package fncSbrapi

import (
	mdlSbrapi "back/sbrapi/model"
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

// MessageHeader
func FncSbrapiMsghdrMainob(pcc, srv, act string, hdr mdlSbrapi.MdlSbrapiMsghdrParams,
) mdlSbrapi.MdlSbrapiMsghdrMainob {
	return mdlSbrapi.MdlSbrapiMsghdrMainob{
		From:           mdlSbrapi.MdlSbrapiMsghdrPartyx{PartyId: "LionAir"},
		To:             mdlSbrapi.MdlSbrapiMsghdrPartyx{PartyId: "Sabre"},
		CPAId:          pcc,
		ConversationId: hdr.Convid,
		Service:        mdlSbrapi.MdlSbrapiMsghdrSrvice{Type: "sabreXML", Name: srv},
		Action:         act,
		MessageData: mdlSbrapi.MdlSbrapiMsghdrMsgdta{
			MessageId: hdr.Mssgid,
			Timestamp: hdr.Timefm,
		},
		XmlnsEb:        "http://www.ebxml.org/namespaces/messageHeader",
		MustUnderstand: "1",
		Version:        "1.0",
	}
}

// Function Treatment XML data hit API
func FncSbrapiMsghdrXmldta(bdyxml interface{}) ([]byte, error) {

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
