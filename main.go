package main

import (
    "fmt"
    "github.com/lbwise/SSTable/database"
)

type Metadata struct {
    id string
    counter int
    initialisation int
}

func main() {
    db := database.CreateDatabase(100, 10)
    db.Write("name", "Liam")
    db.Write("age", 20)

    metadata := Metadata{ id: "1234abcd", initialisation: 0, counter: 20 }
    db.Write("metadata", metadata)

    name, err := db.Read("name")
    if err != nil {
       fmt.Println("Error: ", err)
    }
    fmt.Println("Name:", name)

    age, err := db.Read("age")
    if err != nil {
       fmt.Println("Error: ", err)
    }
    fmt.Println("Age: ", age)

    md, err := db.Read("metadata")
    if err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Println("Metadata: ", md)
}
