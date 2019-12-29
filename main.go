package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	router := gin.Default()
	router.Static("/", "./public")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return exitCodeFailed
	}
	return exitCodeOK
}
