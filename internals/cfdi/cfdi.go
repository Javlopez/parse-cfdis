package cfdi

import (
	"encoding/xml"
)

type Emisor struct {
	XMLName xml.Name `xml:"Emisor"`
	Name    string   `xml:"Nombre,attr"`
}

type Receptor struct {
	XMLName xml.Name `xml:"Receptor"`
	Rfc     string   `xml:"Rfc,attr"`
	Uso     string   `xml:"UsoCFDI,attr"`
}

type Impuestos struct {
	XMLName xml.Name `xml:"Impuestos"`
	Total   string   `xml:"TotalImpuestosTrasladados,attr"`
}

type Concepto struct {
	XMLName     xml.Name `xml:"Concepto"`
	Descripcion string   `xml:"Descripcion,attr"`
}

//Comprobante
type Comprobante struct {
	XMLName     xml.Name `xml:"Comprobante"`
	Folio       string   `xml:"Folio,attr"`
	Fecha       string   `xml:"Fecha,attr"`
	FormaDePago string   `xml:"FormaPago,attr"`
	SubTotal    string   `xml:"SubTotal,attr"`
	Total       string   `xml:"Total,attr"`
	Emisor      Emisor
	Receptor    Receptor
	Conceptos   Concepto `xml:"Conceptos>Concepto"`
	Impuestos   Impuestos
}
