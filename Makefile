.PHONY: build
build:
	GOMOUDULE=on CGO_ENABLED=0 go build -o bin/o11ybench cmd/o11ybench/main.go

.PHONY: clean
clean:
	rm -rf bin
