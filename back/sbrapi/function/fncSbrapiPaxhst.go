package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiPaxhstMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix, seqnce int) ([]mdlSbrapi.MdlSbrapiPaxhstDtbase,
	[]mdlSbrapi.MdlSbrapiPaxhstRawdta, error) {
	fnlPaxhst := []mdlSbrapi.MdlSbrapiPaxhstDtbase{}
	fnlRawhst := []mdlSbrapi.MdlSbrapiPaxhstRawdta{}

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))
	if seqnce == 0 {
		seqnce = 1
	}

	// Looping all page
	idxSeqnce := 0
	for i := 1; i <= seqnce; i++ {

		// Isi struktur data
		strComand := fmt.Sprintf("G*H/%v/%v%v*%v", params.Flnbfl,
			ddmDatefl, params.Depart, i)
		strOutput, err := FncSbrapiCmdscrMainob(unqhdr, strComand)
		if err != nil {
			return fnlPaxhst, fnlRawhst, err
		}

		//  Final data
		tmpPaxhst, tmpRawhst := FncSbrapiPaxhstPrcess(strOutput, params, &idxSeqnce)
		fnlPaxhst = append(fnlPaxhst, tmpPaxhst...)
		fnlRawhst = append(fnlRawhst, tmpRawhst...)
	}
	return fnlPaxhst, fnlRawhst, nil
}

// Function Treatment for API LC AND PUN
func FncSbrapiPaxhstPrcess(output string, params mdlSbrapi.MdlSbrapiMsghdrApndix, sqence *int,
) ([]mdlSbrapi.MdlSbrapiPaxhstDtbase, []mdlSbrapi.MdlSbrapiPaxhstRawdta) {

	// Declare first output
	var result []mdlSbrapi.MdlSbrapiPaxhstDtbase
	var tmprsl mdlSbrapi.MdlSbrapiPaxhstDtbase
	var rawrsl []mdlSbrapi.MdlSbrapiPaxhstRawdta

	// Looping data
	outlne := strings.Split(output, "\n")
	for _, outrow := range outlne {
		rawrsl = append(rawrsl, mdlSbrapi.MdlSbrapiPaxhstRawdta{
			Prmkey: fmt.Sprintf("%v%v%v%v%v%v", params.Datefl, params.Depart,
				params.Airlfl, params.Flnbfl, "-", fmt.Sprintf("%06d", *sqence)),
			Rawhst: outrow})
		(*sqence)++
		if strings.Contains(outrow, "RES AMEND:") {
			tmprsl = mdlSbrapi.MdlSbrapiPaxhstDtbase{
				Agtdie: outrow[11:15],
				Agtcty: outrow[:3],
				Lniata: "lniata:" + outrow[27:33],
				Itemhs: outrow[38:],
				Depart: params.Depart,
				Flnbfl: params.Flnbfl,
				Datefl: params.Datefl}

			// Get time and date
			strDateup := fncGlobal.FncGlobalMainprDaymnt(outrow[16:21])
			strTimeup := strDateup + outrow[22:26]
			intTimeup, _ := strconv.Atoi(strTimeup)
			tmprsl.Timeup = int64(intTimeup)
			continue
		}

		// Add additional data
		if tmprsl.Agtdie != "" && len(outrow) > 35 {
			getPnrcde := outrow[4:10]
			if regexp.MustCompile(`^[A-Za-z0-9]{6}$`).MatchString(getPnrcde) {
				tmprsl.Pnrcde = getPnrcde
				tmprsl.Nmefst = outrow[11:21]
				tmprsl.Nmelst = outrow[22:28]
				tmprsl.Arrivl = outrow[28:31]
				tmprsl.Seatpx = outrow[34:37]
				if len(outrow) >= 43 {
					tmprsl.Qntybt = outrow[40:43]
					if len(outrow) > 43 {
						slcfld := strings.Fields(outrow[43:])
						tmprsl.Codels = strings.Join(slcfld, "|")
					}
				}

				// Add item historry
				strItemhs := strings.ReplaceAll(outrow[:35], " ", "")
				if strItemhs == "" {
					tmprsl.Itemhs += strItemhs
				}
				tmprsl.Prmkey = fmt.Sprintf("%v%v%v%v%v%v", tmprsl.Lniata, tmprsl.Pnrcde,
					tmprsl.Timeup, tmprsl.Agtdie, params.Flnbfl, params.Depart)
				result = append(result, tmprsl)
			}
		}
	}

	// Return final data
	return result, rawrsl
}
