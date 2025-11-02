package main

import (
	"fmt"
	"os"

	structcopygen "github.com/bookweb/structcopy-gen"
	"github.com/bookweb/structcopy-gen/config"
)

func main() {
	fmt.Println(config.Version)
	if err := structcopygen.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
