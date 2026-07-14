package main

import (
	"diskSync/src/internal/config"
	"fmt"
)

func main() {
	_, err := config.Load("src/internal/config/conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
}
