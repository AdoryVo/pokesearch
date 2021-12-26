VERSION = vlatest
RELEASE_DIR = dist/$(VERSION)

build: release build_win build_mac build_lin

release:
	mkdir $(RELEASE_DIR) 

build_win:
	GOOS=windows GOARCH=amd64 go build
	mv pokesearch.exe $(RELEASE_DIR)/pokesearch$(VERSION)-win-amd64.exe

build_mac:
	GOOS=darwin GOARCH=amd64 go build
	mv pokesearch $(RELEASE_DIR)/pokesearch$(VERSION)-mac-amd64

build_lin:
	GOOS=linux GOARCH=amd64 go build
	mv pokesearch $(RELEASE_DIR)/pokesearch$(VERSION)-lin-amd64
