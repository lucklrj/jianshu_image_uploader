package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	jianshu "github.com/lucklrj/jianshu_image_uploader"
)

func main() {
	cookiePath := "cookie.txt"
	cookie := "sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2213345379%22%2C%22first_id%22%3A%221700124e2227bf-0efdfc111e4ea4-4b516c-1296000-1700124e223b7b%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_utm_source%22%3A%22weibo%22%2C%22%24latest_utm_medium%22%3A%22writer_share%22%2C%22%24latest_utm_campaign%22%3A%22maleskine%22%2C%22%24latest_utm_content%22%3A%22note%22%2C%22%24latest_referrer_host%22%3A%22%22%7D%2C%22%24device_id%22%3A%221700124e2227bf-0efdfc111e4ea4-4b516c-1296000-1700124e223b7b%22%7D; __yadk_uid=yG49RRkoVOKVhBCUC2T2VB0FFjEmwrr4; Hm_lvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1596787417,1596810955,1596817269,1596817287; __gads=ID=00bd0cbd407b3016:T=1580567160:S=ALNI_MaDndu6rrKEkG874f_xhIOqHEZTRQ; _ga=GA1.2.1432392703.1581697984; _gid=GA1.2.637572874.1596732011; _m7e_session_core=482b09c7141f7b371d6d9f6070837679; web_login_version=MTU5NjgxOTgzNw%3D%3D--d3a27af0d00e9d2bf6b259114c5d98d1a64c7275; read_mode=day; default_font=font1; locale=zh-CN; Hm_lpvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1596819854; remember_user_token=W1sxMzM0NTM3OV0sIiQyYSQxMSRFMnR6c1JULkVHN2pwcEhZTjJadHAuIiwiMTU5NjgxOTgzNy4xMjgzNTE0Il0%3D--1d47d2ba522639492512de7ab299438f36776064; _gat=1"
	if cookie != "" {
		saveResult := jianshu.SaveCookie(cookie, cookiePath)
		if saveResult == false {
			color.Red("保存cookie出错")
			os.Exit(0)
		}
	}

	httpCookie, _ := jianshu.ParserCookie(cookiePath)

	token, key := jianshu.GetToken(httpCookie, "test.jpg")
	body := jianshu.UploadImg("http://pic-bucket.ws.126.net/photo/0001/2020-08-08/FJHQFD8P00AN0001NOS.jpg", token, key)
	fmt.Println(body)
}
