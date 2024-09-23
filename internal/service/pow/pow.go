package pow

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync/atomic"
)

type Pow struct {
	difficulty int
}

func New(difficulty uint8) *Pow {
	return &Pow{difficulty: int(difficulty)}
}

func (s *Pow) Calculate(ctx context.Context, data string) uint64 {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var br atomic.Bool
	target := strings.Repeat("0", s.difficulty)

	go func() {
		<-ctx.Done()
		br.Store(true)
	}()

	// 1. Start with nonce at 0
	nonce := uint64(0)

	for {
		// 2. Combine the block data and Nonce, i.e., the calculation target is the string "Hello, world!0."
		dataToHash := data + fmt.Sprint(nonce)

		// 3. Apply the SHA-256 hash function to "Hello, world!0" and calculate the hash value.
		hasher := sha256.New()
		hasher.Write([]byte(dataToHash))
		hash := hex.EncodeToString(hasher.Sum(nil))

		// 4. Check if the calculated hash value meets the Difficulty conditions (starting with "0000").
		if strings.HasPrefix(hash, target) {
			// 6. If true, the calculation is complete.
			fmt.Printf("Found a valid hash: %s (Nonce: %d)\n", hash, nonce)
			break
		} else {
			// 5. If not met, increase the Nonce by 1 and repeat steps 2–4.
			nonce++
		}

		if br.Load() {
			nonce = 0
			break
		}
	}

	return nonce
}

func (s *Pow) Check(data string, nonce uint64) bool {
	target := strings.Repeat("0", s.difficulty)
	dataToHash := data + fmt.Sprint(nonce)

	hasher := sha256.New()
	hasher.Write([]byte(dataToHash))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return strings.HasPrefix(hash, target)
}
