package main

import (
	"fmt"
	"log"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/usecase/client"
)

func main() {
	var cfg config.ClientConfig
	c, err := config.Load(cfg)
	if err != nil {
		log.Fatal("Client: ", "failed to load config: ", err.Error())
	}

	quote, err := client.New(c).GetQuote()
	if err != nil {
		log.Fatal("Client: ", "failed to get quote: ", err.Error())
		return
	}

	fmt.Println("Client: ", "message from server: ", quote)
}
