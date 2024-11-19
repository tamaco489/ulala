AWS_PROFILE_NAME := sample_profile
MINIO_CLI_ACCESS_PORT := 9000
MINIO_BUCKET_NAME := sample_bucket
AWS_PROFILE_NAME_USE_ECR := sample_profile
AWS_REGION := sample_region
AWS_PROJECT_ID := aws_project_id
ECR_FRONTEND_REPO_NAME := repository_name
ECR_BACKEND_REPO_NAME := repository_name

minio-sync:
	aws --profile ${AWS_PROFILE_NAME} \
		--endpoint-url http://localhost:${MINIO_CLI_ACCESS_PORT}/ \
		s3 sync ./public/proto/ s3://${MINIO_BUCKET_NAME}/

minio-purge:
	aws --profile ${AWS_PROFILE_NAME} \
		--endpoint-url http://localhost:${MINIO_CLI_ACCESS_PORT}/ \
		s3 rm s3://${MINIO_BUCKET_NAME}/ --recursive

dc-up:
	docker-compose -f ./docker-compose.yml up -d
	docker-compose -f ./docker-compose.yml logs -f

dc-down:
	docker-compose -f ./docker-compose.yml down

mysql:
	docker-compose exec mysql mysql -u proto -ppassword -h localhost -P 3306 proto

migrate:
	cd ./go/cmd/seed; make go-db-init; cd -

go-run-server:
	cd ./go/; go run ./main.go; cd -

all-ecr-push:
	ecr-auth frontend-ecr-push backend-ecr-push

ecr-auth:
	aws ecr get-login-password --region ${AWS_REGION} --profile ${AWS_PROFILE_NAME_USE_ECR} | docker login --username AWS --password-stdin ${AWS_PROJECT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

frontend-ecr-push:
	docker build -f Dockerfile.frontend -t ${ECR_FRONTEND_REPO_NAME} .
	docker tag ${ECR_FRONTEND_REPO_NAME}:latest ${AWS_PROJECT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_FRONTEND_REPO_NAME}:latest
	docker push ${AWS_PROJECT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_FRONTEND_REPO_NAME}:latest

backend-ecr-push:
	docker build -f Dockerfile.backend -t ${ECR_BACKEND_REPO_NAME} .
	docker tag ${ECR_BACKEND_REPO_NAME}:latest ${AWS_PROJECT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_BACKEND_REPO_NAME}:latest
	docker push ${AWS_PROJECT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_BACKEND_REPO_NAME}:latest
