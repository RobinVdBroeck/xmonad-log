BINARY_NAME=xmonad-log

LD_FLAGS=-X 'main.BuildTime=$(shell date)' -X 'make.GoVersion=$(shell go version)'


all: build

dep: 
	go mod download

build: main.go
	go build -ldflags="${LD_FLAGS}" -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean

install: build
	cp ${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

uninstall:
	rm /usr/local/bin/${BINARY_NAME}

