package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Init() {
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	port := os.Getenv("PORT")
	port = "1234"
	if port == "" {
		return errors.New("env PORT is empty")
	}
	addr := fmt.Sprintf(":%s", port)
	r := gin.New()

	r.GET("/", ping)

	return r.Run(addr)
}

func ping(ctx *gin.Context) {
	now := time.Now()
	mess := fmt.Sprintf("pong time %v", now)
	ctx.JSON(http.StatusOK, gin.H{"message": mess})
}
