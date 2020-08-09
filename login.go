package jianshu_image_uploader

import (
	"encoding/json"
	"io/ioutil"
	http2 "net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func SaveCookie(newCookie string, cookiePath string) bool {
	newCookieArray := strings.Split(newCookie, "; ")
	httpCookie := make(map[string]string)
	if len(newCookieArray) > 0 {
		for _, singleCookie := range newCookieArray {
			line := strings.Split(singleCookie, "=")
			httpCookie[line[0]] = line[1]
		}
		httpCookieBytes, _ := json.Marshal(httpCookie)
		file, err := os.OpenFile(cookiePath, os.O_RDWR|os.O_CREATE, os.ModePerm|os.ModeAppend)
		if err != nil {
			color.Red(err.Error())
		}
		file.Write(httpCookieBytes)
		defer file.Close()
		return true
	} else {
		return false
	}
}

func ParserCookie(cookiePath string) ([]*http2.Cookie, error) {

	cookieFile, err := os.OpenFile(cookiePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		color.Red("错误：" + err.Error())
		os.Exit(0)
	}
	cookieContent, err := ioutil.ReadAll(cookieFile)

	if len(cookieContent) == 0 {
		return nil, nil
	}
	cookieContentJson := gjson.ParseBytes(cookieContent).Value().(map[string]interface{})
	if len(cookieContentJson) == 0 {
		return nil, nil
	}
	cookies := make([]*http2.Cookie, 0)
	for key, val := range cookieContentJson {
		cookies = append(cookies, &http2.Cookie{
			Name:  key,
			Value: val.(string),
		})
	}
	color.Green("解析cookie成功")
	return cookies, err
}

func DeleteCookie(cookiePath string) {
	os.Remove(cookiePath)
}
