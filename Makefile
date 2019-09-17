.PHONY: all clean build deps install uninstall reinstall

PREFIX = /usr/local
GOBUILD = go build
GORUN = go run

all: clean build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/goautoplex cmd/api/main.go
	install -m0755 srv/filebot bin/filebot

install: bin
	install -d -m0755 $(PREFIX)/plexbot/bin
	install -m0755 bin/* $(PREFIX)/plexbot/bin

uninstall:
	rm -rf $(CONFIG_DIR)
	rm -rf $(PREFIX)/plexbot

reinstall:
	rm -rf $(PREFIX)/plexbot
	install -d -m0755 $(PREFIX)/plexbot/bin
	install -m0755 bin/* $(PREFIX)/plexbot/bin
