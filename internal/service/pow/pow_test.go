package pow

import "testing"

func TestCalculate(t *testing.T) {
	p := New(10000)

	a, b, err := p.Calculate(1814503) //1297 1399
	if err != nil {
		t.Error(err)
	}

	if a != 1297 {
		t.Errorf("a != 1297: %d", a)
	}

	if b != 7 {
		t.Errorf("b != 1399: %d", b)
	}
}
