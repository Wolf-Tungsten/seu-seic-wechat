# 小猴偷米 Golang 后端模板项目

为小猴偷米非核心业务产品后端开发提供模板项目，请不要直接修改本项目，而应在 fork 之后修改以适应业务逻辑。

## 基本概况

* 使用 Go 语言进行开发
* 采用 [Gin Web Framework](https://github.com/gin-gonic/gin) 作为基础框架
* 使用 MongoDB 数据库，数据库驱动使用 MongoDB 官方提供的 [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
* 尽可能遵循 RESTful 风格

## 路由结构设计

* URL 路径结构和包结构对应设计，定义 handlers 包为顶级路径
* 在 `handlers/handlers.go` 中注册顶层路由，其他路由注册在其上一层的路由中
* 每个包文件夹下建立一个和包同名的源码文件，在这个文件中绑定该路径的路由和下一级的路由

```$go
func Handler(router *gin.RouterGroup){

	route := router.Group("/apiGroup1")
	route.GET("", GET)
	route.POST("", POST)
	// 注册下一层 handler
	api1.Handler(router)

}

func GET(ctx *gin.Context){
	// 处理 GET - /apiGroup1 请求
}

func POST(ctx *gin.Context){
	// 处理 POST - /apiGroup1 请求
}
```


