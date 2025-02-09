package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/btcthirst/gin-001/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	port := viper.Get("PORT")
	if port == "" {
		port = "8080"
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
