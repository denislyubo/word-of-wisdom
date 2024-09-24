package main

import (
	"log"
	"sync"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/usecase/client"
)

func main() {
	var (
		cfg config.ClientConfig
		wg  sync.WaitGroup
	)
	err := config.Load(&cfg)
	if err != nil {
		log.Fatal("Client: ", "failed to load config: ", err.Error())
	}

	client := client.New(&cfg)

	var i uint64
	for i = 0; i < cfg.ClientRequests; i++ {
		wg.Add(1)
		go func(i uint64) {
			quote, err := client.GetQuote()
			if err != nil {
				log.Fatal("Client #", i, ": failed to get quote: ", err.Error())
				return
			}

			log.Println("Client #", i, ": message from server: ", *quote)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
