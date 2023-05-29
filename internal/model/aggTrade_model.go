package model

/*
AggTradeModel
	{
	  "e": "aggTrade",  // 事件类型
	  "E": 123456789,   // 事件时间
	  "s": "BNBBTC",    // 交易对
	  "a": 12345,       // 归集交易ID
	  "p": "0.001",     // 成交价格
	  "q": "100",       // 成交数量
	  "f": 100,         // 被归集的首个交易ID
	  "l": 105,         // 被归集的末次交易ID
	  "T": 123456785,   // 成交时间
	  "m": true,        // 买方是否是做市方。如true，则此次成交是一个主动卖出单，否则是一个主动买入单。
	}
*/
type AggTradeModel struct {
	Stream string `json:"stream"`
	Data   struct {
		E  string `json:"e"` // 事件类型
		E0 int64  `json:"E"` // 事件时间
		S  string `json:"s"` // 交易对
		A  int    `json:"a"` // 归集交易ID
		P  string `json:"p"` // 成交价格
		Q  string `json:"q"` // 成交数量
		F  int    `json:"f"` // 被归集的首个交易ID
		L  int    `json:"l"` // 被归集的末次交易ID
		T  int64  `json:"T"` // 成交时间
		M  bool   `json:"m"` // 买方是否是做市方。如true，则此次成交是一个主动卖出单，否则是一个主动买入单。
	} `json:"data"`
}
