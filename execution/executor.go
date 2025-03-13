package execution

import (
	"fmt"
	"sojebsql/parser"
	"sojebsql/storage"
	"strconv"
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
		table := e.Engine.Tables[query.Table]

		if table == nil {
			fmt.Println("Table does not exist.")
			return
		}

		// If WHERE condition is provided, filter rows
		if query.WhereCol != "" && query.WhereOp != "" && query.WhereVal != "" {
			colIndex := -1
			for i, col := range table.Columns {
				if col == query.WhereCol {
					colIndex = i
					break
				}
			}

			if colIndex == -1 {
				fmt.Println("Column not found in WHERE clause.")
				return
			}

			// Convert WHERE value to appropriate type
			var filterVal interface{} = query.WhereVal
			if num, err := strconv.Atoi(query.WhereVal); err == nil {
				filterVal = num
			}

			// Filter rows based on operator
			var filteredRows [][]interface{}
			for _, row := range rows {
				if compare(row[colIndex], filterVal, query.WhereOp) {
					filteredRows = append(filteredRows, row)
				}
			}

			fmt.Println("Filtered data from table:", query.Table)
			for _, row := range filteredRows {
				fmt.Println(row)
			}
		} else {
			fmt.Println("All data from table:", query.Table)
			for _, row := range rows {
				fmt.Println(row)
			}
		}
	}
}

// compare function applies WHERE filtering based on operator
func compare(value interface{}, filter interface{}, operator string) bool {
	switch v := value.(type) {
	case int:
		filterNum, ok := filter.(int)
		if !ok {
			return false
		}
		switch operator {
		case "=":
			return v == filterNum
		case "!=":
			return v != filterNum
		case ">":
			return v > filterNum
		case "<":
			return v < filterNum
		case ">=":
			return v >= filterNum
		case "<=":
			return v <= filterNum
		}
	case string:
		// Handle cases where integers are stored as strings
		filterStr, ok := filter.(string)
		if !ok {
			return false
		}
		// Try to convert both to integers
		vInt, vErr := strconv.Atoi(v)
		filterInt, fErr := strconv.Atoi(filterStr)

		if vErr == nil && fErr == nil {
			// Both are numbers, compare as numbers
			switch operator {
			case "=":
				return vInt == filterInt
			case "!=":
				return vInt != filterInt
			case ">":
				return vInt > filterInt
			case "<":
				return vInt < filterInt
			case ">=":
				return vInt >= filterInt
			case "<=":
				return vInt <= filterInt
			}
		}

		// If either conversion failed, compare as strings
		switch operator {
		case "=":
			return v == filterStr
		case "!=":
			return v != filterStr
		}
	}
	return false
}
