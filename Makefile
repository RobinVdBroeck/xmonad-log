BINARY_NAME=xmonad-log

all: build

dep: 
	go mod download

build: main.go
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean

install: build
	cp ${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

uninstall:
	rm /usr/local/bin/${BINARY_NAME}

