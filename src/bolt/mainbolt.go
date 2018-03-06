package main

import (
	"fmt"
)

func main() {

	fmt.Println("Bolt  ")
	
	dbOpen("myboltdb.db")
	defer dbClose()
	
	
	fmt.Println("Bolt db path = " + dbPath() )
	
	fmt.Println("Bolt - creating bucket 'countries'... ")
	dbCreateBucket("countries");

	fmt.Println("Bolt - set countries ... ")
	dbPut("countries", "FR", "France")
	dbPut("countries", "JP", "Japan")
	dbPut("countries", "PL", "Poland")

	fmt.Println("Bolt - get countries ... ")
	fmt.Println(" 'PL' : " + dbGet("countries", "PL") )
	fmt.Println(" 'FR' : " + dbGet("countries", "FR") )
	fmt.Println(" 'ZZ' : " + dbGet("countries", "ZZ") )
	if  dbGet("countries", "ZZ") == "" {
		fmt.Println(" 'ZZ' is void => not found "  )
	}
	
	fmt.Println(" 'JP' : " + dbGet("countries", "JP") )
	fmt.Println("Bolt - delete JP ... ")
	dbDelete("countries", "JP")
	fmt.Println(" 'JP' : " + dbGet("countries", "JP") )
	
	fmt.Println("Bolt - new 'foo' bucket... ")
	dbCreateBucket("foo");
	dbPut("foo", "FR", "France in foo")
	fmt.Println(" 'FR' in countries : " + dbGet("countries", "FR") )
	fmt.Println(" 'FR' in foo       : " + dbGet("foo", "FR") )
	
	dbPut("countries", "ES", "Spain")
	dbPut("countries", "BE", "Belgium")
	fmt.Println("Bolt - get all from 'countries' bucket... ")
	countries := dbGetAll("countries")
	for i, country := range countries {
		fmt.Printf(" %d : %+v \n", i, country )
	}
	
}