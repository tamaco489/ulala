### ulala
---

### building development environment
---
* create .env file
  ```
  $ cat << EOF > .env
  GO_ENV='dev'
  API_DOMAIN='localhost'
  PERMISSION='user'
  GOOGLE_APPLICATION_CREDENTIALS ='./credentials/xxx.json'
  GOOGLE_SPREADSHEET_ID='xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
  EOF
  ```

* starting docker containers
  ```
  docker-compose -f ./docker-compose.yml up -d
  ```
  Launching the container starts both frontend and backend, and bastion migrates the DB.

### DB migrate
  ```
  make mysql
  ```
  ```
  make migrate
  ```


### backend
---
* start backend server
  ```
  make go-run-server
  ```

* api verification
  ```
  curl -v -X GET -H "Authorization: Bearer ABCDEFG123456789" \
      -H "Content-Type: application/json" \
      localhost:8080/movies/type\?id=1 | jq
  ```
  ```
  {
    "movie_id": 10000001,
    "title": "Die Hard (ダイ・ハード)",
    "release_year": 1988,
    "description": "ニューヨーク市警察の刑事",
    "type_name": "action",
    "movie_format": "mp4"
  }
  ```



### frontend
---
* start frontend server
  ```
  yarn dev
  ```

* browser access
  ```
  open http://localhost:3000
  ```

* prettier
  ```
  $ npx prettier --write src/*
  ```

### minio (strage)
---
* local data sync and purge
  ```
  make minio-sync
  ```
  ```
  make minio-purge
  ```

### deploy (AWS ECR)
* build and push container images for ECR
  ```
  make -j 2 all-ecr-push
  ```