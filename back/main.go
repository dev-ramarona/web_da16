package main

import (
	fncGlobal "back/global"
	"context"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Initialize MongoDB connection
	contxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fncGlobal.FncGlobalMainprLoadnv("../.env")
	defer cancel()
	portdb := options.Client().ApplyURI(os.Getenv("URI_MONGOS"))
	fncGlobal.Client, _ = mongo.Connect(contxt, portdb)

	// Framework Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     strings.Split(os.Getenv("IPV_ALLOWD"), "|"),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"Authorization", "Cookie", "Content-Disposition"},
	}))

	// // Handle global
	// r.GET("/global/status", fncGlobal.FncGlobalApisbrStatus)

	// // Handle web link API all user
	// r.POST("/allusr/loginx", fncGlobal.FncGlobalAllusrLoginx)
	// r.GET("/allusr/tokenx", fncGlobal.FncGlobalAllusrTokenx)
	// r.GET("/allusr/logout", fncGlobal.FncGlobalAllusrLogout)
	// r.GET("/allusr/applst", fncGlobal.FncGlobalAllusrApplst)

	// // Handle web link API jeddah
	// // r.POST("/jeddah/addfln", fnc_jeddah.FncJeddahAddflnTodtbs)
	// r.POST("/jeddah/prcess", fnc_jeddah.FncJeddahPrcessMainpg)
	// r.POST("/jeddah/agtnul", fnc_jeddah.FncJeddahAgtnmeNullnm)
	// r.GET("/jeddah/logact", fnc_jeddah.FncJeddahLogactGetall)
	// r.POST("/jeddah/agtupd", fnc_jeddah.FncJeddahAgtnmeUpdate)
	// r.POST("/jeddah/pnrdtl", fnc_jeddah.FncJeddahDtlpnrFrntnd)
	// r.POST("/jeddah/pnrdtl/:downld", fnc_jeddah.FncJeddahDtlpnrFrntnd)
	// r.POST("/jeddah/pnrsmr", fnc_jeddah.FncJeddahSmrpnrFrntnd)
	// r.POST("/jeddah/pnrsmr/:downld", fnc_jeddah.FncJeddahSmrpnrFrntnd)
	// r.POST("/jeddah/flnsmr", fnc_jeddah.FncJeddahSmrflnFrntnd)
	// r.POST("/jeddah/flnsmr/:downld", fnc_jeddah.FncJeddahSmrflnFrntnd)
	// r.GET("/jeddah/agtsrc/:newidn/:newdtl", fnc_jeddah.FncJeddahAgtnmeAgtsrc)
	// r.POST("/jeddah/rtlsrs/update", fnc_jeddah.FncJeddahRtlsrsUpdate)
	// r.POST("/jeddah/rtlsrs/upload/:upldby", fnc_jeddah.FncJeddahRtlsrsUpload)
	// r.GET("/jeddah/rtlsrs/status", fnc_jeddah.FncJeddahRtlsrsStatus)

	// // Handle web link API File PPN Lookup
	// r.POST("/ppnlkp/prcess", fnc_ppnlkp.FncPpnlkpPrcessMainpg)
	// r.GET("/ppnlkp/agrgte", fnc_ppnlkp.FncPpnlkpPrcessAgrgte)

	r.Run(os.Getenv("IPV_GOLANG"))
}
