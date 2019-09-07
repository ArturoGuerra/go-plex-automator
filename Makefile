.PHONY: all clean build deps install uninstall reinstall

PREFIX = /usr/local
GOBUILD = go build
GORUN = go run
CONFIG_DIR = /etc/plexbot
CONFIG = $(CONFIG_DIR)/plexbot.conf

all: clean build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/goautoplex -ldflags "-X main.configDir=$(CONFIG)" cmd/goautoplex/main.go
	install -m0755 srv/deluge bin/deluge
	install -m0755 srv/nzbget bin/nzbget
	install -m0755 srv/filebot bin/filebot

install: bin
	install -d -m0755 $(PREFIX)/plexbot/bin
	install -m0755 bin/* $(PREFIX)/plexbot/bin
	if [ ! -f $(CONFIG) ]; then \
    	install -m0755 src/plexbot.conf $(CONFIG); \
	fi

uninstall:
	rm -rf $(CONFIG_DIR)
	rm -rf $(PREFIX)/plexbot

reinstall:
	rm -rf $(PREFIX)/plexbot
	install -d -m0755 $(PREFIX)/plexbot/bin
	install -m0755 bin/* $(PREFIX)/plexbot/bin

test:
	$(GORUN) cmd/goautoplex/main.go
