# Go

<br>

### Setup Database
---

* setup database (create DDL, seed loader)
  migration and seed loader:
  ```
  cd ./go/cmd/seed; make go-db-init; cd -
  ```
  options:
    * db reset only: `cd ./go/cmd/seed; make go-db-reset; cd -`
    * db migrate only: `cd ./go/cmd/seed; make go-db-migrate; cd -`
    * seed loader only: `cd ./go/cmd/seed; make go-seed; cd -`

* create DDL
  ```
  migrate create -ext sql -dir ./go/cmd/migrate/ddl -seq ${schema_file_name}
  ```
  ```
  tree ./go/cmd/migrate/ddl
  ./go/cmd/migrate/ddl
  ├── 000001_create_users_table.down.sql
  ├── 000001_create_users_table.up.sql
  ├── 000002_create_auth_table.down.sql
  ├── 000002_create_auth_table.up.sql
  ├── 000003_create_movies_table.down.sql
  ├── 000003_create_movies_table.up.sql
  ...
  ```

<br>

* seed loader using google spreadsheets (Pass environment variables when executing commands)
  ```
  cd ./go; GOOGLE_SPREADSHEET_ID='XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX' go run ./cmd/seed/movieType/main.go; cd -
  ```

* firebase env setting (default env)
  ```
  % grep -n "GoogleApplicationCredentials" ./go/config/env.go
  17:     GoogleApplicationCredentials string `env:"GOOGLE_APPLICATION_CREDENTIALS" envDefault:"./credentials/xxx.json"`
  ```

### Go Test
---

* integration test (CLI)
  ```
  $ go test -v ./test/integration/sample.go -run TestXXX
  ```

* unit test (CLI)
  ```
  $ go test -v ./test/unit/user_test.go -run TestSignUp
  === RUN   TestSignUp
  === RUN   TestSignUp/SignUp_Test
      user_test.go:59: result: &{}
      user_test.go:60: SignUp Test End
  --- PASS: TestSignUp (0.01s)
      --- PASS: TestSignUp/SignUp_Test (0.01s)
  PASS
  ok      command-line-arguments  0.013s
  ```
  ```
  mysql> SELECT uid, name, email FROM users WHERE name = 'hoge1' AND email = 'hoge1@example.com';
  +----------+-------+-------------------+
  | uid      | name  | email             |
  +----------+-------+-------------------+
  | 10000057 | hoge1 | hoge1@example.com |
  +----------+-------+-------------------+
  1 row in set (0.01 sec)
  ```

<br>

### Go Script
---
* resize thumbnail images
  ```
  cd ./go/cmd; make go-thumbnails; cd -
  ```