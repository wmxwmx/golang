package worker

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parse(html []byte) string {

	root, _ := htmlquery.Parse(bytes.NewReader(html))
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
	//去除音频播放部分
	removeNode(htmlquery.FindOne(root, "//div[@class='post-actions post__block post__block_post-actions']"))
	removeNode(htmlquery.FindOne(root, "//div[@class='post-cover post__block post__block_cover']"))
	return htmlquery.OutputHTML(root, true)
}


func removeNode(childrenNodes *html.Node) () {
	if childrenNodes != nil {
		childrenNodes.Parent.RemoveChild(childrenNodes)
	}
}
