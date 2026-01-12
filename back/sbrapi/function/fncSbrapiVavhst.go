package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiVavhstMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix) ([]mdlSbrapi.MdlSbrapiVavhstDtbase, error) {

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))
	strComand := "VAV" + params.Flnbfl + "/" + ddmDatefl

	// Isi struktur data
	fnlVavhst := []mdlSbrapi.MdlSbrapiVavhstDtbase{}
	strOutput, err := FncSbrapiCmdscrMainob(unqhdr, strComand)
	if err != nil {
		return fnlVavhst, err
	}

	//  Final data
	fnlVavhst = FncSbrapiVavhstPrcess(strOutput, params)
	return fnlVavhst, nil
}

// Function Treatment for API LC AND PUN
func FncSbrapiVavhstPrcess(output string, params mdlSbrapi.MdlSbrapiMsghdrApndix,
) []mdlSbrapi.MdlSbrapiVavhstDtbase {

	// Declare first output
	var result []mdlSbrapi.MdlSbrapiVavhstDtbase
	var tmprsl mdlSbrapi.MdlSbrapiVavhstDtbase

	// Looping data
	outlne := strings.Split(output, "\n")
	for idx, outrow := range outlne {
		slcstr := strings.Fields(outrow)
		if len(slcstr) > 1 {

			// If first data
			if idx%2 == 0 && strings.Contains(slcstr[0], ".") {
				slcDateup := strings.Split(slcstr[0], ".")

				// Get date
				if strings.Contains(slcDateup[1], "/") {
					slcDateup := strings.Split(slcDateup[1], "/")
					rawDateup := slcDateup[1]
					strDateup := fncGlobal.FncGlobalMainprDaymnt(rawDateup)
					rawTimeup := slcDateup[0]
					intTimeup, _ := strconv.Atoi(strDateup + rawTimeup)
					tmprsl.Timeup = int64(intTimeup)
				}

				// Get other data
				tmprsl.Clssfl = slcstr[1]
				tmprsl.Routfl = slcstr[2]
				tmprsl.Depart = slcstr[2][:3]
				tmprsl.Statfl = slcstr[3]
				intTotpax, _ := strconv.Atoi(slcstr[4])
				tmprsl.Totpax = int64(intTotpax)
				continue
			}

			// Second data
			if len(slcstr) > 2 {
				tmprsl.Lniata = "lniata:" + slcstr[0]
				tmprsl.Agtdie = slcstr[1]
				tmprsl.Airlfl = slcstr[2]
				tmprsl.Prmkey = fmt.Sprintf("%v%v%v%v%v", tmprsl.Airlfl,
					tmprsl.Lniata, tmprsl.Timeup, params.Flnbfl, tmprsl.Routfl)
				tmprsl.Flnbfl = params.Flnbfl
				tmprsl.Datefl = params.Datefl
				result = append(result, tmprsl)
			}
		}
	}

	// Return final data
	return result
}
