.PHONY: build-darwin
build-darwin:
	GOOS=darwin go build -ldflags="-s -w" -o pbl

.PHONY: build-linux
build-linux:
	go build -ldflags="-s -w" -o pbl
