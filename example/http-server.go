package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/fatih/color"
	//_ "gopkg.in/gographics/imagick.v2/imagick"

	"github.com/gin-gonic/gin"

	jianshu "github.com/lucklrj/jianshu_image_uploader"
)

type JsonResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    ImageData `json:"data"`
}

type ImageData struct {
	Url string `json:"url"`
}

type SaveData struct {
	Name    string   `json:"name"`
	ImgList []string `json:imgList`
}

var cookie string

func init() {
	flag.StringVar(&cookie, "cookie", "", "保存新的cookie")
}
func main() {
	cookiePath := "cookie.txt"

	flag.Parse()
	if cookie != "" {
		saveResult := jianshu.SaveCookie(cookie, cookiePath)
		if saveResult == false {
			color.Red("保存cookie出错")
			os.Exit(0)
		}
	}

	httpCookie, _ := jianshu.ParserCookie(cookiePath)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")

		_, filename, _, ok := runtime.Caller(1)
		var cwdPath string
		if ok {
			cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
		} else {
			cwdPath = "./"
		}
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, cwdPath+file.Filename)

		token, key := jianshu.GetToken(httpCookie, "test.jpg")
		if token == "" {
			color.Red("cookie已失效")
			os.Exit(0)
		}
		remoteUrl := jianshu.UploadImg(cwdPath+file.Filename, token, key)
		os.Remove(cwdPath + file.Filename)

		//c.String(http.StatusOK, remoteUrl)
		r := JsonResponse{}
		img := ImageData{Url: remoteUrl}
		r.Code = 200
		r.Message = "ok"
		r.Data = img
		c.JSON(200, r)
	})
	// post
	router.POST("/save", func(c *gin.Context) {
		var reqInfo SaveData
		err := c.BindJSON(&reqInfo)
		if err != nil {
			c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
			return
		} else {
			fmt.Println(reqInfo.ImgList)
			fmt.Println(reqInfo.Name)
			c.JSON(http.StatusOK, gin.H{"name": reqInfo.Name, "img": reqInfo.ImgList})
		}

	})
	router.Run(":9999")

}
