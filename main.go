package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	MinLat string `xml:"minlat,attr"`
	MaxLat string `xml:"maxlat,attr"`
	
	MinLon string `xml:"minlon,attr"`
	MaxLon string `xml:"maxlon,attr"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

type Ref struct {
	XMLName xml.Name `xml:"nd"`
	Ref string `xml:"ref,attr"`
}

type Member struct {
	XMLName xml.Name `xml:"member"`
	Type string `xml:"type,attr"`
	Ref string `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}


type Way struct {
	XMLName xml.Name `xml:"way"`
	Id string `xml:"id,attr"`
	Tags []Tag `xml:"tag"`
	Refs []Ref `xml:"nd"`
}


type Node struct {
	XMLName xml.Name `xml:"node"`
	Id string `xml:"id,attr"`
	Lat string `xml:"lat,attr"`
	Lon string `xml:"lon,attr"`
	Changeset string `xml:"changeset,attr"`
	Tags []Tag `xml:"tag"`
}

type Relation struct {
	XMLName xml.Name `xml:"relation"`
	Id string `xml:"id,attr"`
	Changeset string `xml:"changeset,attr"`
	Tags []Tag `xml:"tag"`
	Members []Member `xml:"member"`
}


//result of the parse operation
var bounds Bounds
var nodes []*Node
var ways []*Way
var relations []*Relation

func main() {
	xmlFile, err := os.Open("input.osm")//DISABLE YOUR AV CUZ ITS A FUCKIN BIG FILE AND AVs LIKE TO SCAN FILES BEFORE IT LETS THEM OPEN!!!
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	
	decoder := xml.NewDecoder(xmlFile)

	var inElement string
	
	for {
		t, _ := decoder.Token()// Read tokens from the XML document in a stream.
		if t == nil {
			break
		}

		switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local

				if inElement == "bounds" {
					decoder.DecodeElement(&bounds, &se)
					fmt.Println( "Latitude:", bounds.MinLat, "to", bounds.MaxLat, "- Longitude:", bounds.MinLon, "to", bounds.MaxLon)
				}else if inElement == "node" {
					var node Node
					decoder.DecodeElement(&node, &se)
					nodes = append(nodes, &node)
					//fmt.Println(node)
				}else if inElement == "way" {
					var way Way
					decoder.DecodeElement(&way, &se)
					ways = append(ways, &way)
					//fmt.Println(way)
				}else if inElement == "relation" {
					var relation Relation
					decoder.DecodeElement(&relation, &se)
					relations = append(relations, &relation)
					//fmt.Println(relation)
				}else {
					fmt.Println(inElement)
				}
		}

	}
	
	//for i, way := range relations {
	//		fmt.Println(i, way)
	//}
	
	fmt.Println("Nodes:    ", len(nodes))
	fmt.Println("Ways:     ", len(ways))
	fmt.Println("Relations:", len(relations))
	
}
