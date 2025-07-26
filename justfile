GOOS := "linux"
GOARCH := "arm"
BINARY := "bin/bootstrap"
IPADD := '192.168.6.71'

build-cli:
    GOOS=linux GOARCH=arm go build -o bin/bootstrap ./cmd/cli/
    
build-reader:
    GOOS=linux GOARCH=amd64 go build -o bin/reader ./cmd/reader/

push:
    scp {{BINARY}} sstanton@{{IPADD}}:/home/sstanton/

stop:
    ssh sstanton@{{IPADD}} "pkill bootstrap || true"

start:
    ssh sstanton@{{IPADD}} "nohup ./bootstrap > /dev/null 2>&1 &"

ssh:
    ssh sstanton@{{IPADD}}

deploy: build-cli build-reader stop push start