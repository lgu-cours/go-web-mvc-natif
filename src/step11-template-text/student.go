package main

type Student struct {
        //exported field 'Name'
        FirstName string
        LastName string
        secretCode string // not exported ('private')
}
