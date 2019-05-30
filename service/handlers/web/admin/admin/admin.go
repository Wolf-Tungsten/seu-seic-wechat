package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"wechat-bind/models"
	"wechat-bind/pkg"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/admin")
	route.GET("", GET)
	route.POST("", POST)

}

func GET(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	if user.AdminLevel >= 0 {
		pkg.Return(ctx, 403, "无权访问")
		return
	}
	db := ctx.MustGet("db").(*mongo.Database)
	cursor, err := db.Collection("user").Find(ctx, bson.M{"adminLevel": bson.M{"$lte": -1}})
	var result []gin.H
	if err != nil {
		pkg.Return(ctx, 500, "数据库访问出错")
		log.Fatal(err)
		return
	}
	for cursor.Next(ctx) {
		var admin models.User
		_ = cursor.Decode(&admin)
		result = append(result, gin.H{"cardnum": admin.Cardnum, "name": admin.Name, "id": admin.Id, "level": admin.AdminLevel})
	}
	pkg.Return(ctx, 200, gin.H{"myLevel": user.AdminLevel, "list": result})
}

func POST(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	if user.AdminLevel >= 0 {
		pkg.Return(ctx, 403, "无权访问")
		return
	}

	db := ctx.MustGet("db").(*mongo.Database)
	requestBody := struct {
		Cardnum string `json:"cardnum"`
	}{}
	_ = ctx.Bind(&requestBody)

	var userToBeAdmin models.User
	_ = db.Collection("user").FindOne(ctx, bson.M{"cardnum": requestBody.Cardnum}).Decode(&userToBeAdmin)

	if userToBeAdmin.AdminLevel < 0 && userToBeAdmin.AdminLevel >= user.AdminLevel {
		pkg.Return(ctx, 400, "请勿重复设置管理员")
		return
	}

	// 权限让渡
	upsert, _ := db.Collection("user").UpdateOne(ctx, bson.M{"cardnum": requestBody.Cardnum}, bson.M{"$set": bson.M{"adminLevel": user.AdminLevel - 1}})

	switch upsert.ModifiedCount {
	case 0:
		{
			pkg.Return(ctx, 400, "用户不存在")
		}
	case 1:
		{
			pkg.Return(ctx, 200, "设置成功")
		}
	}
}

func DELETE(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	if user.AdminLevel >= 0 {
		pkg.Return(ctx, 403, "无权访问")
		return
	}

	db := ctx.MustGet("db").(*mongo.Database)
	cardnum := ctx.Query("cardnum")

	var userToBeAdmin models.User
	_ = db.Collection("user").FindOne(ctx, bson.M{"cardnum": cardnum}).Decode(&userToBeAdmin)

	if userToBeAdmin.AdminLevel < 0 && userToBeAdmin.AdminLevel >= user.AdminLevel {
		pkg.Return(ctx, 400, "无权操作")
		return
	}

	// 权限撤销
	upsert, _ := db.Collection("user").UpdateOne(ctx, bson.M{"cardnum": cardnum}, bson.M{"$set": bson.M{"adminLevel": 0}})

	switch upsert.ModifiedCount {
	case 0:
		{
			pkg.Return(ctx, 400, "用户不存在")
		}
	case 1:
		{
			pkg.Return(ctx, 200, "取消成功")
		}
	}
}
