package word_of_wisdom

import "context"

type Power interface {
	Calculate(ctx context.Context, data string) uint64
	Check(data string, nonce uint64) bool
}

type Quoter interface {
	GetQuote() (string, error)
}
