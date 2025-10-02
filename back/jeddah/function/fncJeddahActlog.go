package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Agent name not complete format slice data from database
func FncJeddahActlogGetall(c *gin.Context) {

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_actlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []mdlJeddah.MdlJeddahActlogDtbase{}
	for datarw.Next(contxt) {
		var object mdlJeddah.MdlJeddahActlogDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
		}
	}

	// Send token to frontend
	c.JSON(http.StatusOK, slices)
}

// Get Agent name not complete format slice data from database
func FncJeddahActlogLstdta() string {

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_actlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get Highest dateup
	var slices mdlJeddah.MdlJeddahActlogDtbase
	option := options.FindOne().SetSort(bson.D{{Key: "dateup", Value: -1}})
	err := tablex.FindOne(contxt, bson.M{}, option).Decode(&slices)
	if err != nil {
		nowTimenb := time.Now().AddDate(0, 0, -1).Format("060102")
		return nowTimenb
	}

	// Send token to frontend
	prvDateup, _ := strconv.Atoi(time.Now().AddDate(0, 0, -1).Format("060102"))
	if slices.Dateup >= int32(prvDateup) {
		return strconv.Itoa(int(prvDateup))
	}
	return strconv.Itoa(int(slices.Dateup))
}
