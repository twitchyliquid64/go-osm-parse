package main

import "database/sql"
import "fmt"

var tableCreateCmd = `
	CREATE TABLE IF NOT EXISTS Nodes
	(
		NodeID int64,
		Lat float64,
		Lon float64,
		Changeset int64
	);
	CREATE INDEX IF NOT EXISTS NodesID ON Nodes (NodeID);
	CREATE INDEX IF NOT EXISTS NodesLat ON Nodes (Lat);
	CREATE INDEX IF NOT EXISTS NodesLon ON Nodes (Lon);
	
	
	CREATE TABLE IF NOT EXISTS NodeTags
	(
		NodeID int64,
		Key string,
		Value string
	);
	CREATE INDEX IF NOT EXISTS NodeTagsID ON NodeTags (NodeID);
	
	
	
	
	
	
	
	CREATE TABLE IF NOT EXISTS Ways
	(
		WayID int64,
	);
	CREATE INDEX IF NOT EXISTS WaysID ON Ways (WayID);
	
	
	CREATE TABLE IF NOT EXISTS WayTags
	(
		WayID int64,
		Key string,
		Value string
	);
	CREATE INDEX IF NOT EXISTS WayTagsID ON WayTags (WayID);
	
	
	CREATE TABLE IF NOT EXISTS WayRefs
	(
		WayID int64,
		RefID int64
	);
	CREATE INDEX IF NOT EXISTS WayRefsWayID ON WayRefs (WayID);
	CREATE INDEX IF NOT EXISTS WayRefsRefID ON WayRefs (RefID);
	
`


func initialise_db(db *sql.DB) {
	
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	
	if _, err = tx.Exec(tableCreateCmd); err != nil {
		panic(err)
	}
	
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}


func write_nodes(db *sql.DB, nodeList *[]*Node){
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	
	for x, node := range *nodeList {
		if _, err = tx.Exec("INSERT INTO Nodes (NodeID,Lat,Lon,Changeset) VALUES ($1,$2,$3,$4)", node.Id, node.Lat, node.Lon, node.Changeset); err != nil {
			panic(err)
		}
		
		for _, tag := range node.Tags{
			if _, err = tx.Exec("INSERT INTO NodeTags (NodeID,Key,Value) VALUES ($1,$2,$3)", node.Id, tag.K, tag.V); err != nil {
				panic(err)
			}
		}
		
		if ((x % 10000) == 0) && (x > 1000){
			fmt.Println("Written", x, "/", len(*nodeList), "Nodes to DB (", int(float64(x)/float64(len(*nodeList))*100.0), "percent", ")")
		}
	}
	
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}



func write_ways(db *sql.DB, wayList *[]*Way){
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	
	for x, way := range *wayList {
		if _, err = tx.Exec("INSERT INTO Ways (WayID) VALUES ($1)", way.Id); err != nil {
			panic(err)
		}
		
		for _, tag := range way.Tags{
			if _, err = tx.Exec("INSERT INTO WayTags (WayID,Key,Value) VALUES ($1,$2,$3)", way.Id, tag.K, tag.V); err != nil {
				panic(err)
			}
		}
		
		for _, ref := range way.Refs{
			if _, err = tx.Exec("INSERT INTO WayRefs (WayID,RefID) VALUES ($1,$2)", way.Id, ref.Ref); err != nil {
				panic(err)
			}
		}
		
		if ((x % 1000) == 0) && (x > 100){
			fmt.Println("Written", x, "/", len(*wayList), "Ways to DB (", int(float64(x)/float64(len(*wayList))*100.0), "percent", ")")
		}
	}
	
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

