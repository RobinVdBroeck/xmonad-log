DIST_DIRS:= find * -type d -exec
VERSION:=$(shell git describe --tags)

all: vendor xmonad-log

xmonad-log: main.go
	go build -o $@ $^

vendor: go.sum go.mod
	go mod download

clean:
	rm -rf ./dist
	rm -f ./xmonad-log

build-all: vendor
	gox -verbose \
		-os="linux" \
		-arch="amd64 386" \
		-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}"

dist: build-all
	cd dist && \
		$(DIST_DIRS) tar -zcf xmonad-log-${VERSION}-{}.tar.gz -C {} xmonad-log \; && \
		$(DIST_DIRS) zip -r xmonad-log-${VERSION}-{}.zip -j {}/xmonad-log \; && \
		cd ..
		

.PHONY: all build-all clean dist
