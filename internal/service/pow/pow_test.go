package pow

import (
	"context"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	ctx := context.Background()

	p := New(6)

	data := "Hello, world!"
	nonce := p.Calculate(ctx, data)

	res := p.Check(data, nonce)
	if !res {
		t.Error("check error")
	}
}

func TestCalculateTimeout(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	p := New(6)

	data := "Hello, world!"
	nonce := p.Calculate(ctx, data)

	res := p.Check(data, nonce)
	if res {
		t.Error("check error")
	}
}
