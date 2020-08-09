package jianshu_image_uploader

import (
	"io/ioutil"
	http2 "net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddliu/go-httpclient"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

var httpClient *httpclient.HttpClient

func init() {
	httpClient = httpclient.NewHttpClient().Defaults(httpclient.Map{
		httpclient.OPT_REFERER:   "https://www.jianshu.com/writer",
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64. 0) Gecko/20100101 Firefox/64.0",
		//httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Windows NT 6.1; rv:24.0) Gecko/20100101 Firefox/24.0",
		httpclient.OPT_UNSAFE_TLS: true,
	})
}

func GetToken(cookies []*http2.Cookie, fileName string) (token string, key string) {
	data := map[string]string{"filename": fileName}

	url := "https://www.jianshu.com/upload_images/token.json"
	res, _ := httpClient.Begin().WithCookie(cookies...).Get(url, data)
	body, _ := res.ToString()
	json := gjson.Parse(body)
	return json.Get("token").String(), json.Get("key").String()
}

func UploadImg(filePath string, token string, key string) (body string) {
	if strings.HasPrefix(filePath, "http://") == true || strings.HasPrefix(filePath, "https://") == true {
		color.Green("从远程下载图片")
		res, _ := httpClient.Begin().Get(filePath)
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		filePath = getImageName(filePath)

		file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm|os.ModeAppend)
		file.Write(bodyBytes)
		defer file.Close()
		defer os.Remove(filePath)
	}
	postData := map[string]string{}
	postData["token"] = token
	postData["key"] = key
	postData["@file"] = filePath
	res, _ := httpClient.Begin().Post("http://upload.qiniup.com", postData)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json := gjson.ParseBytes(bodyBytes)
	return json.Get("url").String()
}

func getImageName(imagePath string) string {
	return filepath.Base(imagePath)
}
