package worker

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func parse(htmlData []byte) string {

	root, _ := htmlquery.Parse(bytes.NewReader(htmlData))
	//查找无用节点并删除：
	useLessNode2 := htmlquery.FindOne(root, "//div[@class='header-zone layout__header']")
	removeNode(useLessNode2)
	//去除底部无用节点
	removeNode(htmlquery.FindOne(root, "//div[@class='post__separator-line']"))
	removeNode(htmlquery.FindOne(root, "//div[@class='tags-list post__block post__block_tags']"))
	removeNode(htmlquery.FindOne(root, "//div[@class='related-list']"))
	/////去除查看更多节点
	removeNode(htmlquery.FindOne(root, "//div[@class='posts-listing__more-wrp']"))
	//去除分享部分
	shareNodes := htmlquery.FindOne(root, "//div[@class='post-page']")
	if shareNodes != nil {
		removeNode(shareNodes.FirstChild.NextSibling)
	}
	//去除头部作者部分
	removeNode(htmlquery.FindOne(root, "//div[@class='post-meta post__block post__block_post-meta']"))
	//去除底部推荐
	removeNode(htmlquery.FindOne(root, "//div[@class='post-page__sidebar-col']"))

	//去除搜索部分
	removeNode(htmlquery.FindOne(root, "//div[@class='header-mobile-search layout__mobile-search']"))

	//去除音频播放部分
	//removeNode(htmlquery.FindOne(root, "//div[@class='post-actions post__block post__block_post-actions']"))
	//removeNode(htmlquery.FindOne(root, "//div[@class='post-cover post__block post__block_cover']"))

	//去除免责声明
	removeNode(htmlquery.FindOne(root, "//p[@class='post-content__disclaimer']"))
	removeNode(htmlquery.FindOne(root, "//div[@class='post__separator-line']"))

	//处理音视频的播放和显示
	body := htmlquery.FindOne(root, "/html/body")
	body.AppendChild(&html.Node{
		Data:     "script",
		Type:     html.ElementNode,
		DataAtom: atom.Script,
		FirstChild: &html.Node{
			Type:     html.TextNode,
			DataAtom: 0,
			Data:     "let audio=document.getElementsByTagName(\"audio\")[0];let mediaEle=document.getElementsByClassName(\"post-audio-player__meta\")[0];audio.addEventListener(\"loadedmetadata\",function(){setText()})\n    function formatTime(d){let min=Math.floor(d/60)\n        let second=Math.floor(d-min*60)\n        return\"\"+min+\":\"+second}\n    function setText(){mediaEle.innerText=formatTime(audio.duration-audio.currentTime)}\n    document.getElementsByClassName(\"post-audio-player__play-btn\")[0].onclick=function(){let playerIcon=document.getElementsByClassName(\"post-audio-player__player-icon\")[0]\n        if(audio.paused){audio.play()\n            playerIcon.className=\"btn__icon post-audio-player__player-icon post-audio-player__pause-icon\"}else{audio.pause()\n            playerIcon.className=\"btn__icon post-audio-player__player-icon post-audio-player__play-icon\"}\n        audio.addEventListener(\"timeupdate\",function(t){setText()})}",
		},
	})
	return htmlquery.OutputHTML(root, true)
}

func removeNode(childrenNodes *html.Node) {
	if childrenNodes != nil {
		childrenNodes.Parent.RemoveChild(childrenNodes)
	}
}

//写文件
//func w(str string) {
//	var d1 = []byte(str)
//	err2 := ioutil.WriteFile("./a.html", d1, 0666)
//	fmt.Println(err2)
//}
