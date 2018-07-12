package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Gnidleif/discripter/logger"
	"github.com/Gnidleif/discripter/server"
)

var (
	token     string
	scriptDir string
	logDir    string
)

func main() {
	if err := logger.Start(logDir); err != nil {
		log.Fatal(err)
	}
	defer logger.Stop()

	if err := server.Start(token, scriptDir); err != nil {
		log.Fatal(err)
	}
	defer server.Stop()

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	<-make(chan struct{})
}

func init() {
	flag.StringVar(&token, "t", os.Getenv("DGB_TOKEN"), "Bot token")
	flag.StringVar(&scriptDir, "d", os.Getenv("DG_SCRIPTS"), "Script directory")
	flag.StringVar(&logDir, "l", "/tmp/disclog", "Log directory")
	flag.Parse()

	if token == "" || scriptDir == "" {
		log.Fatal(errors.New("flag error: token and script directory can't be empty"))
	}
}
