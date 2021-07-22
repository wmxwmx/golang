package worker

import (
	"fmt"
	"test/data/data"
)

type Worker struct {
	In  *chan data.ListViewData
	Out *chan data.ListViewData
}

func (w Worker) Prepare() {
	go func() {
		for {
			requestData := <-*w.In
			resultData, err := fetchDetail(&requestData)
			fmt.Println(resultData.Author)
			if err != nil {
				*w.Out <- *resultData
				continue
			}
			*w.Out <- *resultData
		}
	}()

}
