package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jordanadams/tagsky/internal/cmd"
)

func main() {
	err := cmd.Execute(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
