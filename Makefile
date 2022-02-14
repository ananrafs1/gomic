
host=template
title=random

runScrap:
	go run ./pkg/scrap/main.go -host=$(host) -title=$(title)

updateProto:
	protoc ./orchestrator/proto/scrapper.proto --go-grpc_out=. --go_out=. --go-grpc_opt=require_unimplemented_servers=false