
host=template
title=random

runScrap:
	go run ./pkg/scrap/main.go -host=$(host) -title=$(title)