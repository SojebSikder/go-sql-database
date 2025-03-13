package parser

import (
	"fmt"
	"strings"
)

// Query structure
type Query struct {
	Type     string
	Table    string
	Columns  []string
	Values   []string
	WhereCol string
	WhereOp  string
	WhereVal string
}

// Supported operators
var operators = []string{"=", "!=", ">", "<", ">=", "<="}

// Parse processes SQL-like queries
func Parse(query string) (*Query, error) {
	words := strings.Fields(strings.ToUpper(query))
	if len(words) == 0 {
		return nil, fmt.Errorf("empty query")
	}

	switch words[0] {
	case "CREATE":
		if len(words) < 4 || words[1] != "TABLE" {
			return nil, fmt.Errorf("invalid CREATE TABLE syntax")
		}
		return &Query{
			Type:    "CREATE",
			Table:   words[2],
			Columns: words[3:],
		}, nil

	case "INSERT":
		if len(words) < 4 || words[1] != "INTO" || words[3] != "VALUES" {
			return nil, fmt.Errorf("invalid INSERT syntax")
		}
		return &Query{
			Type:   "INSERT",
			Table:  words[2],
			Values: words[4:],
		}, nil

	case "SELECT":
		if len(words) < 2 || words[1] != "FROM" {
			return nil, fmt.Errorf("invalid SELECT syntax")
		}
		queryStruct := &Query{
			Type:  "SELECT",
			Table: words[2],
		}

		// Handle WHERE clause
		for i := 3; i < len(words)-2; i++ {
			if words[i] == "WHERE" {
				for _, op := range operators {
					if i+2 < len(words) && words[i+2] == op {
						queryStruct.WhereCol = words[i+1]
						queryStruct.WhereOp = words[i+2]
						queryStruct.WhereVal = words[i+3]
						return queryStruct, nil
					}
				}
				return nil, fmt.Errorf("invalid WHERE clause syntax")
			}
		}

		return queryStruct, nil

	default:
		return nil, fmt.Errorf("unsupported query type")
	}
}

// SplitQueries breaks a multi-query string into individual queries
func SplitQueries(queryStr string) []string {
	queries := strings.Split(queryStr, ";")
	var result []string
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query != "" {
			result = append(result, query)
		}
	}
	return result
}
