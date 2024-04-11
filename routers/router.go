package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/YukiJuda111/go-gin-blog/pkg/setting"
	"github.com/YukiJuda111/go-gin-blog/routers/api"
	v1 "github.com/YukiJuda111/go-gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	/*
		返回Gin的type Engine struct{...}
		里面包含RouterGroup
		相当于创建一个路由Handlers
		可以后期绑定各类的路由规则和函数、中间件等
	*/
	r := gin.Default()

	gin.SetMode(setting.ServerSetting.RunMode)

	// /*
	// 	创建不同的HTTP方法绑定到Handlers中
	// 	也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
	// */
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })
	// /*
	// 	Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等
	// 	在gin中包含大量Context的方法
	// 	例如DefaultQuery、Query、DefaultPostForm、PostForm等等
	// */

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
