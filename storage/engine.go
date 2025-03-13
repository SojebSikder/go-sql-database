package storage

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
)

// Table structure
type Table struct {
	Name    string
	Columns []string
	Rows    [][]interface{}
}

// StorageEngine handles persistent storage
type StorageEngine struct {
	mu     sync.RWMutex
	Tables map[string]*Table
}

// NewStorageEngine initializes the storage engine
func NewStorageEngine() *StorageEngine {
	return &StorageEngine{
		Tables: make(map[string]*Table),
	}
}

// CreateTable creates a new table
func (se *StorageEngine) CreateTable(name string, columns []string) {
	se.mu.Lock()
	defer se.mu.Unlock()

	table := &Table{Name: name, Columns: columns, Rows: [][]interface{}{}}
	se.Tables[name] = table
	se.saveTable(name)
}

// InsertIntoTable adds a new row
func (se *StorageEngine) InsertIntoTable(name string, values []interface{}) {
	se.mu.Lock()
	defer se.mu.Unlock()

	table, exists := se.Tables[name]
	if !exists {
		fmt.Println("Table does not exist.")
		return
	}

	table.Rows = append(table.Rows, values)
	se.saveTable(name)
}

// SelectFromTable retrieves rows
func (se *StorageEngine) SelectFromTable(name string) [][]interface{} {
	se.mu.RLock()
	defer se.mu.RUnlock()

	se.loadTable(name)
	if table, exists := se.Tables[name]; exists {
		return table.Rows
	}
	return nil
}

// saveTable persists a table to disk
func (se *StorageEngine) saveTable(name string) {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		fmt.Println("Error creating data directory:", err)
		return
	}

	file, err := os.Create("data/" + name + ".table")
	if err != nil {
		fmt.Println("Error creating table file:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(se.Tables[name])
	if err != nil {
		fmt.Println("Error encoding table:", err)
	}
}

// loadTable reads a table from disk
func (se *StorageEngine) loadTable(name string) {
	filePath := "data/" + name + ".table"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening table file:", err)
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var table Table
	err = decoder.Decode(&table)
	if err != nil {
		fmt.Println("Error decoding table:", err)
		return
	}

	se.Tables[name] = &table
}
