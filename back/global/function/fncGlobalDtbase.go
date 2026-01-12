package fncGlobal

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Input Bulkwrite to the database All
func FncGlobalDtbaseBlkwrt(dtamdl []mongo.WriteModel, cltion string) error {

	// Select database and collection
	tablex := Client.Database(Dbases).Collection(cltion)
	contxt, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	// Insert batch into MongoDB
	optyns := options.BulkWrite().SetOrdered(false) // Tidak harus urut
	_, errorx := tablex.BulkWrite(contxt, dtamdl, optyns)
	if errorx != nil {
		fmt.Println(errorx)
	}
	return errorx
}

// Input to the database All
func FncGlobalDtbaseBtcwrt(mapMgomdl map[string]*[]mongo.WriteModel, blimit int) {
	for key, mgo := range mapMgomdl {
		if len(*mgo) > blimit {
			rsupdt := FncGlobalDtbaseBlkwrt(*mgo, key)
			if rsupdt != nil {
				panic("Error Insert/Update to DB:" + rsupdt.Error())
			}
			*mgo = []mongo.WriteModel{}
		}
	}
}
