package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data LC and PUN prev day
func FncPsglstDepartGetslc() []string {

	// Inisialisasi variabel
	fnldta := []string{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_depart")
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
		var object mdlPsglst.MdlPsglstDepartDtbase
		datarw.Decode(&object)
		fnldta = append(fnldta, object.Depart)
	}

	// return data
	return fnldta
}
