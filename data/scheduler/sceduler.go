package scheduler

import (
	"test/data/data"
)

type DetailScheduler struct {
	sendToWorkChan *chan data.ListViewData
}

func (s *DetailScheduler) Submit(data data.ListViewData) {
	//这里可以添加判断，对需要再抓取的数据不再抛入chan
	if GoOn(&data) {
		*s.sendToWorkChan <- data
	}

}

func (s *DetailScheduler) ConfigChan(ch *chan data.ListViewData) {
	s.sendToWorkChan = ch
}

// GoOn
//是否需要抓取次新闻/**/
func GoOn(data *data.ListViewData) bool {
	return true
}
