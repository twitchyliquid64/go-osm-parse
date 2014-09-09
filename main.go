package main

import (
	"database/sql"
	_ "github.com/cznic/ql/driver"
	"encoding/xml"
	"fmt"
	"os"
)


//result of the parse operation
var bounds Bounds
var nodes []*Node
var ways []*Way
var relations []*Relation

func decode(decoder *xml.Decoder){
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
					//fmt.Println(inElement)
				}
		}

	}
}


func main() {
	xmlFile, err := os.Open("input.osm")//DISABLE YOUR AV CUZ ITS A FUCKIN BIG FILE AND AVs LIKE TO SCAN FILES BEFORE IT LETS THEM OPEN!!!
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	db, err := sql.Open("ql", "osmdump.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	initialise_db(db)
	defer db.Close()
	
	
	decoder := xml.NewDecoder(xmlFile)
	fmt.Println("Decoding file to memory...Please Wait.")
	decode(decoder)
	fmt.Println("Decode completed.")

	write_ways(db, &ways)
	write_nodes(db, &nodes)
	fmt.Println("Relations:", len(relations))
	
}
