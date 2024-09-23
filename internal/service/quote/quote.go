package quote

import (
	"bufio"
	"bytes"
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed quotes.txt
var quotes []byte

type QuotesService struct {
	quotes []string
}

func New() *QuotesService {
	qs := QuotesService{}

	reader := bytes.NewReader(quotes)
	s := bufio.NewScanner(reader)

	for s.Scan() {
		if q := strings.TrimSpace(s.Text()); q != "" {
			qs.quotes = append(qs.quotes, q)
		}
	}

	return &qs
}

func (q *QuotesService) GetQuote() (string, error) {
	return q.quotes[rand.Intn(len(q.quotes))], nil
}
