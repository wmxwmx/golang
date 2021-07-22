package data

import "time"

type ListViewData struct {
	GetTime     time.Time //抓取时间
	Title       string    //标题
	LeadText    string    //引导文本
	PublishTime time.Time //发布时间时间
	Author      string    //作者
	Image       string    //图片
	Slug        string    //网络请求key
	Url         string    //原详情页url
	Html        string //详情html

}
