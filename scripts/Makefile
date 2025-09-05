run:
	cd ../ && air

migration-up:
	cd ../ && go run cmd/migration-up/main.go

swag init:
	cd ../ && swag init -g cmd/api/main.go -d ./ -o ./cmd/api/docs

docker-compose db-dev up:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev up -d

docker-compose db-dev shutdown:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev down

docker-compose db-dev logs:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev logs -f