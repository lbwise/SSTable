package database

import (
    "errors"
    "fmt"
    "os"
    "strings"
)

type DB struct {
    Size int
    capacity int
    index *Index
    table *os.File
    blockSize int
    //indexes []*Index
}


func CreateDatabase(capacity int, blockSize int) *DB {
    return &DB{
        capacity: capacity,
        blockSize: blockSize,
    }
}

func (db *DB) CreateIndex(key string) error {
    db.index = CreateIndex(key, db.blockSize)
    return nil
}

func (db *DB) Write(key string, value any) error {
    log := fmt.Sprintf("%s,%s;", key, value)
    return db.writeTable(log)
}

func (db *DB) Read(key string) (any, error) {
    loc, err := db.index.Search(key)
    if err != nil {
        return nil, err
    }

    table, err := db.loadTable()
    if err != nil {
        return nil, err
    }

    value, err := db.parse(key, table, loc)
    if err != nil {
        return nil, err
    }

    return value, nil
}


func (db *DB) parse(key string, table string, index int) (string, error) {
    keyStart := -1
    readLine := false
    for i := index; i < index + db.blockSize; i++ {
        if keyStart == -1 && table[i] == key[0] {
            keyStart = i
        } else if keyStart != -1 && table[keyStart: i] == key + ","{
            readLine = true
            // Reads from first occurance of character (fuck it switch to regex)
        } else if readLine && string(table[i]) == ";" {
            return strings.Split(table[keyStart:i], ",")[1], nil
        }
    }
    return "", errors.New("couldn't find key")
}

func (db *DB) Merge() {
}

func (db *DB) loadTable() (string, error) {
    f, err := os.OpenFile("database.txt", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return "", err
    }

    tableBytes := make([]byte, db.capacity)
    f.Read(tableBytes)
    table := fmt.Sprintf("%s", tableBytes)
    return table, nil
}

func (db *DB) writeTable(log string) error {
    f, _ := os.OpenFile("database.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    _, err := f.WriteString(log)
    return err
}

