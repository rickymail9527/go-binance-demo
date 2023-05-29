# go-binance-demo

* Version 1.0.1

### What is this repository for? ###

接口文檔：
https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams_CN.md

Demo to wss://stream.binance.com:9443

#### 1. 訂閱近期成交(BTCUSDT歸集)功能，
wss://stream.binance.com:9443/stream?streams=btcusdt@aggTrade

#### 2. 將最新的一筆行情存入redis ，鍵名為“streams=btcusdt@aggTrade”

#### 3. 客戶端能訪問獲取最新一筆行情 (ReSTFull API)
請求
```
curl --location --request GET 'localhost:8080/api/v1/binance/agg_trade'
```
接收
```
{
    "code": 0,
    "data": {
        "stream": "btcusdt@aggTrade",
        "data": {
            "e": "aggTrade",
            "E": 1655217984681,
            "s": "BTCUSDT",
            "a": 1205353908,
            "p": "22264.00000000",
            "q": "0.00309000",
            "f": 1406422503,
            "l": 1406422503,
            "T": 1655217984681,
            "m": true
        }
    }
}
```

#### 4. 客戶端獲取最新行情的推送 (Websocket)
請求
```
ws://127.0.0.1:8080/socket/get_agg_trade
```
接收
```
{
    "stream": "btcusdt@aggTrade",
    "data": {
        "e": "aggTrade",
        "E": 1655220702627,
        "s": "BTCUSDT",
        "a": 1205418652,
        "p": "22393.95000000",
        "q": "0.00081000",
        "f": 1406504277,
        "l": 1406504277,
        "T": 1655220702626,
        "m": true
    }
}
```
