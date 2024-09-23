package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

type server struct {
	config   *config.ServerConfig
	listener net.Listener
}

func New(config *config.ServerConfig) *server {
	return &server{config: config}
}

func (s *server) ListenAndServe(ctx context.Context) (err error) {
	lc := net.ListenConfig{
		KeepAlive: s.config.ServerKeepAlive,
	}

	s.listener, err = lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", s.config.ServerPort))
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		err := s.listener.Close()
		log.Println("Server: ", "listener close error: ", err.Error())
	}()

	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			log.Println("Server: ", "closed")
			return nil
		}

		go func(conn net.Conn) {
			if err = s.handler(conn); err != nil {
				log.Println("Server: ", "handler error: ", err.Error())
			}
		}(conn)
	}
}

func (s *server) handler(conn net.Conn) error {
	defer func() {
		conn.Close()
		log.Println("Server: ", "connection closed")
	}()
	_ = conn.SetDeadline(time.Now().Add(s.config.ServerDeadline))

	msg, err := utils.Read(conn, utils.DELIMITER)
	if err != nil {
		return err
	}
	log.Println("Server: ", "message from client: ", msg)

	_, err = utils.Write(conn, "Hello, solve puzzle: 21\n")
	log.Println("Server: ", "message to client: ", "Hello, solve puzzle: 21")
	if err != nil {
		log.Println("Client: ", "Write Error: ", err.Error())
		return err
	}

	msg, err = utils.Read(conn, utils.DELIMITER)
	if err != nil {
		return err
	}

	strs := strings.Split(msg, ":")
	if len(strs) != 2 {
		log.Println("Server: ", "Unexpected message from client: ", msg)
		return errors.New("unexpected message from client")
	}

	if strs[0] != "3" || strs[1] != "7" {
		log.Println("Server: ", "Wrong puzzle answer: ")
		return errors.New("wrong puzzle answer")
	}

	log.Println("Server: ", "Success")
	_, err = utils.Write(conn, "Quote: la-la-la\n")
	if err != nil {
		return err
	}

	return nil
}
