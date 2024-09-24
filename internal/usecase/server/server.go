package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	schema "github.com/denislyubo/word-of-wisdom"
	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/service/pow"
	"github.com/denislyubo/word-of-wisdom/internal/service/quote"
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

type server struct {
	config   *config.ServerConfig
	listener net.Listener
	pow      schema.Power
	quote    schema.Quoter
}

func New(config *config.ServerConfig) *server {
	p := pow.New(config.Difficulty)
	q := quote.New()
	return &server{config: config, pow: p, quote: q}
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
		if err != nil {
			log.Println("Server: ", "listener close error: ", err.Error())
		}
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
	puzzle := "Puzzle" + string(rand.Intn(100))
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

	msg = fmt.Sprintf("Hello, solve puzzle: %s\n", puzzle)
	_, err = utils.Write(conn, msg)
	log.Println("Server: ", "message to client: ", msg)
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

	if strs[0] != "nonce" {
		log.Println("Server: ", "Wrong puzzle answer: ")
		return errors.New("wrong puzzle answer")
	}

	nonce, err := strconv.ParseUint(strs[1], 10, 64)
	if err != nil {
		log.Println("Server: ", "Wrong puzzle answer: ")
		return errors.New("wrong puzzle answer")
	}

	if !s.pow.Check(puzzle, nonce) {
		log.Println("Server: ", "Wrong puzzle answer")
		return errors.New("wrong puzzle answer")
	}

	log.Println("Server: ", "Success")
	q, err := s.quote.GetQuote()
	if err != nil {
		log.Println("Server: ", "Get Quote Error: ", err.Error())
	}
	_, err = utils.Write(conn, q)
	if err != nil {
		return err
	}

	return nil
}
