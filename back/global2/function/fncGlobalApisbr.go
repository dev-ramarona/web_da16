package fnc_global

import (
	mdl_global "back/global/model"
	"bytes"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get status data process
func FncGlobalApisbrStatus(c *gin.Context) {
	c.JSON(http.StatusOK, IsStat)
}

// MessageHeader
func FncGlobalApisbrMsghdr(pcc string, serv string, actn string) mdl_global.Msghdr {
	MessageHeader := mdl_global.Msghdr{
		From:           mdl_global.Partyx{PartyId: "LionAir"},
		To:             mdl_global.Partyx{PartyId: "Sabre"},
		CPAId:          Pcckey,
		ConversationId: "V1@280b16ec-5eac-46c0-893f-c88f8e8cb632@310b16ecxxxyz",
		Service:        mdl_global.Service{Type: "sabreXML", Name: serv},
		Action:         actn,
		MessageData: mdl_global.Msgdta{
			MessageId: "mid:20001209-133003-2333@clientofsabre.com",
			Timestamp: "Z",
		},
		XmlnsEb:        "http://www.ebxml.org/namespaces/messageHeader",
		MustUnderstand: "1",
		Version:        "1.0",
	}
	return MessageHeader
}

// Function Treatment XML data hit API
func FncGlobalApisbrXmldta(bdyxml interface{}) ([]byte, error) {

	// Convert structure to XML
	xmlAlldta, err := xml.MarshalIndent(bdyxml, "", "  ")
	if err != nil {
		return []byte{}, err
	}

	// Send request to API
	xmlAlldta = append([]byte(xml.Header), xmlAlldta...)
	rspAlldta, err := http.Post("https://webservices.platform.sabre.com",
		"text/xml", bytes.NewBuffer(xmlAlldta))
	if err != nil {
		return []byte{}, err
	}
	defer rspAlldta.Body.Close()

	// Read response
	fnlAlldta, err := io.ReadAll(rspAlldta.Body)
	if err != nil {
		return []byte{}, err
	}

	// Response byte and error
	return fnlAlldta, nil
}
