package jobs

import (
	"fmt"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/jobs/app/jobs"
	"voice/app/models"
)

type twitter_url struct{}

func init() {
	fmt.Println("runnning")
	revel.OnAppStart(func() {
		fmt.Println("start")
		//一分間隔
		//jobs.Schedule("0 * * * * *", jobs.Func(func() { fmt.Println("1minute") }))
		//jobs.Schedule("0 * * * * *", twitter_url{})
		jobs.Schedule("@every 15s", twitter_url{})
	})
}

func (t twitter_url) Run() {
	fmt.Println("search start")
	var m map[string]int
	m = make(map[string]int)
	d := Search()

	//todo
	for _, v := range d {
		for _, vv := range v.Entities.Urls {
			m[vv.Expanded_url]++
			fmt.Println("url:%v", vv.Expanded_url)
			fmt.Printf("url:%v\n", vv.Expanded_url)
			fmt.Println(m[vv.Expanded_url])
		}
	}

	//redisに投入
	var r = models.Myredis{}
	r.Connect()
	r.Zincr(m)

	fmt.Println("search end")
}

//todo : websocket
//
