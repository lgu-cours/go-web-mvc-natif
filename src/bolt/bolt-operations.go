package main

import (
	"github.com/boltdb/bolt"
)

// The current Bolt Database
var db *bolt.DB

func dbOpen(fileName string) {
	// Open the Bolt data file in the current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open(fileName, 0600, nil)
	if err != nil {
		panic("Cannot open database. File '" + fileName + "' ")
	}	
}

func dbCheckIfOpen() {
	if db == nil {
		panic("Bolt database is not open!")
	}
}

func dbClose() {
	if db != nil {
		db.Close()
	}
}

func dbPath() string {
	dbCheckIfOpen()
	return db.Path()
}


func dbCreateBucket(bucketName string) {
	dbCheckIfOpen()
	err := db.Update(func(tx *bolt.Tx) error {
	    //_, err := tx.CreateBucket( []byte(bucketName) ) 
	    _, err := tx.CreateBucketIfNotExists( []byte(bucketName) )       
	    if err != nil {
	        return err
	    }
	    return nil
	});	
	
	if err != nil {
		panic("Cannot create bucket '" + bucketName + "' ")
	}
}

func dbPut(bucketName string, key string, value string ) {
	dbCheckIfOpen()
	
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    bucket.Put([]byte(key), []byte(value))                        
	    return nil
	});	
	
	if err != nil {
		panic("Cannot put key '" + key + "' in bucket '" + bucketName + "'")
	}
}

func dbGet(bucketName string, key string ) string {
	dbCheckIfOpen()
	
	var value string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    value = string(bucket.Get([]byte(key))) // void string if value == nil (not found)      
	    return nil
	});	
	
	
	if err != nil {
		panic("Cannot get key '" + key + "' from bucket '" + bucketName + "'")
	}
	return value
}

func dbDelete(bucketName string, key string ) {
	dbCheckIfOpen()
	
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    bucket.Delete([]byte(key))                                                     
	    return nil
	});	
	
	if err != nil {
		panic("Cannot delete key '" + key + "' in bucket '" + bucketName + "'")
	}
}

