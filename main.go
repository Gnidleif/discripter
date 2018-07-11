package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/Gnidleif/discripter/logger"
	"github.com/Gnidleif/discripter/server"
)

var (
	Token     string
	ScriptDir string
)

func main() {
	if err := logger.Start("disclog"); err != nil {
		log.Fatal(err)
	}
	defer logger.Stop()

	if err := server.Start(Token, ScriptDir); err != nil {
		log.Fatal(err)
	}
	defer server.Stop()

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	<-make(chan struct{})
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&ScriptDir, "d", "", "Script directory")
	flag.Parse()

	if Token == "" || ScriptDir == "" {
		log.Fatal(errors.New("flag error: token and script directory can't be empty"))
	}
}
