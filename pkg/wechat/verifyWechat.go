package wechat

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-common/secret"
	"sort"
	"strings"
)

func CheckWechatSignature(ctx *gin.Context) bool{

	signature := ctx.Query("signature")
	timestamp := ctx.Query("timestamp")
	nonce := ctx.Query("nonce")
	signatureList := [] string {timestamp, nonce, secret.WechatToken}
	sort.Strings(signatureList)
	signatureCheck := strings.Join(signatureList, "")
	h := sha1.New()
	h.Write([]byte(signatureCheck))
	signatureCheckSha1 := h.Sum(nil)
	signatureCheckSha1Str := fmt.Sprintf("%x", signatureCheckSha1)
	return signatureCheckSha1Str == signature
}

