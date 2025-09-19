package fnc_global

import (
	mdl_global "back/global/model"
	"encoding/xml"
)

func FncGlobalPnrallApisbr(tokenx, pnrcde string, sbarea []string) mdl_global.PnrallReqEnvlpe {
	bdyPnrall := mdl_global.PnrallReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.PnrallReqHeader{
			MessageHeader: FncGlobalApisbrMsghdr(Pcckey, "Retrieve Itinerary", "GetReservationRQ"),
			Security: mdl_global.PnrallReqSecrty{
				BinarySecurityToken: mdl_global.PnrallReqBsctkn{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: tokenx,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.PnrallReqBodyxx{
			GetReservationRQ: mdl_global.PnrallReqGetRsvtion{
				Xmlns:       "http://webservices.sabre.com/pnrbuilder/v1_19",
				Version:     "1.19.0",
				Locator:     pnrcde,
				RequestType: "Stateless",
				ReturnOptions: mdl_global.PnrallReqRtropt{
					SubjectAreas:   mdl_global.PnrallReqSbjare{SubjectArea: sbarea},
					ViewName:       "Simple",
					ResponseFormat: "STL",
				},
			},
		},
	}
	return bdyPnrall
}

// Login Sabre and create Session API
func FncGlobalItnaryApisbr(tokenx, pnrcde string) ([]mdl_global.ItnaryRspSegmnt, mdl_global.RemarksRspBokdtl, error) {

	// Isi struktur data
	bdyPnrall := FncGlobalPnrallApisbr(tokenx, pnrcde, []string{"ITINERARY"})

	// Declare first output
	var bstRspnse []mdl_global.ItnaryRspSegmnt
	var bokRspnse mdl_global.RemarksRspBokdtl

	// Read response
	rspSession, err := FncGlobalApisbrXmldta(bdyPnrall)
	if err != nil {
		return bstRspnse, bokRspnse, err
	}

	// Parsing XML ke dalam struktur Go
	var envlpeRspnse mdl_global.ItnaryRspEnvlpe
	err = xml.Unmarshal([]byte(rspSession), &envlpeRspnse)
	if err != nil {
		return bstRspnse, bokRspnse, err
	}

	// Return non error data
	bstRspnse = envlpeRspnse.Body.GetReservationRS.Reservation.PassengerReservation.Segments.Segment
	bokRspnse = envlpeRspnse.Body.GetReservationRS.Reservation.BookingDetails
	return bstRspnse, bokRspnse, nil
}

// Login Sabre and create Session API
func FncGlobalRmarksApisbr(tokenx, pnrcde string) (mdl_global.RemarksRspBokdtl, error) {

	// Isi struktur data
	bdyPnrall := FncGlobalPnrallApisbr(tokenx, pnrcde, []string{"REMARKS"})

	// Declare first output
	var bstRspnse mdl_global.RemarksRspBokdtl

	// Read response
	rspSession, err := FncGlobalApisbrXmldta(bdyPnrall)
	if err != nil {
		return bstRspnse, err
	}

	// Parsing XML ke dalam struktur Go
	var envlpeRspnse mdl_global.RemarksRspEnvlpe
	err = xml.Unmarshal([]byte(rspSession), &envlpeRspnse)
	if err != nil {
		return bstRspnse, err
	}

	// Return non error data
	bstRspnse = envlpeRspnse.Body.GetReservationRS.Reservation.BookingDetails
	return bstRspnse, nil
}

// Login Sabre and create Session API
func FncGlobalItrmrlApisbr(tokenx, pnrcde string) (mdl_global.ItrmrlRspRsvtion, error) {

	// Isi struktur data
	bdyPnrall := FncGlobalPnrallApisbr(tokenx, pnrcde, []string{"REMARKS", "ITINERARY", "RECORD_LOCATOR"})

	// Declare first output
	var bstRspnse mdl_global.ItrmrlRspRsvtion

	// Read response
	rspSession, err := FncGlobalApisbrXmldta(bdyPnrall)
	if err != nil {
		return bstRspnse, err
	}

	// Parsing XML ke dalam struktur Go
	var envlpeRspnse mdl_global.ItrmrlRspEnvlpe
	err = xml.Unmarshal([]byte(rspSession), &envlpeRspnse)
	if err != nil {
		return bstRspnse, err
	}

	// Return non error data
	bstRspnse = envlpeRspnse.Body.GetReservationRS.Reservation
	return bstRspnse, nil
}
