package fncGlobal

import (
	"bufio"
	"os"
	"strings"

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
var Status = "Done"
var Client *mongo.Client
var jwtkey []byte
var Ipalow []string
var Dbases, Urlmgo, Pcckey, Usrnme,
	Psswrd, Ptgolg, Ipadrs string

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
	Ipalow = strings.Split(os.Getenv("NEXT_PUBLIC_IPV_ALLOWD"), "|")
}
