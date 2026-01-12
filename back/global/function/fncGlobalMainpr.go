package fncGlobal

import (
	mdlGlobal "back/global/model"
	"bufio"
	"crypto/rand"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Load Environment Variables from .env file
func FncGlobalMainprLoadnv(filenm string) {
	filenw, err := os.Open(filenm)
	if err != nil {
		panic("Error opening .env file:" + err.Error())
	}
	defer filenw.Close()

	// Scan the file line by line
	scnner := bufio.NewScanner(filenw)
	for scnner.Scan() {
		linenw := strings.TrimSpace(scnner.Text())
		if linenw == "" || strings.HasPrefix(linenw, "#") {
			continue
		}
		partnw := strings.SplitN(linenw, "=", 2)
		if len(partnw) == 2 {
			prmkey := strings.TrimSpace(partnw[0])
			valuex := strings.TrimSpace(partnw[1])
			os.Setenv(prmkey, valuex)
		}
	}
}

// manual load
var Status = mdlGlobal.MdlGlobalAllusrStatus{Sbrapi: 0, Action: 0}
var Client *mongo.Client
var jwtkey []byte
var Ipalow []string
var Secure = false
var Dbases, Urlmgo, Pcckey, Usrnme,
	Psswrd, Ptgolg, Ipadrs, Usrcok, Tknnme string

// Initial load environment variables
func init() {
	FncGlobalMainprLoadnv("../front/.env")
	jwtkey = []byte(os.Getenv("NEXT_PUBLIC_JWT_SECRET"))
	Dbases = os.Getenv("NEXT_PUBLIC_VAR_DTBASE")
	Urlmgo = os.Getenv("NEXT_PUBLIC_URI_MONGOS")
	Pcckey = os.Getenv("NEXT_PUBLIC_SBR_PCCKEY")
	Usrnme = os.Getenv("NEXT_PUBLIC_SBR_USRNME")
	Psswrd = os.Getenv("NEXT_PUBLIC_SBR_PSSWRD")
	Ptgolg = os.Getenv("NEXT_PUBLIC_PRT_GOLANG")
	Ipadrs = os.Getenv("NEXT_PUBLIC_IPV_ADRESS")
	Tknnme = os.Getenv("NEXT_PUBLIC_TKN_COOKIE")
	Ipalow = strings.Split(os.Getenv("NEXT_PUBLIC_IPV_ALLOWD"), "|")
	Tmpscr, err := strconv.ParseBool(os.Getenv("NEXT_PUBLIC_IPV_SECURE"))
	if err == nil {
		Secure = Tmpscr
	}
}

// Get status data process
func FncGlobalMainprStatus(c *gin.Context) {
	c.JSON(http.StatusOK, Status)
}

// Generate UUID
func FncGlobalMainprCduuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// versi 4 UUID (random)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// Convert time 12.30 to float 12.5
func FncGlobalMainprHstory(prvval, nowval any, hstory string,
	datend, datenw int32) (int32, string) {
	var fnlDatend, fnlHstory = datenw, ""
	if prvval == nowval {
		fnlDatend = datend
	} else if nowval != "" && nowval != 0 {
		arrHstory := []string{}
		if hstory != "" {
			arrHstory = strings.Split(hstory, "|")
		}
		arrHstory = append(arrHstory, fmt.Sprintf("%v:%v", datend, prvval))
		fnlHstory = strings.Join(arrHstory, "|")
	}
	return fnlDatend, fnlHstory
}

// Convert time 12.30 to float 12.5
func FncGlobalMainprFlhour(tmestr string) (float64, error) {

	// Pisahkan berdasarkan titik
	flhour := strings.Split(strings.Trim(tmestr, " "), ".")
	if len(flhour) != 2 {
		return 0, fmt.Errorf("format waktu tidak valid")
	}

	// Get hours and minutes
	hournw, er1 := strconv.Atoi(flhour[0])
	minute, er2 := strconv.Atoi(flhour[1])
	if er1 != nil || er2 != nil {
		return 0, fmt.Errorf("gagal mengonversi angka")
	}

	// Convert format to float
	dcimal := float64(hournw) + float64(minute)/60
	return dcimal, nil
}

// Treatment 920A / 1230P to string format time
func FncGlobalMainprFltime(timefl string) (string, error) {

	// Pastikan fltime memiliki minimal 3 karakter (contoh: "305A")
	if len(timefl) < 3 {
		return "0000", fmt.Errorf("format fltime hhmmA/P tidak valid")
	}

	// Ambil bagian menit (2 digit terakhir sebelum A/P)
	minute := timefl[len(timefl)-3 : len(timefl)-1]
	amorpm := timefl[len(timefl)-1:]
	hournb := timefl[:len(timefl)-3]

	// Konversi jam ke integer
	newHournb, err := strconv.Atoi(hournb)
	if err != nil {
		return "0000", fmt.Errorf("format jam hhmmA/P tidak valid")
	}

	// Konversi menit ke integer
	newMinute, err := strconv.Atoi(minute)
	if err != nil {
		return "0000", fmt.Errorf("format menit hhmmA/P tidak valid")
	}

	// Konversi format AM/PM ke 24 jam
	switch amorpm {
	case "P":
		if newHournb != 12 {
			newHournb += 12
		}
	case "A":
		if newHournb == 12 {
			newHournb = 0
		}

	}

	// Format hasil menjadi "hhmm"
	hhmm := fmt.Sprintf("%02d%02d", newHournb, newMinute)
	return hhmm, nil
}

// Note error manage
func FncGlobalMainprNoterr(strerr *string, varerr string) string {
	if !strings.Contains(*strerr, varerr) {
		if *strerr == "" {
			*strerr = varerr
			return *strerr
		}
		*strerr += "|" + varerr
	}
	return *strerr
}

// get year format DDMMM and change the data
func FncGlobalMainprDaymnt(daymnt string) string {
	strDatenw := time.Now().Format("060102")
	fmtDatenw, _ := time.Parse("060102", strDatenw)
	difFinald, strDatenw := 0, ""
	for idx, yearvl := range []int{-1, 0, 1} {
		strYearnw := time.Now().AddDate(yearvl, 0, 0).Format("06")
		fmtSbrenw, _ := time.Parse("02Jan06", daymnt+strYearnw)
		difDatenw := fmtDatenw.Sub(fmtSbrenw)
		difAbslte := int(math.Abs(difDatenw.Hours() / 24))
		if idx == 0 {
			difFinald = difAbslte
			strDatenw = fmtSbrenw.Format("060102")
		} else if difFinald > difAbslte {
			difFinald = difAbslte
			strDatenw = fmtSbrenw.Format("060102")
		}
	}
	return strDatenw
}
