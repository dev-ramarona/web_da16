package fncPsglst

import (
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func FncPsglstFrcalcSplitd(psglst *mdlPsglst.MdlPsglstPsgdtlDtbase,
	mapCurrcv map[string]mdlPsglst.MdlPsglstCurrcvDtbase, sycMilege *sync.Map,
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams) bool {
	cekNonrev := false
	slcFrcalc := strings.Fields(psglst.Frcalc)
	nowFrbase := [20]mdlPsglst.MdlPsglstFrcalcFrbase{}
	regDstrc1 := regexp.MustCompile(`^[A-Z]{3}$`)
	regDstrc2 := regexp.MustCompile(`(^I-|^X/|^S-.+|^S-|^F-)([A-Z]{3}$)`)
	regDstrc3 := regexp.MustCompile(`(^[A-Z]{3})(/IT$)`)
	regDstrc4 := regexp.MustCompile(`(^I-.+)([A-Z]{3}$)`)
	regDstrc5 := regexp.MustCompile(`(^\d{2}[A-Z]{3}\d{2})([A-Z]{3}$)`)
	regCrrat1 := regexp.MustCompile(`(^ROE)(\d+\.\d+|\d+)`)
	regCrrat2 := regexp.MustCompile(`\d+\.\d+|\d+`)
	regAirlfl := regexp.MustCompile(`^[A-Z]{2}$`)
	regFrend1 := regexp.MustCompile(`(^[A-Z]{3})(\d+\.\d{2}|\d+)([A-Z]{3})(\d+\.\d{2}|\d+)(END)`)
	regFrend2 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)([A-Z]{3})(\d+\.\d{2}|\d+)(END)`)
	regFrend3 := regexp.MustCompile(`(^[A-Z]{3})(\d+\.\d{2}|\d+)(END)`)
	regFrend4 := regexp.MustCompile(`(^[A-Z]{6})(\d+\.\d{2}|\d+)([A-Z]{3})(\d+\.\d{2}|\d+)(END$)`)
	regFrbse1 := regexp.MustCompile(`(^[A-Z]{3})(\d+\.\d{2}|\d+)([A-Z]{2}$)`)
	regFrbse2 := regexp.MustCompile(`(^[A-Z]{3}\d+M|^[A-Z]{3})(\d+\.\d{2}|\d+)([A-Z]+)`)
	regFrbse3 := regexp.MustCompile(`(^[A-Z]{3})(\d+\.\d{2}|\d+)(/-)([A-Z]{3}$)`)
	regFrbse4 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)(/-)([A-Z]{3}$)`)
	regFrbse5 := regexp.MustCompile(`(^/-)([A-Z]{3}$)`)
	regFrbse6 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)([A-Z]{2}$)`)
	regFrbse7 := regexp.MustCompile(`^\d+\.\d{2}$|^\d+$`)
	regFrbse8 := regexp.MustCompile(`^XT\d+`)
	regFrbse9 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)([A-Z]{2})([A-Z]{3}$)`)
	regFrbs10 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)([A-Z]+$)`)
	regFrbs11 := regexp.MustCompile(`(^\d+\.\d{2}|^\d+)([A-Z]{2})(\d+\.\d{2}$|\d+$)`)
	regQschg1 := regexp.MustCompile(`(^[A-Z]{6})(\d+$|\d+.\d{2}$)`)
	regQschg2 := regexp.MustCompile(`(^[A-Z]{6})(\d+$|\d+.\d{2})`)
	regQschg3 := regexp.MustCompile(`Q\d+\.\d{2}|Q\d+|\d+.\d{2}|\d+|[A-Z]{3}`)
	regQschg4 := regexp.MustCompile(`^Q$`)
	regQsrate := regexp.MustCompile(`(^[A-Z]{6})(\d+\.\d{2}|\d+)(NUC)(\d+\.\d{2}|\d+)(ENDROE)(\d+\.\d{2}|\d+)`)
	regNotuse := regexp.MustCompile(`^PD\d.+|^M/IT|^/IT|^/BT|^Q$`)
	nowDstrct := ""
	idxCountd := 1
	prvRateof := false
	prvQsrchg := false
	prvCrrate := false

	// Looping all fare
brk:
	for _, slc := range slcFrcalc {
		switch {

		// Logic match dstrct (normal)
		case len(regDstrc1.FindAllString(slc, -1)) >= 1:
			rsl := regDstrc1.FindAllString(slc, -1)
			if rsl[0] == "END" || rsl[0] == "ROE" {
				prvRateof = true
				continue
			}
			if nowDstrct != "" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[0]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[0]
				idxCountd++
			}
			nowDstrct = rsl[0]

		// Logic match dstrct (upnormal)
		case len(regDstrc2.FindStringSubmatch(slc)) >= 1:
			rsl := regDstrc2.FindStringSubmatch(slc)
			if nowDstrct != "" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[2]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[2]
				if rsl[1] == "X/" {
					nowFrbase[idxCountd].Isitpr = "Next Frbase"
				}
				idxCountd++
			}
			nowDstrct = rsl[2]

		// Logic match dstrct (/IT)
		case len(regDstrc3.FindStringSubmatch(slc)) >= 1:
			rsl := regDstrc3.FindStringSubmatch(slc)
			if nowDstrct != "" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[1]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[1]
			}
			nowDstrct = rsl[1]

		// Logic match dstrct (/I-DATE)
		case len(regDstrc4.FindStringSubmatch(slc)) >= 1:
			rsl := regDstrc4.FindStringSubmatch(slc)
			if nowDstrct != "" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[2]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[2]
			}
			nowDstrct = rsl[2]

		// Logic match dstrct (DATE DSTRCT)
		case len(regDstrc5.FindStringSubmatch(slc)) >= 1:
			rsl := regDstrc5.FindStringSubmatch(slc)
			if nowDstrct != "" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[2]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[2]
			}
			nowDstrct = rsl[2]

		// Logic match ROE
		case len(regCrrat1.FindStringSubmatch(slc)) >= 1:
			rsl := regCrrat1.FindStringSubmatch(slc)

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Crrate = rsl[2]
				}
			}
			prvCrrate = true

		// Logic match ROE
		case prvRateof, prvCrrate:
			rsl := regCrrat2.FindAllString(slc, -1)

			// Push all currency
			if len(rsl) > 0 {
				for idx := range nowFrbase {
					if nowFrbase[idx].Routfl != "" {
						if prvCrrate {
							nowFrbase[idx].Crrate += rsl[0]
						} else {
							nowFrbase[idx].Crrate = rsl[0]
						}
					}
				}
			}
			break brk

		// Logic match dstrct airlfl
		case len(regAirlfl.FindAllString(slc, -1)) >= 1:
			rsl := regAirlfl.FindAllString(slc, -1)
			nowFrbase[idxCountd].Airlfl = rsl[0]

		// Logic match end (normal)
		case len(regFrend1.FindStringSubmatch(slc)) >= 1:
			rsl := regFrend1.FindStringSubmatch(slc)
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[1]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[1]
			nowFrbase[idxCountd].Frbase = rsl[2]
			idxCountd++
			nowDstrct = rsl[1]

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Curncy = rsl[3]
				}
			}

		// Logic match end (upnormal)
		case len(regFrend2.FindStringSubmatch(slc)) >= 1:
			rsl := regFrend2.FindStringSubmatch(slc)
			nowFrbase[idxCountd-1].Frbase = rsl[1]
			idxCountd++

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Curncy = rsl[2]
				}
			}

		// Logic match end (upnormal 2)
		case len(regFrend3.FindStringSubmatch(slc)) >= 1:
			rsl := regFrend3.FindStringSubmatch(slc)

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Curncy = rsl[1]
				}
			}

		// Logic match end (upnormal 2)
		case len(regFrend4.FindStringSubmatch(slc)) >= 1:
			rsl := regFrend4.FindStringSubmatch(slc)
			if nowFrbase[idxCountd-1].Qsrcrw == "" {
				nowFrbase[idxCountd-1].Qsrcrw = rsl[1] + "Q" + rsl[2]
			} else {
				nowFrbase[idxCountd-1].Qsrcrw += "|" + rsl[1] + "Q" + rsl[2]
			}

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Curncy = rsl[3]
				}
			}

		// Logic match dstrct - frbase - airlfl (normal)
		case len(regFrbse1.FindStringSubmatch(slc)) > 1:
			rsl := regFrbse1.FindStringSubmatch(slc)
			if rsl[1] != "NUC" {
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[1]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[1]
				nowFrbase[idxCountd].Frbase = rsl[2]
				nowFrbase[idxCountd+1].Airlfl = rsl[3]
				idxCountd++
				nowDstrct = rsl[1]
			}

		// Logic match dstrct - frbase - frcode (upnormal)
		case len(regFrbse2.FindStringSubmatch(slc)) > 1:
			rsl := regFrbse2.FindStringSubmatch(slc)
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[1][:3]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[1][:3]
			nowFrbase[idxCountd].Frbase = rsl[2]
			idxCountd++
			nowDstrct = rsl[1][:3]

			// Push all currency
			if len(rsl[3]) == 3 && rsl[3] == "NUC" {
				for idx := range nowFrbase {
					if nowFrbase[idx].Routfl != "" {
						nowFrbase[idx].Curncy = rsl[3]
					}
				}
			}

		// Logic match dstrct - frbase - void - dstrct
		case len(regFrbse3.FindStringSubmatch(slc)) > 1:
			rsl := regFrbse3.FindStringSubmatch(slc)
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[1]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[1]
			nowFrbase[idxCountd].Frbase = rsl[2]
			idxCountd++
			nowDstrct = rsl[1]

			// Void data
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[4]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[4]
			idxCountd++
			nowDstrct = rsl[4]

		// Logic match frbase - void - dstrct
		case len(regFrbse4.FindStringSubmatch(slc)) > 1:
			rsl := regFrbse4.FindStringSubmatch(slc)
			nowFrbase[idxCountd-1].Frbase = rsl[1]

			// Void data
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[3]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[3]
			idxCountd++
			nowDstrct = rsl[3]

		// Logic match void - dstrct
		case len(regFrbse5.FindStringSubmatch(slc)) > 1:
			rsl := regFrbse5.FindStringSubmatch(slc)

			// Void data
			nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[2]
			nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
			nowFrbase[idxCountd].Depart = nowDstrct
			nowFrbase[idxCountd].Arrivl = rsl[2]
			idxCountd++
			nowDstrct = rsl[2]

		// Logic match frbase - airlfl
		case len(regFrbse6.FindStringSubmatch(slc)) >= 1:
			rsl := regFrbse6.FindStringSubmatch(slc)
			nowFrbase[idxCountd-1].Frbase = rsl[1]
			nowFrbase[idxCountd].Airlfl = rsl[2]

		// Logic match frbase
		case len(regFrbse7.FindAllString(slc, -1)) >= 1:
			rsl := regFrbse7.FindAllString(slc, -1)
			nowFrbase[idxCountd-1].Frbase = rsl[0]

		// Logic match frbase
		case len(regFrbse8.FindAllString(slc, -1)) >= 1:
			regFrbse8.FindAllString(slc, -1)

		// Logic match frbase
		case len(regFrbse9.FindStringSubmatch(slc)) >= 1:
			rsl := regFrbse9.FindStringSubmatch(slc)
			if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, rsl[2]) {
				nowFrbase[idxCountd].Airlfl = rsl[2]
				nowFrbase[idxCountd].Routfl = nowDstrct + "-" + rsl[3]
				nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
				nowFrbase[idxCountd].Depart = nowDstrct
				nowFrbase[idxCountd].Arrivl = rsl[3]
				idxCountd++
				nowDstrct = rsl[3]
			}
			nowFrbase[idxCountd-1].Frbase = rsl[1]

		// Logic match frbase
		case len(regFrbs10.FindStringSubmatch(slc)) >= 1:
			rsl := regFrbs10.FindStringSubmatch(slc)
			nowFrbase[idxCountd-1].Frbase = rsl[1]

		// Logic match frbase zero
		case len(regFrbs11.FindStringSubmatch(slc)) >= 1:
			rsl := regFrbs11.FindStringSubmatch(slc)
			nowFrbase[idxCountd-1].Frbase = rsl[1]

		// Logic match routefl - qsrchg
		case len(regQschg1.FindStringSubmatch(slc)) >= 1:
			rsl := regQschg1.FindStringSubmatch(slc)
			if nowFrbase[idxCountd-1].Qsrcrw == "" {
				nowFrbase[idxCountd-1].Qsrcrw = rsl[1] + "Q" + rsl[2]
			} else {
				nowFrbase[idxCountd-1].Qsrcrw += "|" + rsl[1] + "Q" + rsl[2]
			}

		// Logic match routefl - qsrchg
		case len(regQschg2.FindStringSubmatch(slc)) >= 1 && prvQsrchg:
			prvQsrchg = false
			rsl := regQschg2.FindStringSubmatch(slc)
			if nowFrbase[idxCountd-1].Qsrcrw == "" {
				nowFrbase[idxCountd-1].Qsrcrw = rsl[1] + "Q" + rsl[2]
			} else {
				nowFrbase[idxCountd-1].Qsrcrw += "|" + rsl[1] + "Q" + rsl[2]
			}

		// Logic match qsrchg
		case len(regQschg3.FindAllString(slc, -1)) >= 1:
			rsl := regQschg3.FindAllString(slc, -1)
			slcQsrchg := []string{}
			if !strings.Contains(slc, "/") {
				for _, val := range rsl {
					if strings.Contains(val, "Q") {
						slcQsrchg = append(slcQsrchg, val)
					} else if val == "END" {
						break
					} else if len(val) == 3 && val != "END" {
						if len(rsl) <= 2 && val != "NUC" {
							nowFrbase[idxCountd].Routfl = nowDstrct + "-" + val
							nowFrbase[idxCountd].Cpnbfl = int32(idxCountd)
							nowFrbase[idxCountd].Depart = nowDstrct
							idxCountd++
							nowDstrct = val
						} else {
							for idx := range nowFrbase {
								if nowFrbase[idx].Routfl != "" &&
									nowFrbase[idx].Curncy == "" {
									nowFrbase[idx].Curncy = val
								}
							}
						}
					} else if _, err := strconv.ParseFloat(val, 64); err == nil {
						if nowFrbase[idxCountd-1].Frbase == "" {
							nowFrbase[idxCountd-1].Frbase = val
						}
					}
				}
				if len(slcQsrchg) > 0 {
					nowFrbase[idxCountd-1].Qsrcrw = strings.Join(slcQsrchg, "|")
				}
			}

		// Logic match qsrchg
		case len(regQschg4.FindAllString(slc, -1)) >= 1:
			regQschg4.FindAllString(slc, -1)
			prvQsrchg = true

		// Logic match routefl - qsrchg - nuc - tot - end - roe - crrate
		case len(regQsrate.FindStringSubmatch(slc)) >= 1:
			rsl := regQsrate.FindStringSubmatch(slc)
			if nowFrbase[idxCountd-1].Qsrcrw == "" {
				nowFrbase[idxCountd-1].Qsrcrw = rsl[1] + "Q" + rsl[2]
			} else {
				nowFrbase[idxCountd-1].Qsrcrw += "|" + rsl[1] + "Q" + rsl[2]
			}

			// Push all currency
			for idx := range nowFrbase {
				if nowFrbase[idx].Routfl != "" {
					nowFrbase[idx].Crrate = rsl[6]
				}
			}

		// Logic match not use data after end
		case len(regNotuse.FindAllString(slc, -1)) >= 1:
			rsl := regNotuse.FindAllString(slc, -1)
			if slices.Contains(rsl, "IT") || slices.Contains(rsl, "M/IT") {
				for idx := range nowFrbase {
					if nowFrbase[idx].Routfl != "" {
						nowFrbase[idx].Isitit = "IT FARE"
					}
				}
			}

		// Default
		default:
			// fmt.Println("default", slc, "|", idxCountd, "|", psglst.Prmkey,
			// 	"|", psglst.Pnrcde, "|", psglst.Tktnfl, "|", psglst.Frcalc)
		}
	}

	// Final data
	mapRoutcn := map[string]int{}
	mapFrcacl := map[string]mdlPsglst.MdlPsglstFrcalcFrbase{}
	prvPrrate := map[string]string{}
	slcRoutfr := []string{}
	for _, val := range nowFrbase {
		if val.Routfl != "" {
			slcRoutsg := strings.Split(psglst.Routsg, "-")
			keyFrcalc := strconv.Itoa(int(val.Cpnbfl)) + "|" + val.Routfl
			fltFrbase, _ := strconv.ParseFloat(val.Frbase, 64)
			fltCrrate, _ := strconv.ParseFloat(val.Crrate, 64)

			// Mapping kode bandara → group rute
			var mapDstrct = map[string][]string{
				"JKT": {"CGK", "HLP", "JKT"},
				"SRI": {"AAP", "SRI"},
				"JOG": {"YIA", "JOG"}}
			val.Depart = func() string {
				if slcDepart, ist := mapDstrct[val.Depart]; ist {
					for _, v := range slcDepart {
						for i := 0; i < len(slcRoutsg); i++ {
							if slcRoutsg[i] == v {
								slcRoutsg = append(slcRoutsg[:i], slcRoutsg[i+1:]...)
								val.Routfl = v + "-" + val.Routfl[4:]
								return v
							}
						}
					}
				}
				return val.Depart
			}()
			val.Arrivl = func() string {
				if slcArrivl, ist := mapDstrct[val.Arrivl]; ist {
					for _, v := range slcArrivl {
						for i := 0; i < len(slcRoutsg); i++ {
							if slcRoutsg[i] == v {
								slcRoutsg = append(slcRoutsg[:i], slcRoutsg[i+1:]...)
								val.Routfl = val.Routfl[:3] + "-" + v
								return v
							}
						}
					}
				}
				return val.Arrivl
			}()

			// Push count route calculation
			mapRoutcn[keyFrcalc] += 1
			if val.Arrivl != psglst.Depart && val.Depart != psglst.Arrivl {
				if val.Depart == psglst.Depart || val.Arrivl == psglst.Arrivl {
					mapRoutcn[keyFrcalc] += 5
					if val.Depart == psglst.Depart && val.Arrivl == psglst.Arrivl {
						mapRoutcn[keyFrcalc] += 15
					}
				}
				if len(psglst.Routvc) >= 7 {
					if val.Depart == psglst.Routvc[:3] || val.Arrivl == psglst.Routvc[4:] {
						mapRoutcn[keyFrcalc] += 5
						if val.Depart == psglst.Routvc[:3] && val.Arrivl == psglst.Routvc[4:] {
							mapRoutcn[keyFrcalc] += 15
						}
					}
				}
				if int(val.Cpnbfl) == int(psglst.Cpnbfl) || int(val.Cpnbfl) == int(psglst.Cpnbvc) {
					mapRoutcn[keyFrcalc] += 1
				}
			}

			// Get total Q
			var totQsrcvc float64
			if val.Qsrcrw != "" {
				regQsrcrw := regexp.MustCompile(`\d+(\.\d+)?`)
				mtcQsrcrw := regQsrcrw.FindAllString(val.Qsrcrw, -1)
				for _, qsrcrw := range mtcQsrcrw {
					valQsrcrw, _ := strconv.ParseFloat(qsrcrw, 64)
					totQsrcvc += valQsrcrw
				}
			}

			// Change currency
			nucNtafvc := fltFrbase
			nucQsrcvc := totQsrcvc
			// if psglst.Prmkey == "2512227051HLP8DRUMINTO" {
			// 	fmt.Println("xxx", fltFrbase, "*", fltCrrate, "=", fltFrbase*fltCrrate, val.Routfl)
			// }
			// if psglst.Prmkey == "2512227051HLP8DRUMINTO" {
			// 	fmt.Println("yyy", val)
			// }
			if val.Curncy == "NUC" {
				nucNtafvc = fltFrbase * fltCrrate
				nucQsrcvc = totQsrcvc * fltCrrate
			}
			val.Frbcnv = fmt.Sprintf("%v", nucNtafvc)
			val.Qsrcnv = fmt.Sprintf("%v", nucQsrcvc)
			if psglst.Curncy != "IDR" {
				if vlx, ist := mapCurrcv[psglst.Curncy]; ist {
					val.Frbcnv = fmt.Sprintf("%v", nucNtafvc/vlx.Crrate)
					val.Qsrcnv = fmt.Sprintf("%v", nucQsrcvc/vlx.Crrate)
				} else {
					val = mdlPsglst.MdlPsglstFrcalcFrbase{}
				}
			}
			mapFrcacl[keyFrcalc] = val
			slcRoutfr = append(slcRoutfr, val.Routfl)

			// Get prorate if isset
			if val.Isitit == "IT FARE" {
				prvPrrate = map[string]string{}
			} else if val.Isitpr == "Next Frbase" && val.Frbase == "" {
				prvPrrate[keyFrcalc] = val.Routfl
			} else if len(prvPrrate) > 0 {
				prvPrrate[keyFrcalc] = val.Routfl
				mapPrrate := map[string]float64{}
				totFrrate := float64(0)
				for keymap, routfl := range prvPrrate {

					// Reuse func frtaxs
					fncMilege := func(nowRoutac string) bool {
						istMilege, ist := sycMilege.Load(nowRoutac)
						if mtcMilege, mtc := istMilege.(mdlPsglst.MdlPsglstMilegeDtbase); mtc && ist {
							mapPrrate[keymap] = float64(mtcMilege.Milege)
							totFrrate += float64(mtcMilege.Milege)
						}
						return ist
					}
					if rspfnc := fncMilege(routfl); !rspfnc {
						rspMilege, err := fncSbrapi.FncSbrapiMilegeMainob(nowObjtkn, routfl)
						if err == nil {
							for _, milege := range rspMilege {
								if _, ist := sycMilege.Load(milege.Routfl); !ist {
									sycMilege.Store(milege.Routfl, milege)
								}
							}
							fncMilege(routfl)
						}
					}
				}

				// Calculate prorate
				prvPrrate = map[string]string{}
				nowFrbcnv := val.Frbcnv
				for keymap, milege := range mapPrrate {
					nowFrrate := milege / totFrrate
					if nowval, ist := mapFrcacl[keymap]; ist &&
						(nowval.Frbase == "" || nowval.Routfl == val.Routfl) {
						tmpFrbcnv, _ := strconv.ParseFloat(nowFrbcnv, 64)
						nowval.Frbcnv = fmt.Sprintf("%v", nowFrrate*tmpFrbcnv)
						nowval.Isitpr = "PRORTE"
						mapFrcacl[keymap] = nowval
						// if psglst.Prmkey == "2512227051HLP8DRUMINTO" {
						// 	fmt.Println("prorate", keymap, nowFrrate, "*", tmpFrbcnv, "=", nowFrrate*tmpFrbcnv)
						// }
					}
				}
			}
		}
	}

	// Get data vcr
	hghest := struct {
		key string
		cpn string
		val int
	}{}
	for keyseg, valint := range mapRoutcn {
		slckey := strings.Split(keyseg, "|")
		if slckey[0] > hghest.cpn && hghest.val == valint && hghest.val != 0 {
			continue
		}
		if hghest.val <= valint {
			hghest.key = keyseg
			hghest.val = valint
			hghest.cpn = slckey[0]
		}
	}
	if getFlsgmn, ist := mapFrcacl[hghest.key]; ist {
		if len(getFlsgmn.Routfl) == 7 {
			psglst.Ntafvc, _ = strconv.ParseFloat(getFlsgmn.Frbcnv, 64)
			psglst.Qsrcvc, _ = strconv.ParseFloat(getFlsgmn.Qsrcnv, 64)
			psglst.Qsrcrw = getFlsgmn.Qsrcrw
			psglst.Routfr = getFlsgmn.Routfl
			psglst.Isittf = "FRCALC"
			// if psglst.Prmkey == "2512227051HLP8DRUMINTO" {
			// 	fmt.Println("final", hghest.key, getFlsgmn.Frbcnv, psglst.Ntafvc)
			// }
			if getFlsgmn.Isitpr == "PRORTE" {
				psglst.Isittf = "PRCALC"
			}
			cekNonrev = getFlsgmn.Isitit == "" &&
				(getFlsgmn.Frbase == "0" || getFlsgmn.Frbase == "0.00")
			psglst.Routfx = strings.ReplaceAll(strings.Join(slcRoutfr, "|"), getFlsgmn.Routfl, "")
		}
	}

	return cekNonrev
}
