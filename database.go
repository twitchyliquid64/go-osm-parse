package main

import "database/sql"

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
	
	for _, node := range *nodeList {
		if _, err = tx.Exec("INSERT INTO Nodes (NodeID,Lat,Lon,Changeset) VALUES ($1,$2,$3,$4)", node.Id, node.Lat, node.Lon, node.Changeset); err != nil {
			panic(err)
		}
		
		for _, tag := range node.Tags{
			if _, err = tx.Exec("INSERT INTO NodeTags (NodeID,Key,Value) VALUES ($1,$2,$3)", node.Id, tag.K, tag.V); err != nil {
				panic(err)
			}
		}
	}
	
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}
