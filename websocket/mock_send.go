package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// mock数据测试发送消息
func mockSendHandler(c *gin.Context) {
	value := c.Query("data")

	strArray := strings.Split(string(value), ",")
	time, _ := strconv.ParseInt(strArray[3], 10, 32)
	o, _ := strconv.ParseFloat(strArray[4], 64)
	h, _ := strconv.ParseFloat(strArray[5], 64)
	l, _ := strconv.ParseFloat(strArray[6], 64)
	close, _ := strconv.ParseFloat(strArray[7], 64)
	response := &ResponseMessage{
		Time:   int32(time),
		Symbol: strArray[2],
		Data:   []float64{o, h, l, close},
	}
	WsManager.send(response)
}
