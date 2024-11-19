package infrastructure_test

import (
	"context"
	"testing"
	"time"

	"github.com/miyabiii1210/ulala/go/datastore"
)

func TestNewConnection(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Database Connection Test",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := datastore.NewDBConnection()
			if client.Error != nil {
				t.Errorf("db connection error: %v", client.Error)
				return
			}

			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestNewRedisConnection(t *testing.T) {
	type args struct {
		ctx        context.Context
		key, value string
		expiration time.Duration // keyの有効期限
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Redis Connection Test",
			args: args{
				ctx:        context.Background(),
				key:        "haru",
				value:      "ulala",
				expiration: 20 * time.Second, // n秒後にkeyを消去
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := datastore.NewRedisConnection()
			t.Logf("%s redis client: \n", client) // Redis<127.0.0.1:6379 db:0>

			p := client.Ping(context.Background())
			if p.Err() != nil {
				t.Errorf("ping error: %v", p.Err())
				return
			}
			t.Logf("ping: %v", p.Val()) // ping: PONG

			setRst := client.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.expiration)
			if setRst.Err() != nil {
				t.Errorf("set error: %v", setRst.Err())
				return
			}

			getRst, err := client.Get(tt.args.ctx, tt.args.key).Result()
			if err != nil {
				t.Errorf("get error: %v", err)
				return
			}
			t.Logf("get: %v", getRst) // get: ulala

			t.Logf("%s End\n", tt.name)
		})
	}
}
