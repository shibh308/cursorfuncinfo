package a

import "testing"

func TestG(t *testing.T) {
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
			if G(c.x) != c.out {
				t.Error("not correct")
			}
		})
	}
}

func Test_f(t *testing.T) {
	tests := map[string]struct {
		x   int
		y   int
		out int
	}{
		"case1": {5, 3, 15},
		"case2": {2, 4, 8},
	}
	for name, c := range tests {
		t.Run(name, func(t *testing.T) {
			if f(c.x, c.y) != c.out {
				t.Error("not correct")
			}
		})
	}
}
