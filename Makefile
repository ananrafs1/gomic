
host=template
title=random
Page=1
Quantity=10

runScrap:
	go run ./pkg/scrap/main.go -host=$(host) -title=$(title) -Page=$(Page) -Quantity=$(Quantity)


runServer: 
	@go build -o ./bin/server ./pkg/server/main.go
	./bin/server

updateProto:
	protoc ./orchestrator/proto/scrapper.proto --go-grpc_out=. --go_out=. --go-grpc_opt=require_unimplemented_servers=false

rundocker:
	@docker-compose up -d

