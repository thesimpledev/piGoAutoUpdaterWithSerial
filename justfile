GOOS := "linux"
GOARCH := "arm"
BINARY := "bin/bootstrap"
IPADD := '192.168.6.71'

build:
    GOOS={{GOOS}} GOARCH={{GOARCH}} go build -o {{BINARY}} ./cmd/cli/

push:
    scp {{BINARY}} sstanton@{{IPADD}}:/home/sstanton/

stop:
    ssh sstanton@{{IPADD}} "pkill bootstrap || true"

start:
    ssh sstanton@{{IPADD}} "nohup ./bootstrap > /dev/null 2>&1 &"

ssh:
    ssh sstanton@{{IPADD}}

deploy: build stop push start