package cmd

import (
	"fmt"
	"sync"

	"github.com/rickymail9527/go-binance-demo/internal/client"
	"github.com/rickymail9527/go-binance-demo/internal/redisCtrl"
	"github.com/rickymail9527/go-binance-demo/internal/server"
)

func Start() {
	fmt.Println("Start, World!")
	redisCtrl.RedisNewClient()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		server.StartWebSocketServer()
		wg.Done()
	}()
	go func() {
		client.StartWebSocketClient()
		wg.Done()
	}()
	go func() {
		server.HttpStart()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("End, World!")
}
