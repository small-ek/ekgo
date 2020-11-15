package testJob

import (
	"fmt"
	"github.com/small-ek/ginp/queue"
	"time"
)

type Test struct {
	Time   time.Time
	Key    string
	Params []interface{}
}

//testJob.Push(&testJob.Test{
//Time:   time.Now().Add(time.Second * 10),
//Key:    "testJob",
//Params: []interface{}{1, 2, 3},
//})
func Push(p *Test) {
	go func() {
		var dm = queue.NewDelayMessage()
		//添加任务
		dm.AddTask(p.Time, p.Key, func(args ...interface{}) {
			fmt.Println(args...)
			dm.Stop()
		}, p.Params)
		//什么时候结束
		//time.AfterFunc(time.Second*30, func() {
		//	dm.Stop()
		//})
		dm.Start()
	}()
}
