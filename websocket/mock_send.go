package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/websocket/ws"
	"net/http"
	"strconv"
	"time"
)

// mock数据测试发送消息
func mockSendHandler(c *gin.Context) {
	val, b := c.GetQuery("num")
	var num int
	if !b {
		num = 10000
	} else {
		n, _ := strconv.Atoi(val)
		num = n
	}
	val, b = c.GetQuery("sleep")
	var sleep int
	if !b {
		sleep = 1000
	} else {
		n, _ := strconv.Atoi(val)
		sleep = n
	}

	go mockData(num, sleep)
	c.String(http.StatusOK, "%d", num)

	//strArray := strings.Split(string(value), ",")
	//time, _ := strconv.ParseInt(strArray[3], 10, 64)
	//o, _ := strconv.ParseFloat(strArray[4], 64)
	//h, _ := strconv.ParseFloat(strArray[5], 64)
	//l, _ := strconv.ParseFloat(strArray[6], 64)
	//close, _ := strconv.ParseFloat(strArray[7], 64)
	//response := &ResponseMessage{
	//	Time:   int64(time),
	//	Symbol: strArray[2],
	//	Data:   []float64{o, h, l, close},
	//}
	//WsManager.send(response)
}

func mockData(num int, sleep int) {
	symbols := []string{"09988.hk", "00700.hk", "65539.hk", "63312.hk", "56495.hk", "52648.hk", "56335.hk",
		"69461.hk", "62833.hk", "03690.hk", "01810.hk", "03968.hk"}

	for i := 0; i < num; i++ {
		inx := i % len(symbols)
		gap := float64(inx)
		response := &ws.ResponseMessage{
			Time:   time.Now().UnixMilli(),
			Symbol: symbols[inx],
			Data:   []float64{100.12 + gap, 120.01 + gap, 90.8, 90.0 + gap/10},
		}
		ws.Manager.Send(response)
		time.Sleep(time.Duration(sleep * 1000))
	}
}
