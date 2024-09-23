package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

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

	_, err = utils.Write(conn, "3:7\n")
	log.Println("Client: ", "3:7")
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
