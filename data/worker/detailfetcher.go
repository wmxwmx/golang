package worker

import (
	"io/ioutil"
	"net/http"
	"strings"
	"test/data/data"
	"test/global"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)
func fetchDetail(notPrefect *data.ListViewData) (*data.ListViewData, error) {

	<-rateLimiter
	url := strings.Join([]string{global.DetailBaseUrl,notPrefect.Slug},"")
	request, err := http.NewRequest(http.MethodGet,url, nil)
	if err != nil {
		return nil,err
	}
	defer func() {
		var _ = request.Close
	}()
	request.Header.Add("authority", "cointelegraph.com")
	response, err := http.DefaultClient.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		return notPrefect,err
	}
	html,err :=  ioutil.ReadAll(response.Body)
	if err != nil{
		return notPrefect,err
	}
	notPrefect.Html = parse(html)

	return notPrefect,nil

}
