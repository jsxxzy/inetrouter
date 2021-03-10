package inetrouter

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsxxzy/inet"
)

var ginApp *gin.Engine = gin.Default()

//go:embed template/* assets/*
var staticFS embed.FS

// Loop run http server
func Loop() {

	addPingWare(ginApp)

	templ := template.Must(template.New("").ParseFS(staticFS, "template/*.html"))
	ginApp.SetHTMLTemplate(templ)

	ginApp.StaticFS("/public", http.FS(staticFS))

	ginApp.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "大专人雄起",
		})
	})

	ginApp.GET("/api/check", func(ctx *gin.Context) {
		var flag = inet.HasLogin()
		var i = "0"
		if flag {
			i = "1"
		}
		ctx.String(http.StatusOK, i)
	})

	ginApp.GET("/api/info", func (ctx *gin.Context) {
		var info, err = inet.QueryInfo()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, info)
	})

	ginApp.Run()
}

func addPingWare(app *gin.Engine) {
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
}
