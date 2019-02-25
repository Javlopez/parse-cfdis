package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Javlopez/cfdi-parser/internals/cfdi"
)

func main() {
	data := [][]string{}

	dir := flag.String("dir", "", "Por favor ingresa el nombre del directorio")
	output := flag.String("output", "output.csv", "Por favor ingresa el destino de tu archivo csv")

	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: --dir=dir123.zip\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *dir == "" {
		usage()
		os.Exit(1)
	}

	zipDir := *dir
	outputFile := *output

	r, err := zip.OpenReader(zipDir)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	data = append(data, []string{"Folio", "Fecha", "Forma de Pago", "Emisor", "Concepto", "Subtotal", "Impuestos", "Total"})

	for _, f := range r.File {

		rc, _ := f.Open()
		xmlData := ReadXML(rc)
		c := cfdi.Comprobante{Folio: "none"}
		err := xml.Unmarshal(xmlData, &c)

		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		data = append(data, []string{c.Folio, c.Fecha, c.FormaDePago, c.Emisor.Name, c.Conceptos.Descripcion, c.SubTotal, c.Impuestos.Total, c.Total})
	}

	WriteCsvFile(outputFile, data)
}

//ReadXML method
func ReadXML(handle io.Reader) []byte {
	byteValue, _ := ioutil.ReadAll(handle)
	return byteValue
}

func WriteCsvFile(name string, data [][]string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal("Cannot create csv file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to csv file", err)
		}
	}
}
