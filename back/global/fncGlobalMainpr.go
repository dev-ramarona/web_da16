package fncGlobal

import (
	"bufio"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func FncGlobalMainprLoadnv(filenm string) {
	filenw, err := os.Open(filenm)
	if err != nil {
		panic("Error opening .env file:" + err.Error())
	}
	defer filenw.Close()

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
var (
	Client *mongo.Client
	Status = "Done"
	// Dbases = os.Getenv("VAR_DTBASE")
	// jwtkey = []byte(os.Getenv("JWT_SECRET"))
	// Urlmgo = os.Getenv("URI_MONGOS")
	// Pcckey = os.Getenv("SBR_PCCKEY")
	// Usrnme = os.Getenv("SBR_USRNME")
	// Psswrd = os.Getenv("SBR_PSSWRD")
	// Ipgolg = os.Getenv("IPV_GOLANG")
	// Ipadrs = os.Getenv("IPV_ALLOWD")
	// Ipalow = strings.Split(os.Getenv("IPV_ALLOWD"), "|")
)
