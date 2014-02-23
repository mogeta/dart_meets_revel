package tests

import (
	"github.com/robfig/revel"
	"voice/app/jobs"
)

type TwitTest struct {
	revel.TestSuite
}

func (t TwitTest) Before() {
	println("Set up")
}

func (t TwitTest) TestThatIndexPageWorks() {
	jobs.Search()
}

func (t TwitTest) After() {
	println("Tear down")
}
