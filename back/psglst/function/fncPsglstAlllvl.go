package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Sync map data LC and PUN prev day
func FncPsglstClslvlMapobj() map[string]mdlPsglst.MdlPsglstClsslvDtbase {

	// Inisialisasi variabel
	fnldta := map[string]mdlPsglst.MdlPsglstClsslvDtbase{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_clsslv")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlPsglst.MdlPsglstClsslvDtbase
		datarw.Decode(&object)
		fnldta[object.Clssfl] = object
	}

	// return data
	return fnldta
}

// Get Sync map data LC and PUN prev day
func FncPsglstHfbalvMapobj() []mdlPsglst.MdlPsglstHfbalvDtbase {

	// Inisialisasi variabel
	fnldta := []mdlPsglst.MdlPsglstHfbalvDtbase{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_hfbalv")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{},
		options.Find().SetSort(bson.D{{Key: "levelr", Value: 1}}))
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlPsglst.MdlPsglstHfbalvDtbase
		datarw.Decode(&object)
		fnldta = append(fnldta, object)
	}

	// return data
	return fnldta
}
