package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/nhz-io/go-gen-equality"
	"github.com/clipperhouse/typewriter"
)

func main() {
	filter := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go") && !strings.HasSuffix(f.Name(), "_equality.go")
	}

	a, err := typewriter.NewAppFiltered("+test", filter)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := a.WriteAll(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
