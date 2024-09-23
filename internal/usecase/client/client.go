package client

import (
	"fmt"
	"log"
	"net"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

type client struct {
	config *config.ClientConfig
}

func New(config *config.ClientConfig) *client {
	return &client{config: config}
}

func (c *client) GetQuote() (*string, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.config.ServerHost, c.config.ServerPort))
	defer func() {
		if err != nil {
			log.Println("Client: ", "dial err:", err)
		}
		if err := conn.Close(); err != nil {
			log.Println("Client: ", "close err:", err)
		}
	}()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = utils.Write(conn, "Hello from client\n")
	if err != nil {
		log.Println("Client: ", "Write Error: ", err.Error())
		return nil, err
	}

	return nil, nil
}
