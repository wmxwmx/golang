package main

import (
	"test/data/engin"
	"test/data/scheduler"
	"time"
)

func main() {

	engin.Engine{
		Scheduler: scheduler.DetailScheduler{},
	}.Run()
	time.Sleep(300 * time.Second)

}
