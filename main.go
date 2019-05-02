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
	Program Program  `xml:"program>,any"`
	Payload Payload  `xml:"payload>,any"`
}

//Program : has the top level data
type Program struct {
	XMLName xml.Name `xml:"program"`
	Values  []Value  `xml:",any"`
}

//Payload : Contains all of the data
type Payload struct {
	XMLName xml.Name `xml:"payload"`
	Values  []Value  `xml:",any"`
}

//Value : contains the nodes of the payload
type Value struct {
	Node    string
	Content string
}

func main() {
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
	processXML(xmlFile)
	fmt.Println("The End********************")
} // /Users/brianmarx/Desktop/baselineFake.txt

//ProcessXML : the XML nodes.
func processXML(xmlFile []byte) {
	var root Root
	var prog Program
	xml.Unmarshal(xmlFile, &prog)
	xml.Unmarshal(xmlFile, &root)
	progs := prog.Values
	// inputs := root.DataInput
	fmt.Println(len(progs))
}
