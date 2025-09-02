go run:
	cd ../ && go run cmd/api/main.go

docker-compose db-dev up:
	cd ../deployments/docker && docker-compose -f docker-compose.db.yml -p learn-go-api-dev up -d
docker-compose db-dev shutdown:
	cd ../deployments/docker && docker-compose -f docker-compose.db.yml -p learn-go-api-dev down
docker-compose db-dev logs:
	cd ../deployments/docker && docker-compose -f docker-compose.db.yml -p learn-go-api-dev logs -f