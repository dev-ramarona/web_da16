package fncGlobal

import (
	mdlGlobal "back/global/model"
	"bufio"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"strings"

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
}

// Get status data process
func FncGlobalMainprStatus(c *gin.Context) {
	c.JSON(http.StatusOK, Status)
}

// Generate UUID
func FncGlobalMakevrCduuid() string {
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
