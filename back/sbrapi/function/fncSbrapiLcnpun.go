package fncSbrapi

import (
	mdlSbrapi "back/sbrapi/model"
	"strconv"
	"strings"
	"time"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiLcnpunMainob(lcnpun string, unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix) ([]mdlSbrapi.MdlSbrapiLcnpunDtbase, error) {

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))
	strComand := lcnpun + params.Flnbfl + "/" + ddmDatefl + params.Depart

	// Isi struktur data
	fnlLcnpun := []mdlSbrapi.MdlSbrapiLcnpunDtbase{}
	strOutput, err := FncSbrapiCmdscrMainob(unqhdr, strComand)
	if err != nil {
		return fnlLcnpun, err
	}

	//  Final data
	fnlLcnpun = FncSbrapiLcnpunPrcess(lcnpun, strOutput, params)
	return fnlLcnpun, nil
}

// Function Treatment for API LC AND PUN
func FncSbrapiLcnpunPrcess(lcnpun, output string, params mdlSbrapi.MdlSbrapiMsghdrApndix,
) []mdlSbrapi.MdlSbrapiLcnpunDtbase {

	// Declare first output
	var result []mdlSbrapi.MdlSbrapiLcnpunDtbase
	var allpnr = map[string]bool{}
	var timenw = time.Now().Format("0601021504")
	var intTimenw, _ = strconv.Atoi(timenw)
	var intDatenw, _ = strconv.Atoi(timenw[0:6])

	// Looping data
	outlne := strings.Split(output, "\n")
	for _, outrow := range outlne {

		// Skip if not LC/PUN/LDN
		if len(outrow) <= 6 {
			continue
		}

		// Split data
		slcrow := strings.Split(outrow, ".")
		clnrow := []string{}
		for _, row := range slcrow {
			if strings.TrimSpace(row) != "" {
				clnrow = append(clnrow, row)
			}
		}

		// Push to database LC AND PUN
		if len(clnrow) >= 3 {
			totpax, _ := strconv.Atoi(strings.TrimSpace(clnrow[0][3:6]))
			agtnme := strings.TrimSpace(clnrow[0][6:len(clnrow[0])])
			pnrcde := clnrow[2]
			clssfl := clnrow[1][:1]
			if len(clnrow) == 4 {
				pnrcde = clnrow[3]
				clssfl = clnrow[2]
			}

			// Assign data
			if _, ist := allpnr[pnrcde]; !ist {
				allpnr[pnrcde] = true
				result = append(result, mdlSbrapi.MdlSbrapiLcnpunDtbase{
					Prmkey: lcnpun + params.Airlfl + params.Flnbfl + params.Depart +
						strconv.Itoa(int(params.Datefl)) + pnrcde + timenw[0:6],
					Airlfl: params.Airlfl, Lcrpun: lcnpun, Totpax: totpax,
					Flnbfl: params.Flnbfl, Depart: params.Depart,
					Routfl: params.Routfl, Clssfl: strings.TrimSpace(clssfl),
					Datefl: params.Datefl, Dateup: int32(intDatenw),
					Timeup: int64(intTimenw), Agtnme: agtnme, Pnrcde: pnrcde})
			}
		}
	}

	// Return final data
	return result
}
