package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	cfg "gospider/config"

	"github.com/gocolly/colly"
)

func getApi(urlpath string) (string, error) {
	req, err := http.NewRequest("GET", urlpath,nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: cfg.CfgInfo.CookieUName, Value: cfg.CfgInfo.CookieUPass})
	req.Header.Add("iottoken",cfg.CfgInfo.AuthIOT);

	// Set the auth for the request.
	//req.SetBasicAuth(cfg.CfgInfo.AuthUser, cfg.CfgInfo.AuthPass)

	resp,_ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body),nil
}

func verify() {
	var data interface{}

	for i := 0 ; i < len(cfg.CfgInfo.ApiList) ; i ++ {
		var urlpath string = cfg.CfgInfo.ApiList[i]
		resp,err	:= getApi(urlpath)
		if err != nil {
			log.Printf("get url %s failed %t\n",urlpath,err)
			continue
		}
		err = json.Unmarshal([]byte(resp), &data)
		// read any json : https://github.com/mushuanli/gowebserver/blob/main/config/config.go
		// or define struct to Unmarshal
	}

	userInfo,err := getApi(cfg.GetApi("ApiUsrInfo"))
	if err != nil {
		log.Printf("get User info failed %t\n",err)
	}
	err = json.Unmarshal([]byte(userInfo), &data)
}

func collect() {
	c := colly.NewCollector()

    c.OnHTML("a", func(e *colly.HTMLElement) {
        e.Request.Visit(e.Attr("href"))
    })

    c.OnRequest(func(r *colly.Request) {
		cookies := *&http.Cookie{
			{
				Name: ""
				Name: cfg.CfgInfo.CookieUName, Value: cfg.CfgInfo.CookieUPass
			}
		}
        fmt.Println("Visiting", r.URL)
    })

    c.Visit("https://wls.cogiot.net/weiduan/#/app/weiduan/monitor/device")
}

func main() {
	cfg.Init("./config.json")
	log.Println("app start");
	collect()
	verify()
}
