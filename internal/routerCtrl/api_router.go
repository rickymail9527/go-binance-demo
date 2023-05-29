package routerCtrl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickymail9527/go-binance-demo/internal/model"
	"github.com/rickymail9527/go-binance-demo/internal/redisCtrl"
	"github.com/rickymail9527/go-binance-demo/internal/until"
)

type ApiRouter struct{}

func (r ApiRouter) Router(router *gin.Engine) {
	base := router.Group("/api")

	v1 := base.Group("/v1")
	{
		v1.GET("/binance/agg_trade", r.GetBinanceNewAggTrade)
	}
}

func (r ApiRouter) GetBinanceNewAggTrade(c *gin.Context) {
	redisClient := redisCtrl.GetClient()
	data, err := redisClient.Get(until.RedisAggTradeKey).Result()
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "",
		})
		return
	}

	var res model.AggTradeModel
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": res,
	})
}
