package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Root : The main doc
type Root struct {
	XMLName   xml.Name    `xml:"root"`
	DataInput []DataInput `xml:"dataInput"`
}

//DataInput : contain the meat of the data
type DataInput struct {
	XMLName xml.Name `xml:"dataInput"`
	Hdr     []Hdr    `xml:"hdr"`
	Payload Payload  `xml:"payload"`
}

//Hdr : has the top level data
type Hdr struct {
	XMLName xml.Name `xml:"hdr"`
	Values  []Value  `xml:",any"`
}

//Payload : Contains all of the data
type Payload struct {
	XMLName xml.Name `xml:"payload"`
	Record  Record   `xml:"record"`
}

//Record : Contains all of the very important data
type Record struct {
	XMLName xml.Name `xml:"record"`
	Values  []string `xml:",any"`
	RecKey  string   `xml:"RECORD.KEY"`
}

//Value : contains the nodes of the payload
type Value struct {
	Node    string
	Content string
}

func main() {
	processXML(findFile())
	fmt.Println("The End********************")
} // /Users/brianmarx/Desktop/baselineFake.txt

func findFile() []byte {
	fmt.Print("Enter the file name: ")
	reader := bufio.NewReader(os.Stdin)

	fileName, _, error := reader.ReadLine()
	if error != nil {
		fmt.Println("Error reading file")
	}

	xmlFile, err := ioutil.ReadFile(string(fileName))
	if err != nil {
		fmt.Println("Error opening the xml file", err)
	}
	return xmlFile
}

//ProcessXML : the XML nodes.
func processXML(xmlFile []byte) {
	var root Root

	error := xml.Unmarshal(xmlFile, &root)
	if error != nil {
		fmt.Println(error)
	}
	processDataInput(root.DataInput, xmlFile)
}

func processDataInput(input []DataInput, xmlFile []byte) error {
	m := make(map[string]DataInput)
	for index, element := range input {
		fmt.Println(element.Payload.Record.RecKey)
		m[element.Payload.Record.RecKey] = input[index]
	}

	fmt.Println(m)
	return nil
}
