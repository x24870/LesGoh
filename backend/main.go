package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/x24870/LesGoh/backend/chain"
)

const (
	port = ":8080"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	adapter := chain.NewAdapter()

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	chain.RegisterAPI(router, adapter)

	serv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := serv.ListenAndServe(); err != nil {
		logrus.Info("stop serv.ListenAndServe: %v", err)
	}

	logrus.Info("end...")
}
