package pow

import (
	"github.com/denislyubo/word-of-wisdom/internal/utils"
)

type pow struct {
	isPrime []bool
	primes  []int
	n       int
}

func New(n int) *pow {
	isPrime := make([]bool, n+1)
	utils.SieveOfEratosthenes(n, isPrime)
	return &pow{n: n, isPrime: isPrime}
}

func (s *pow) Calculate(number int) (int, int, error) {
	a, b, err := utils.FindPrimePair(number)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}
