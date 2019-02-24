package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func main() {
	xmlData := parseFile("0AF59B69-355A-4407-980E-0A240A18D355.xml")
	c := Comprobante{Folio: "none"}
	err := xml.Unmarshal(xmlData, &c)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("XMLName: %#v\n", c.XMLName)
	fmt.Printf("Folio: %q\n", c.Folio)
	fmt.Printf("Fecha: %q\n", c.Fecha)
	fmt.Printf("FormaDePago: %q\n", c.FormaDePago)
	fmt.Printf("Total: %q\n", c.Total)
	fmt.Printf("SubTotal: %q\n", c.SubTotal)
	fmt.Printf("Emisor: %q\n", c.Emisor.Name)
	fmt.Printf("Receptor: %q\n", c.Receptor.Rfc)
	fmt.Printf("Conceptos: %q\n", c.Conceptos.Descripcion)
	fmt.Printf("Impuestos: %q\n", c.Impuestos.Total)
}

func parseFile(file string) []byte {
	handle, err := os.Open(file)

	if err != nil {
		return []byte{}
	}

	defer handle.Close()
	return ReadXML(handle)
}

//ReadXML method
func ReadXML(handle io.Reader) []byte {
	byteValue, _ := ioutil.ReadAll(handle)
	return byteValue
}
