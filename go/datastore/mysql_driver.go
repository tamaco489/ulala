package datastore

import (
	"errors"
	"sync"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/miyabiii1210/ulala/go/config"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var instance *gorm.DB // application側に渡すgorm構造体
var once sync.Once    // 指定した関数を1度だけ実行

func NewDBConnection() *gorm.DB {
	once.Do(func() {
		logMode := logger.Silent // 一旦ローカル環境、開発環境以外は全てログ出力をしない設定にする
		if config.IsLocal() || config.IsDevelopment() {
			logMode = logger.Info
		}

		db, err := gormtrace.Open(
			mysql.New(mysql.Config{
				DSN: getDbDsn(),
			}),
			&gorm.Config{
				Logger: logger.Default.LogMode(logMode),
			},
			gormtrace.WithServiceName(config.EnvConfig.ApplicationName),

			// エラーチェック用の関数、指定したエラーが特定のエラーと一致しない場合にtrueを返す
			gormtrace.WithErrorCheck(func(err error) bool {
				// レコードが見つからない場合のエラー、一旦このエラーとは異なるエラーが発生した場合にエラーチェックを有効にする
				return !errors.Is(err, gorm.ErrRecordNotFound)
			}),
		)

		if err != nil {
			panic(err)
		}

		// application側で使用するDB接続子 (gorm構造体) を生成
		if err = db.Use(dbresolver.Register(
			// リゾルバ設定 (レプリカ用のDB接続子を生成)
			dbresolver.Config{
				Replicas: []gorm.Dialector{
					mysql.New(mysql.Config{
						DSN: getDbSlaveUrl(),
					}),
				},
				Policy: dbresolver.RandomPolicy{},
			}).

			// マスターデータベース、レプリカデータベースの共通設定
			SetConnMaxIdleTime(time.Hour).      // 接続がアイドル状態で利用されていない場合、最大でどれくらいの時間プールに保持するか (一旦1hで設定)
			SetConnMaxLifetime(24 * time.Hour). // 接続されてから、最大でどれくらいの時間その接続を有効と見做すか (一旦24hで設定)
			SetMaxIdleConns(100).               // データベースに接続されてからアイドル状態でプールするコネクションの最大数
			SetMaxOpenConns(200),               // データベースに接続できるコネクションの最大数
		); err != nil {
			panic(err)
		}

		// gormプリロード機能を使用して、gormのセッションを作成 (トランザクションやクエリの実行、DB操作のコンテキスト管理のために使用)
		instance = db.Preload(clause.Associations).Session(&gorm.Session{
			// 関連するモデル (オブジェクト) の保存操作が行われた際に、そのモデルに関連するデータを含めて保存処理を行う設定
			FullSaveAssociations: true,
		})
	})

	return instance
}

func getDbDsn() string {
	c := mysqldriver.Config{
		DBName:               config.EnvConfig.DBName,
		User:                 config.EnvConfig.DBUser,
		Passwd:               config.EnvConfig.DBPassword,
		Addr:                 config.EnvConfig.DBHost,
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
	}

	return c.FormatDSN()
}

func getDbSlaveUrl() string {
	return getDbDsn() // 一旦マスターと同じ接続先を指定
}
