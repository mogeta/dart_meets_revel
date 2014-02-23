package controllers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"github.com/robfig/revel"
	"voice/app/models"
)

type App struct {
	*revel.Controller
}

type Statuses struct {
	*Entities
}

type Entities struct {
	*Urls
}

type Urls struct {
	Expanded_url string
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}

func (c App) List() revel.Result {
	var r = models.Myredis{}
	r.Connect()
  //return json.Marshal(r.Get())
  return c.RenderJson(r.Get())
}

func (c App) Twitter() revel.Result {

	//revel.INFO.Println("test")
	//result, _ := ioutil.ReadAll(resp.Body)
	//revel.INFO.Println(string(result))
	//revel.INFO.Println(resp.Body)
	//revel.INFO.Println(mentions)

	//redisに投入
	var r = models.Myredis{}
	r.Connect()

	rankData, err := json.Marshal(r.Get())
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("%s", rankData)
	return c.Render(rankData)
}

func (c App) D3() revel.Result {
	//redisに投入
	var r = models.Myredis{}
	r.Connect()

	rankData, err := json.Marshal(r.Get())
	if err != nil {
		fmt.Println("err")
	}

	result := string(rankData[0:])
	return c.Render(result)
}

func (c App) Index() revel.Result {
	return c.Render()
}
