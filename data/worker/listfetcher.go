package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"test/data/data"
	"test/global"
	"time"
)

func FetchList() ([]data.ListViewData, error) {
	paramsData, _ := json.Marshal(getBitcoinListNetParams())
	request, err := http.NewRequest(http.MethodPost, global.ListBaseUrl, bytes.NewReader(paramsData))
	if err != nil {
		return nil, err
	}
	request.Header.Add("content-type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, err
	}
	jsonStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	netData, err := JsonToNetData(jsonStr)
	if err != nil {
		panic(err)
	}
	var _ = request.Close
	return NetToListData(netData), nil

}

//https://cointelegraph.com/markets
func getMarketListNetParams() map[string]interface{} {
	param := map[string]interface {
	}{
		"operationName": "TagPageQuery",
		"variables": map[string]interface {
		}{
			"slug":          "markets",
			"order":         "postPublishedTime",
			"offset":        0,
			"length":        20,
			"short":         "en",
			"cacheTimeInMS": 300000,
		},
		"query": "query TagPageQuery($short: String, $slug: String!, $order: String, $offset: Int!, $length: Int!) {\n locale(short: $short) {\n tag(slug: $slug) {\n cacheKey\n id\n slug\n avatar\n createdAt\n updatedAt\n redirectRelativeUrl\n alternates {\n cacheKey\n short\n domain\n id\n code\n __typename\n }\n tagTranslates {\n cacheKey\n id\n title\n metaTitle\n pageTitle\n description\n metaDescription\n keywords\n __typename\n }\n posts(order: $order, offset: $offset, length: $length) {\n data {\n cacheKey\n id\n slug\n views\n postTranslate {\n cacheKey\n id\n title\n avatar\n published\n publishedHumanFormat\n leadText\n __typename\n }\n category {\n cacheKey\n id\n __typename\n }\n author {\n cacheKey\n id\n slug\n authorTranslates {\n cacheKey\n id\n name\n __typename\n }\n __typename\n }\n postBadge {\n cacheKey\n id\n label\n postBadgeTranslates {\n cacheKey\n id\n title\n __typename\n }\n __typename\n }\n showShares\n showStats\n __typename\n }\n postsCount\n __typename\n }\n __typename\n }\n __typename\n }\n}\n",
	}
	return param
}

//https://cointelegraph.com/tags/bitcoin
func getBitcoinListNetParams() map[string]interface{} {
	param := map[string]interface {
	}{
		"operationName": "TagPageQuery",
		"variables": map[string]interface {
		}{
			"slug":          "bitcoin",
			"order":         "postPublishedTime",
			"offset":        0,
			"length":        122,
			"short":         "en",
			"cacheTimeInMS": 300000,
		},
		"query": "query TagPageQuery($short: String, $slug: String!, $order: String, $offset: Int!, $length: Int!) {\n locale(short: $short) {\n tag(slug: $slug) {\n cacheKey\n id\n slug\n avatar\n createdAt\n updatedAt\n redirectRelativeUrl\n alternates {\n cacheKey\n short\n domain\n id\n code\n __typename\n }\n tagTranslates {\n cacheKey\n id\n title\n metaTitle\n pageTitle\n description\n metaDescription\n keywords\n __typename\n }\n posts(order: $order, offset: $offset, length: $length) {\n data {\n cacheKey\n id\n slug\n views\n postTranslate {\n cacheKey\n id\n title\n avatar\n published\n publishedHumanFormat\n leadText\n __typename\n }\n category {\n cacheKey\n id\n __typename\n }\n author {\n cacheKey\n id\n slug\n authorTranslates {\n cacheKey\n id\n name\n __typename\n }\n __typename\n }\n postBadge {\n cacheKey\n id\n label\n postBadgeTranslates {\n cacheKey\n id\n title\n __typename\n }\n __typename\n }\n showShares\n showStats\n __typename\n }\n postsCount\n __typename\n }\n __typename\n }\n __typename\n }\n}\n",
	}
	return param
}

//https://cointelegraph.com/tags/ethereum
func getEthereumListNetParams() map[string]interface{} {
	param := map[string]interface {
	}{
		"operationName": "TagPageQuery",
		"variables": map[string]interface {
		}{
			"slug":          "ethereum",
			"order":         "postPublishedTime",
			"offset":        0,
			"length":        20,
			"short":         "en",
			"cacheTimeInMS": 300000,
		},
		"query": "query TagPageQuery($short: String, $slug: String!, $order: String, $offset: Int!, $length: Int!) {\n locale(short: $short) {\n tag(slug: $slug) {\n cacheKey\n id\n slug\n avatar\n createdAt\n updatedAt\n redirectRelativeUrl\n alternates {\n cacheKey\n short\n domain\n id\n code\n __typename\n }\n tagTranslates {\n cacheKey\n id\n title\n metaTitle\n pageTitle\n description\n metaDescription\n keywords\n __typename\n }\n posts(order: $order, offset: $offset, length: $length) {\n data {\n cacheKey\n id\n slug\n views\n postTranslate {\n cacheKey\n id\n title\n avatar\n published\n publishedHumanFormat\n leadText\n __typename\n }\n category {\n cacheKey\n id\n __typename\n }\n author {\n cacheKey\n id\n slug\n authorTranslates {\n cacheKey\n id\n name\n __typename\n }\n __typename\n }\n postBadge {\n cacheKey\n id\n label\n postBadgeTranslates {\n cacheKey\n id\n title\n __typename\n }\n __typename\n }\n showShares\n showStats\n __typename\n }\n postsCount\n __typename\n }\n __typename\n }\n __typename\n }\n}\n",
	}
	return param
}

// JsonToNetData json???netData
func JsonToNetData(jsonStr []byte) (data.ListNetData, error) {
	data1 := data.ListNetData{}
	err := json.Unmarshal(jsonStr, &data1)
	TranslateErr(err)
	return data1, err
}

func TranslateErr(err error) {
	if err != nil {
		fmt.Print("ErrHappen\n", err)
		panic(err)
	}
}

// NetToListData net?????????view??????
func NetToListData(netData data.ListNetData) []data.ListViewData {

	listData := make([]data.ListViewData, 0)
	for _, netPostBean := range netData.Data.Locale.Tag.Posts.Data {
		listData = append(listData, data.ListViewData{
			//????????????
			Title: netPostBean.PostTranslate.Title,
			//??????????????????
			LeadText: netPostBean.PostTranslate.LeadText,
			//????????????
			Image: netPostBean.PostTranslate.Avatar,
			//????????????
			GetTime: time.Now(),
			//??????????????????
			PublishTime: netPostBean.PostTranslate.Published,
			//??????????????????????????????key
			Slug: netPostBean.Slug,
			//??????????????????
			Url: strings.Join([]string{global.DetailBaseUrl, netPostBean.Slug}, ""),
			//??????????????????
			Author: func() string {
				name := ""
				//????????????????????????????????????????????????????????????????????????
				for _, da := range netPostBean.Author.AuthorTranslates {
					return da.Name
				}
				return name
			}(),
		})

	}
	return listData

}
