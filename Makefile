GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
BIN_NAME=discripter
BIN_UNIX=$(BIN_NAME)_unix
BIN_PI=$(BIN_NAME)_pi
COVER_NAME=cover.out
COVER_HTML=cover.html

all: test build
build:
	$(GOBUILD) -o $(BIN_NAME) -v
test:
	$(GOTEST) -v -cover -coverprofile=$(COVER_NAME) ./...
	$(GOTOOL) cover -html=$(COVER_NAME) -o $(COVER_HTML)
clean:
	$(GOCLEAN)
	rm -f $(BIN_NAME)
	rm -f $(BIN_UNIX)
	rm -f $(COVER_NAME)
	rm -f $(COVER_HTML)
deps:
	$(GOGET) github.com/bwmarrin/discordgo

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_UNIX) -v
