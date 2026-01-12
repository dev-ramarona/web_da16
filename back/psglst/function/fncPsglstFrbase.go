package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"

	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data farebase
func FncPsglstFrbaseSycmap() *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_frbase")
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
		var object mdlPsglst.MdlPsglstFrbaseDtbase
		datarw.Decode(&object)
		fnldta.Store(object.Prmkey, object)
	}

	// return data
	return fnldta
}
