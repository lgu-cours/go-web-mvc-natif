package main

import (
	"fmt"
)

func main() {

	fmt.Println("Bolt  ")
	
	/*
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open("myboltdb.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	*/
	dbOpen("myboltdb.db")
	defer dbClose()
	
	
	fmt.Println("Bolt db path = " + dbPath() )
	
	fmt.Println("Bolt - creating bucket 'countries'... ")
	dbCreateBucket("countries");

	fmt.Println("Bolt - set countries ... ")
//	dbPut("countries", "FR", "France")
//	dbPut("countries", "JP", "Japan")
//	dbPut("countries", "PL", "Poland")

	fmt.Println("Bolt - get countries ... ")
	fmt.Println(" 'PL' : " + dbGet("countries", "PL") )
	fmt.Println(" 'FR' : " + dbGet("countries", "FR") )
	fmt.Println(" 'ZZ' : " + dbGet("countries", "ZZ") )
	if  dbGet("countries", "ZZ") == "" {
		fmt.Println(" 'ZZ' is void => not found "  )
	}
	
	dbCreateBucket("foo");
	dbPut("foo", "FR", "France in foo")
	fmt.Println(" 'FR' in countries : " + dbGet("countries", "FR") )
	fmt.Println(" 'FR' in foo       : " + dbGet("foo", "FR") )
	
}