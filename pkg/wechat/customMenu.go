package wechat

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wechat-bind/database"
	"wechat-bind/secret"
)

func CustomMenu() {
	ctx := gin.Context{}
	ctx.Set("db", database.Client.Database(secret.DatabaseName))
	file, _ := ioutil.ReadFile(os.Getenv("WECHAT_MENU"))
	accessToken := GetAccessToken(&ctx)
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", accessToken), "applicant/json", bytes.NewReader(file))
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))
}
