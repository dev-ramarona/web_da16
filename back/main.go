package main

import (
	fncGlobal "back/global/function"
	fncJeddah "back/jeddah/function"
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Initialize MongoDB connection
	contxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	portdb := options.Client().ApplyURI(fncGlobal.Urlmgo)
	fncGlobal.Client, _ = mongo.Connect(contxt, portdb)

	// Framework Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: fncGlobal.Ipalow,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"Authorization", "Cookie", "Content-Disposition"},
		AllowCredentials: true,
	}))

	// Handle global
	r.GET("/global/status", fncGlobal.FncGlobalMainprStatus)

	// Handle web link API all user
	r.POST("/allusr/loginx", fncGlobal.FncGlobalAllusrLoginx)
	r.GET("/allusr/tokenx", fncGlobal.FncGlobalAllusrTokenx)
	r.GET("/allusr/logout", fncGlobal.FncGlobalAllusrLogout)
	r.GET("/allusr/applst", fncGlobal.FncGlobalAllusrApplst)

	// // Handle web link API jeddah
	r.POST("/jeddah/prcess", fncJeddah.FncJeddahPrcessMainpg)
	r.GET("/jeddah/actlog/getall", fncJeddah.FncJeddahActlogGetall)
	r.GET("/jeddah/flnbfl/tmplte", fncJeddah.FncJeddahFlnbflTmplte)
	r.POST("/jeddah/flnbfl/upload/:upldby", fncJeddah.FncJeddahFlnbflUpload)
	r.POST("/jeddah/flnbfl/update", fncJeddah.FncJeddahFlnbflUpdate)
	r.POST("/jeddah/agtnme/nullvl", fncJeddah.FncJeddahAgtnmeNullnm)
	r.POST("/jeddah/agtnme/update", fncJeddah.FncJeddahAgtnmeUpdate)
	r.GET("/jeddah/agtnme/search/:newidn/:newdtl", fncJeddah.FncJeddahAgtnmeSearch)
	r.POST("/jeddah/pnrsmr/getall", fncJeddah.FncJeddahPnrsmrFrntnd)
	r.POST("/jeddah/pnrsmr/getall/:downld", fncJeddah.FncJeddahPnrsmrFrntnd)
	r.POST("/jeddah/pnrdtl/getall", fncJeddah.FncJeddahPnrdtlGetall)
	r.POST("/jeddah/pnrdtl/getall/:downld", fncJeddah.FncJeddahPnrdtlGetall)
	r.POST("/jeddah/flnsmr/getall", fncJeddah.FncJeddahFlnsmrGetall)
	r.POST("/jeddah/flnsmr/getall/:downld", fncJeddah.FncJeddahFlnsmrGetall)
	r.GET("/jeddah/rtlsrs/tmplte", fncJeddah.FncJeddahRtlsrsTmplte)
	r.POST("/jeddah/rtlsrs/update", fncJeddah.FncJeddahRtlsrsUpdate)
	r.POST("/jeddah/rtlsrs/upload/:upldby", fncJeddah.FncJeddahRtlsrsUpload)

	// Run server
	r.Run("0.0.0.0:" + fncGlobal.Ptgolg)
}
