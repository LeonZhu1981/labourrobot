package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	apiUrl := "http://hr.51labour.com/data/responseData.ashx"

	data := url.Values{}
	data.Set("pgindex", "1")
	data.Add("pgsize", "10")
	data.Add("addrid", "0")
	data.Add("yearid", "0")
	data.Add("typeid", "13")

	/*Set proxy to debug http request.*/
	// proxyUrl, err := url.Parse("http://localhost:8888")
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	client := &http.Client{}
	r, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err.Error())
	}
	r.Header.Set("Host", "hr.51labour.com")
	r.Header.Set("Origin", "http://hr.51labour.com")
	r.Header.Set("Accept", "text/html, */*; q=0.01")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Set("X-Requested-With", "XMLHttpRequest")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36")
	r.Header.Set("Referer", "http://hr.51labour.com/data/list-0-24.html")
	r.Header.Set("Accept-Language", "en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4")
	r.Header.Set("Cookie", "51labour_login=uid=611854&username=leonzhu1981&email=Leon.X.Zhu@gmail.com&token=74ccd69a22c1caf1d0ed0a33884beba7; Hm_lvt_8fb5c50d02bb75f5b9361280133d7646=1393840036; Hm_lpvt_8fb5c50d02bb75f5b9361280133d7646=1393848678; dnt=userid=739131&password=rhc89TgCG6FNSHTEapBHnVr0ygcnIk7w&avatar=http%3a%2f%2fbbs.51labour.com%2favatars%2fupload%2f000%2f73%2f91%2f31_avatar_medium.jpg&tpp=0&ppp=0&invisible=0&referer=index.aspx&expires=30")

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err == nil {
		if resp.StatusCode == 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
		}
	} else {
		panic(err.Error())
	}
}
