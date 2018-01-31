package main

type Student struct {
        //exported field 'Name'
        FirstName string
        LastName  string
        Age       int 
        secretCode string // not exported ('private')
        Teacher Teacher
}
