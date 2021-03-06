package engin

import (
	"test/data/data"
	"test/data/scheduler"
	"test/data/worker"
)

type Engine struct {
	Scheduler      scheduler.DetailScheduler
	GoroutineCount int //开启携程个数
}

type Scheduler interface {
	Submit(data *data.ListViewData)
	ConfigChan(ch *chan data.ListViewData)
}

func (e Engine) Run() {
	ListData := make([]data.ListViewData, 0)

	listViewData, err := worker.FetchList()
	if err != nil {
		panic(err)
		return
	}
	in := make(chan data.ListViewData)
	out := make(chan data.ListViewData)
	for i := 0; i < e.GoroutineCount; i++ {
		creatWorker(&in, &out)
	}
	//接收每结果
	go func() {
		for i := 1; ; i++ {
			result := <-out
			ListData = append(ListData, result)
			if i == len(listViewData) {
				break
			}
		}
		//fmt.Println("\n", ListData[0].Html)
	}()
	e.Scheduler.ConfigChan(&in)
	//发送任务数据给协调器，进行过滤
	for _, data := range listViewData {
		e.Scheduler.Submit(data)
	}

}

func creatWorker(in *chan data.ListViewData, out *chan data.ListViewData) {
	worker.Worker{
		In:  in,
		Out: out,
	}.Prepare()
}
