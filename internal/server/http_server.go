package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rickymail9527/go-binance-demo/internal/routerCtrl"
)

const httpPort = "8081"

// New Router
func registerMemberRouter(router *gin.Engine) {
	new(routerCtrl.ApiRouter).Router(router)
}

func HttpStart() {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.Use(gin.Recovery())

	registerMemberRouter(engine)
	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: engine,
	}
	go func() {
		_ = server.ListenAndServe()
	}()

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		sig := <-quit

		fmt.Println("got a signal", sig)
		now := time.Now()
		cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(cxt)
		if err != nil {
			fmt.Println("err", err)
		}
		// Time taken to actually exit
		fmt.Println("------exited--------", time.Since(now))
		w.Done()
	}()
	w.Wait()
	log.Println("All servers stopped. Exiting.")
	os.Exit(0)
}
