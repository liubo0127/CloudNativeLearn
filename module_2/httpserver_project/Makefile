export tag=v1.0

build_darwin:
	echo "build httpserver binary"
	mkdir -p bin/darwin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/httpserver .

run_darwin: build_darwin
	echo "run httpserver"
	mkdir -p log
	./bin/darwin/httpserver -log_dir=./log/

build_linux_amd64:
	echo "build httpserver binary"
	mkdir -p bin/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/httpserver .

run_linux_amd64: build_linux_amd64
	echo "run httpserver"
	mkdir -p log
	./bin/linux/httpserver -log_dir=./log/