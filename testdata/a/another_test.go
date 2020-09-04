package a_test

import "github.com/shibh308/testinfo/cursorfuncinfo/testdata/a"
import "testing"

func BenchmarkA(b *testing.B) {
	b.StartTimer()
	for i := 0; i < 100; i++ {
		if a.A(i, i) != i*i {
			b.Error("not correct")
		}
	}
	b.StopTimer()
}

func Test_B(t *testing.T) {
	tests := []struct {
		name string
		x    int
		out  int
	}{
		{"case1", 5, 25},
		{"case2", 2, 10},
	}
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if a.B(c.x) != c.out {
				t.Error("not correct")
			}
		})
	}
}

func TestY(t *testing.T) {
}
