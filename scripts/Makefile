go run:
	cd ../ && go run cmd/api/main.go

go migrate:
	cd ../ && go run cmd/migrate/main.go

docker-compose db-dev up:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev up -d

docker-compose db-dev shutdown:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev down

docker-compose db-dev logs:
	cd ../deployments/docker && docker-compose -f db/docker-compose.yml --env-file ../../.env -p learn-go-api-dev logs -f