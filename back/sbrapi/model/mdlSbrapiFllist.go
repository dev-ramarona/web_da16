package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiFllistReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiFllistReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiFllistReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiFllistReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiFllistReqscr `xml:"wsse:Security"`
}
type MdlSbrapiFllistReqscr struct {
	BinarySecurityToken MdlSbrapiFllistReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiFllistReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiFllistReqbdy struct {
	ACS_AirportFlightListRQ MdlSbrapiFllistReqafl `xml:"v3:ACS_AirportFlightListRQ"`
}
type MdlSbrapiFllistReqafl struct {
	XMLName    xml.Name              `xml:"v3:ACS_AirportFlightListRQ"`
	Xmlns      string                `xml:"xmlns:v3,attr"`
	FlightInfo MdlSbrapiFllistReqinf `xml:"FlightInfo"`
}
type MdlSbrapiFllistReqinf struct {
	Airline            string                `xml:"Airline"`
	DepartureDate      string                `xml:"DepartureDate"`
	Origin             string                `xml:"Origin"`
	DepartureTimeRange MdlSbrapiFllistReqdtr `xml:"DepartureTimeRange"`
}
type MdlSbrapiFllistReqdtr struct {
	StartTime string `xml:"StartTime"`
	EndTime   string `xml:"EndTime"`
}

// Response
type MdlSbrapiFllistRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiFllistRspbdy `xml:"Body"`
}
type MdlSbrapiFllistRspbdy struct {
	ACS_AirportFlightListRS MdlSbrapiFllistRspfls `xml:"http://services.sabre.com/ACS/BSO/airportFlightList/v3 ACS_AirportFlightListRS"`
}
type MdlSbrapiFllistRspfls struct {
	Origin            string                `xml:"Origin"`
	AirportFlightList MdlSbrapiFllistRspafl `xml:"AirportFlightList"`
}
type MdlSbrapiFllistRspafl struct {
	AirportFlight []MdlSbrapiFllistRspapf `xml:"AirportFlight"`
}
type MdlSbrapiFllistRspapf struct {
	Flight           string `xml:"Flight"`
	DepartureDate    string `xml:"DepartureDate"`
	DepartureTime    string `xml:"DepartureTime"`
	DepartureGate    string `xml:"DepartureGate"`
	Destination      string `xml:"Destination"`
	DestinationFinal string `xml:"DestinationFinal"`
	Status           string `xml:"Status"`
	AircraftType     string `xml:"AircraftType"`
	FreeText         string `xml:"FreeText"`
}
