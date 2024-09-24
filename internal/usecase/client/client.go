package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	schema "github.com/denislyubo/word-of-wisdom"
	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/service/pow"
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

type client struct {
	config *config.ClientConfig
	pow    schema.Power
}

func New(config *config.ClientConfig) *client {
	pow := pow.New(config.Difficulty)
	return &client{config: config, pow: pow}
}

func (c *client) GetQuote() (*string, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.config.ServerHost, c.config.ServerPort))
	if err != nil {
		log.Println("Client: ", "dial err:", err)
		return nil, err
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Client: ", "close err:", err)
		}
	}()

	_, err = utils.Write(conn, "Hello server\n")
	log.Println("Client: ", "Hello server")
	if err != nil {
		log.Println("Client: ", "Write Error: ", err.Error())
		return nil, err
	}

	msg, err := utils.Read(conn, utils.DELIMITER)
	if err != nil {
		log.Println("Client: ", "Read Error: ", err.Error())
		return nil, err
	}

	strs := strings.Split(msg, ": ")
	if len(strs) != 2 {
		log.Println("Client: ", "Unexpected message from server: ", msg)
		return nil, errors.New("unexpected message from server")
	}

	if strs[0] != "Hello, solve puzzle" {
		log.Println("Client: ", "Unexpected message from server: ", msg)
		return nil, errors.New("unexpected message from server")
	}

	nonce := c.pow.Calculate(context.TODO(), strs[1])

	_, err = utils.Write(conn, fmt.Sprintf("nonce:%d\n", nonce))
	log.Println("Client: ", "sent nonce: ", nonce)
	if err != nil {
		log.Println("Client: ", "Write Error: ", err.Error())
		return nil, err
	}

	msg, err = utils.Read(conn, utils.DELIMITER)
	if err != nil {
		log.Println("Client: ", "Read Error: ", err.Error())
		return nil, err
	}

	return &msg, nil
}
