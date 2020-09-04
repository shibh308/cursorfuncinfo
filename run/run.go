package run

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/shibh308/cursorfuncinfo/analysis"
	"os"
)

type Flags struct {
	Parse  bool
	Path   string `short:"f" long:"file" description:"file path" required:"true"`
	Offset int    `short:"p" long:"pos" description:"position of cursor" required:"true"`
}

var op = Flags{Parse: true}

func SetFlags(fl Flags) {
	op = fl
}

func Run() (string, error) {
	ctx := context.Background()
	if op.Parse {
		_, err := flags.Parse(&op)
		if err != nil {
			return "", err
		}
	}

	_, err := os.Stat(op.Path)
	if err != nil {
		return "", fmt.Errorf("file stat error %s: \"%s\"", op.Path, err.Error())
	}
	a, err := analysis.New(ctx, op.Path, op.Offset)
	if err != nil {
		return "", fmt.Errorf("failed to make analysis object: \"%s\"", err.Error())
	}
	res, err := a.GetFuncData()
	if err != nil {
		return "", fmt.Errorf("failed to analyze: \"%s\"", err.Error())
	}
	if res == nil {
		return "", fmt.Errorf("pos not in any func decl")
	}
	j, err := json.Marshal(res)
	if err != nil {
		return "", fmt.Errorf("failed to load json: \"%s\"", err.Error())
	}
	return string(j), nil
}
