package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data LC and PUN prev day
func FncJeddahPnrlogSycmap(prvDateup string) *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}
	tmpdta := map[string]map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	intDateup, _ := strconv.Atoi(prvDateup)
	datarw, err := tablex.Find(contxt, bson.M{"dateup": intDateup, "lcrpun": "lc"})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlJeddah.MdlJeddahPnrlogDtbase
		datarw.Decode(&object)
		syckey := object.Airlfl + object.Flnbfl + object.Depart + strconv.Itoa(int(object.Datefl))
		if _, ist := tmpdta[syckey]; !ist {
			if len(tmpdta) < 1 {
				tmpdta = map[string]map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}
			}
			tmpdta[syckey] = map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}
		}
		tmpdta[syckey][object.Pnrcde] = object
	}

	// Push to final data
	for key, val := range tmpdta {
		fnldta.Store(key, val)
	}

	// return data
	return fnldta
}
