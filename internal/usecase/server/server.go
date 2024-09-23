package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

const (
	DELIMITER byte = '\n'
	QUIT_SIGN      = "quit!"
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
	msg, err := utils.Read(conn, DELIMITER)
	if err != nil {
		return err
	}
	fmt.Println("Server: ", "message from client: ", msg)

	return nil
}
