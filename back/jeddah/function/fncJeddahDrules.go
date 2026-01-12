package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Sync map data LC and PUN prev day
func FncJeddahDrulesSlcobj() []mdlJeddah.MdlJeddahRulesjDtbase {

	// Inisialisasi variabel
	fnldta := []mdlJeddah.MdlJeddahRulesjDtbase{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_rulesj")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{},
		options.Find().SetSort(bson.D{{Key: "rllevl", Value: 1}}))
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var nowDrules mdlJeddah.MdlJeddahRulesjDtbase
		datarw.Decode(&nowDrules)
		fnldta = append(fnldta, nowDrules)
	}

	// return data
	return fnldta
}
