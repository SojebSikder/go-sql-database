package execution

import (
	"fmt"
	"sojebsql/parser"
	"sojebsql/storage"
)

// Executor executes SQL queries
type Executor struct {
	Engine *storage.StorageEngine
}

// NewExecutor creates a new query executor
func NewExecutor(engine *storage.StorageEngine) *Executor {
	return &Executor{Engine: engine}
}

// Execute runs a given query
func (e *Executor) Execute(query *parser.Query) {
	switch query.Type {
	case "CREATE":
		e.Engine.CreateTable(query.Table, query.Columns)
		fmt.Println("Table created:", query.Table)

	case "INSERT":
		values := make([]interface{}, len(query.Values))
		for i, v := range query.Values {
			values[i] = v
		}
		e.Engine.InsertIntoTable(query.Table, values)
		fmt.Println("Row inserted into:", query.Table)

	case "SELECT":
		rows := e.Engine.SelectFromTable(query.Table)
		fmt.Println("Data from table:", query.Table)
		for _, row := range rows {
			fmt.Println(row)
		}
	}
}
