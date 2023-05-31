package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rickymail9527/go-binance-demo/internal/model"
	"github.com/rickymail9527/go-binance-demo/internal/redisCtrl"
	"github.com/rickymail9527/go-binance-demo/internal/until"
)

const websocketPort = "8080"

var upGrader = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		message = []byte(string(message) + " from server!!")
		//log.Printf("Received: %s from server log", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func aggTradeHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	redisClient := redisCtrl.GetClient()

	// The event loop
	for {
		var message []byte
		data, _ := redisClient.Get(until.RedisAggTradeKey).Result()
		var res model.AggTradeModel
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			message = []byte("get aggTrade error")
		}
		jsonBytes, _ := json.Marshal(res)
		if err != nil {
			message = []byte("get aggTrade error")
		}
		message = jsonBytes

		//log.Printf("Received: %s from server log", message)
		err = conn.WriteMessage(1, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Index Page")
}

func StartWebSocketServer() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/socket/get_agg_trade", aggTradeHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:"+websocketPort, nil))
}
