package client

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rickymail9527/go-binance-demo/internal/redisCtrl"
	"github.com/rickymail9527/go-binance-demo/internal/until"
)

const socketUrl = "wss://stream.binance.com:9443/stream?streams=btcusdt@aggTrade"

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	redisClient := redisCtrl.GetClient()
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		//log.Printf("Received: %s\n", msg)

		// check redis lock
		data, _ := redisClient.Get(until.RedisLockKey).Result()
		if data != "1" {
			// lock
			redisClient.Set(until.RedisLockKey, "1", 16*time.Second)
			// redis set
			redisClient.Set(until.RedisAggTradeKey, msg, 0)
			// unlock
			redisClient.Set(until.RedisLockKey, "0", 0)
		}
	}
}

func StartWebSocketClient() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	//socketUrl := "ws://localhost:8080" + "/socket"
	//socketUrl := "wss://stream.binance.com:9443/stream?streams=btcusdt@aggTrade"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)

	// Our main loop for the client
	// We send our relevant packets here
	for {
		select {
		//case <-time.After(time.Duration(1) * time.Millisecond * 1000 * 5):
		//	// Send an echo packet every second
		//	err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from GolangDocs!"))
		//	if err != nil {
		//		log.Println("Error during writing to websocket:", err)
		//		return
		//	}

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}
