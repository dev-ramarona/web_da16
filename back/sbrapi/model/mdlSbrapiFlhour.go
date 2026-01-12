package mdlSbrapi

import "encoding/xml"

type MdlSbrapiFlhourRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiFlhourRspbdy `xml:"Body"`
}

type MdlSbrapiFlhourRspbdy struct {
	SabreCommandLLSRS MdlSbrapiFlhourRspcmd `xml:"http://webservices.sabre.com/sabreXML/2011/10 SabreCommandLLSRS"`
}

type MdlSbrapiFlhourRspcmd struct {
	XML_Content MdlSbrapiFlhourRspxml `xml:"XML_Content"`
}

type MdlSbrapiFlhourRspxml struct {
	AIRAALSADSKED0 MdlSbrapiFlhourRspair `xml:"AIRAALSADSKED0"`
}

type MdlSbrapiFlhourRspair struct {
	SKD001 []MdlSbrapiFlhourRspskd `xml:"SKD001"`
}

type MdlSbrapiFlhourRspskd struct {
	TNSCarrierCode           string `xml:"TNSCarrierCode"`
	DateOfDDMM               string `xml:"dateOfDDMM"`
	BoardPoint               string `xml:"boardPoint"`
	DestinationAirportCode   string `xml:"destinationAirportCode"`
	AirportZoneCode          string `xml:"airportZoneCode"`
	ScheduledDepartureTime   string `xml:"scheduledDepartureTime"`
	TimeZoneCode             string `xml:"timeZoneCode"`
	ArrivalTime              string `xml:"arrivalTime"`
	MealCode1                string `xml:"mealCode1"`
	EquipmentCode            string `xml:"equipmentCode"`
	ElapsedTime              string `xml:"elapsedTime"`
	AccumulatedElapsedTime   string `xml:"accumulatedElapsedTime"`
	IATA_NonSmokingIndicator string `xml:"IATA_NonSmokingIndicator"`
	AirMilesFlown            string `xml:"airMilesFlown"`
	FlightNumber             string `xml:"flightNumber"`
}
