package fncPsglst

// Running process hit passanggerlist daily
// func ApiPsglstProcess(c *gin.Context) {
// 	if fncGlobal.Status.Sbrapi == 0.0 {
// 		fncGlobal.Status.Sbrapi = 0.01

// 		// Bind JSON Body input to variable
// 		inputx := mdlPsglst.MdlPsglstErrlogDtbase{} //save
// 		if err := c.BindJSON(&inputx); err != nil {
// 			panic(err)
// 		}

// 		// Declare date format
// 		dteUpload := time.Now().Format("2006-01-02T15:04:05")
// 		dteMinday := time.Now().AddDate(0, 0, -1)
// 		if inputx.Datefl != "All" {
// 			dteMinday, _ = time.Parse("060102", inputx.Datefl)
// 		}

// 		// Final format date
// 		dteFldtef := strings.ToUpper(dteMinday.Format("02Jan"))
// 		dteFullfm := dteMinday.Format("2006-01-02T15:04:05")
// 		dteNumber := dteMinday.Format("060102")

// 		// Declare airline
// 		slcAirlnf := []string{"JT", "ID", "IW", "IU", "OD", "SL"}
// 		if inputx.Airlnf != "All" {
// 			slcAirlnf = []string{inputx.Airlnf}
// 		}

// 		// Declare Flight number
// 		slcFlnumf := []string{"All"}
// 		if inputx.Flnumf != "All" {
// 			slcFlnumf = []string{inputx.Flnumf}
// 		}

// 		// Indicator done data
// 		sycmapClscbn := fncGlobal.GetApndixClslvlSycmap()
// 		sycmapFlhour := fncGlobal.GetApndixFlhourSycmap()
// 		sycidcFlhour := &sync.Map{}
// 		sycmapFrbase := fncGlobal.GetApndixFrbaseSycmap()
// 		sycidcFrbase := &sync.Map{}
// 		sycmapFrtaxs := fncGlobal.GetApndixFrtaxsSycmap()
// 		sycidcFrtaxs := &sync.Map{}
// 		workerTotall := inputx.Worker

// 		// Action log
// 		sycmapErrlog, _ := fncGlobal.GetAlldtaErrlogSycmap("psglst", dteNumber)

// 		// Looping slice airlines
// 		for _, airlnf := range slcAirlnf {
// 			fmt.Println("Processing airline:", airlnf)

// 			// Get 10 API sessions/tokens
// 			slcRspssn, err := fncGlobal.GetAlldbsMltplssn(airlnf, workerTotall)
// 			lgcRspssn := err != nil || slcRspssn[0].Bsttkn == "" || len(slcRspssn) < 1
// 			func() {
// 				errPrmkey := "sssion" + airlnf + dteNumber
// 				errErrlog := alldbsmodel.ModelAlldbsGlobalErrlog{
// 					Prmkey: errPrmkey, Errprt: "fllist", Errout: "API",
// 					Errtxt: "Sessio Cannot get on Sabre Web API", Dvsion: "mnfest",
// 					Datenb: dteNumber, Fldtef: dteFldtef, Airlnf: airlnf,
// 					Depart: "All", Flnumf: "All", Dateup: dteUpload, Flstat: "All",
// 					Flhour: 0.0, Routef: "All", Status: "Pending", Worker: 8}
// 				fncGlobal.PshGlobalErrlogPanicr(errErrlog, lgcRspssn, errPrmkey, sycmapErrlog)
// 			}()

// 			// Prepare job queue
// 			jobFllist := make(chan alldbsmodel.ModelAlldbsApndixFllist, 1500)
// 			var swg sync.WaitGroup

// 			// Launch 10 workers using 10 tokens
// 			for i := 0; i < workerTotall; i++ {
// 				if len(slcRspssn) >= i+1 {
// 					fmt.Println(slcRspssn[i])
// 					if slcRspssn[i].Bsttkn != "" {
// 						swg.Add(1)
// 						go actPsglstWorker(i, &slcRspssn[i], slcFlnumf, inputx.Ignore,
// 							jobFllist, &swg, sycmapFlhour, sycidcFlhour, sycmapFrbase,
// 							sycidcFrbase, sycmapFrtaxs, sycidcFrtaxs, sycmapErrlog, sycmapClscbn)
// 					}
// 				}
// 			}

// 			// Declare Depart
// 			slcDepart := strings.Split(inputx.Depart, "-")
// 			if inputx.Depart == "All" {
// 				slcDepart = fncGlobal.GetApndixDepartSlcarr(airlnf)
// 				slcDepart = slcDepart
// 			}

// 			// Looping slice departures
// 			for _, depart := range slcDepart {

// 				// Get API Flight List data
// 				rspFllist, err := fncGlobal.GetApisbrFllist(slcRspssn[0], depart,
// 					airlnf, dteFldtef, dteFullfm[:10], dteNumber)
// 				func() {
// 					errPrmkey := "fllist" + airlnf + depart + dteNumber
// 					errErrlog := alldbsmodel.ModelAlldbsGlobalErrlog{
// 						Prmkey: errPrmkey, Errprt: "fllist", Errout: "API", Dvsion: "mnfest",
// 						Errtxt: "Flight List Cannot found on Sabre Web API",
// 						Datenb: dteNumber, Fldtef: dteFldtef, Airlnf: airlnf,
// 						Depart: depart, Flnumf: "All", Dateup: dteUpload, Flstat: "All",
// 						Flhour: 0.0, Routef: "All", Status: "Pending", Worker: 10}
// 					fncGlobal.PshGlobalErrlogPanicr(errErrlog, err != nil, errPrmkey, sycmapErrlog)
// 				}()

// 				// Looping Flight List
// 				for _, fllist := range rspFllist {

// 					// Only accept on this route
// 					fllist.Datefm = dteFullfm
// 					fllist.Datefl = dteNumber
// 					fllist.Dateup = dteUpload
// 					fllist.Fldtef = dteFldtef
// 					if slices.Contains(slcFlnumf, fllist.Flnumf) ||
// 						slices.Contains(slcFlnumf, "All") {
// 						jobFllist <- fllist
// 					}
// 				}
// 			}

// 			// Finish
// 			close(jobFllist)
// 			swg.Wait()
// 			// defer fncGlobal.ActAlldbsClosessn(slcRspssn) ini pakai defer sebelumnya dr data dibawwah
// 			fncGlobal.ActAlldbsClosessn(slcRspssn)
// 			fmt.Println("Finished:", airlnf)
// 		}

// 		// Update status if done
// 		rspErrlog := "Updated"
// 		logAction := alldbsmodel.ModelAlldbsMainpgActlog{
// 			Prmkey: dteNumber, Dateup: dteUpload, Fldtef: dteFldtef,
// 			Datenb: dteNumber, Dtstat: "Done"}
// 		if nowErrlog, ist := fncGlobal.GetAlldtaErrlogSycmap("psglst", dteNumber); ist {
// 			logAction.Dtstat = "Pending"
// 			if _, is2 := nowErrlog.Load(inputx.Prmkey); is2 {
// 				rspErrlog = "Failed"
// 			}
// 		}

// 		// Update action log
// 		rsupdt := fncGlobal.PshAlldtaMgomdlBlkwrt([]mongo.WriteModel{
// 			mongo.NewUpdateOneModel().
// 				SetFilter(bson.M{"prmkey": logAction.Prmkey}).
// 				SetUpdate(bson.M{"$set": logAction}).
// 				SetUpsert(true)}, "psglst_actlog")
// 		if !rsupdt {
// 			panic("Error Insert/Update to DB")
// 		}
// 		fncGlobal.Status = "Done"
// 		c.JSON(200, rspErrlog)
// 	}
// }
