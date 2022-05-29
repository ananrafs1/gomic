
host=template
title=random
Page=1
Quantity=10

runScrap:
	go run ./pkg/scrap/main.go -host=$(host) -title=$(title) -Page=$(Page) -Quantity=$(Quantity)


runServer: monitor
	@go build -o ./bin/server ./pkg/server/main.go
	./bin/server

updateProto:
	protoc ./orchestrator/proto/scrapper.proto --go-grpc_out=. --go_out=. --go-grpc_opt=require_unimplemented_servers=false

monitor: resetdocker
	@docker run -d --name jaeger -p 16686:16686 -p 14268:14268 jaegertracing/all-in-one:1.21

resetdocker:
	@docker rm -f jaeger
