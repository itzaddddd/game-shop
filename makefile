# service is one of {auth,player,item,inventory}
service=

run: 
	go run main.go ./env/dev/.env.$(service)

migrate:
	go run ./pkg/db/script/migration.go ./env/dev/.env.$(service)

docker-compose-db-up:
	docker compose -f docker-compose.db.yml up

grpc-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./service/$(service)/proto/$(service).proto


.PHONY: grpc-gen run migrate docker-compose-db-up