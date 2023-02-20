# Linux Makefile | Golang & zip are required for build.
VERSION = @latest
RELEASE_DIR = dist/$(VERSION)
ARCHIVE_DIR = dist/pokesearch

build: release build_win build_mac build_lin

release:
	mkdir -p $(RELEASE_DIR) 

build_win:
	GOOS=windows GOARCH=amd64 go build -o $(ARCHIVE_DIR)
	zip -FSj $(RELEASE_DIR)/pokesearch$(VERSION)-win-amd64.zip $(ARCHIVE_DIR)/*
	rm $(ARCHIVE_DIR)/pokesearch*

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o $(ARCHIVE_DIR)
	tar -C $(ARCHIVE_DIR) -cvzf $(RELEASE_DIR)/pokesearch$(VERSION)-mac-amd64.tar.gz .
	rm $(ARCHIVE_DIR)/pokesearch*

build_lin:
	GOOS=linux GOARCH=amd64 go build -o $(ARCHIVE_DIR)
	tar -C $(ARCHIVE_DIR) -cvzf $(RELEASE_DIR)/pokesearch$(VERSION)-lin-amd64.tar.gz .
	rm $(ARCHIVE_DIR)/pokesearch*
