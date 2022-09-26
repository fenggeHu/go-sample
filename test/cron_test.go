package test

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
func TestCron(t *testing.T) {
	c := newWithSeconds()
	cronExpression := "10 */1 * * * ?" //每分钟的10秒运行一次
	c.AddFunc(cronExpression, func() {
		fmt.Println("cron running: ", time.Now().Unix())
		time.Sleep(2 * time.Minute) // 休眠2分钟 - 没有阻塞下一次执行
	})
	c.Start()

	select {}
}
