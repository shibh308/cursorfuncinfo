package run_test

import (
	"fmt"
	"github.com/shibh308/cursorfuncinfo/run"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		Name string
		File string
		Pos  int
	}{
		{"a_f", "../testdata/a/a.go", 40},
		{"a_G", "../testdata/a/a.go", 80},
		{"a_A", "../testdata/a/a.go", 110},
		{"a_B", "../testdata/a/a.go", 130},
		{"a_X", "../testdata/a/a.go", 170},
		{"a_testF", "../testdata/a/a_test.go", 400},
		{"a_testG", "../testdata/a/a_test.go", 140},
		{"a_testA", "../testdata/a/another_test.go", 110},
		{"a_testB", "../testdata/a/another_test.go", 280},
		{"a_testY", "../testdata/a/another_test.go", 530},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			run.SetFlags(run.Flags{
				Parse:  false,
				Path:   c.File,
				Offset: c.Pos,
			})
			str, err := run.Run()
			fmt.Println(str)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
