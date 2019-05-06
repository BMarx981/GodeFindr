package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var count int

//Node : contains the nodes of the payload
type Node struct {
	XMLName xml.Name
	Nodes   []Node `xml:",any"`
	Content []byte `xml:"innerxml"`
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
	// processXML(xmlFile)

	buf := bytes.NewBuffer(xmlFile)
	dec := xml.NewDecoder(buf)

	var n Node
	err = dec.Decode(&n)
	if err != nil {
		panic(err)
	}

	walk([]Node{n}, func(n Node) bool {
		if n.XMLName.Local == "RECORD.KEY" {
			fmt.Println(string(n.Content))
		}
		return true
	})
	fmt.Println("The End********************")
} // /Users/brianmarx/go/src/NodeFindr/baselineFake.txt

//UnmarshalXML : the XML nodes.
func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type node Node
	count = count + 1
	fmt.Println("called ", start.End)
	return d.DecodeElement((*node)(n), &start)

}

func walk(nodes []Node, f func(Node) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}
