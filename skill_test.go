package forgot

import (
	"testing"
)

func TestB2S(t *testing.T) {
	s := "No zuo no die, why still try?"
	b := []byte(s)
	ret := B2S(b)
	t.Log(s)
	t.Log(b)
	t.Log("B2S result:")
	t.Log(ret)
	t.Log([]byte(ret))
}

func BenchmarkB2S(b *testing.B) {
	buf := []byte("No zuo no die, why still try?")
	for i := 0; i < b.N; i++ {
		_ = B2S(buf)
	}
}

func Benchmark_B2SB_Normal(b *testing.B) {
	buf := []byte("No zuo no die, why still try?")
	for i := 0; i < b.N; i++ {
		_ = string(buf)
	}
}

func TestS2B(t *testing.T) {
	s := "No zuo no die, why still try?"
	b := []byte(s)
	ret := S2B(&s)
	t.Log(s)
	t.Log(b)
	t.Log("S2B result:")
	t.Log(string(ret))
	t.Log(ret)
}

func Benchmark_S2B_Normal(b *testing.B) {
	s := "No zuo no die, why still try?"
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func Benchmark_S2B(b *testing.B) {
	s := "No zuo no die, why still try?"
	for i := 0; i < b.N; i++ {
		_ = S2B(&s)
	}
}
