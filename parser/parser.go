package parser

import (
	"fmt"
	"strings"
)

// Query structure
type Query struct {
	Type    string
	Table   string
	Values  []string
	Columns []string
}

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
		return &Query{
			Type:  "SELECT",
			Table: words[2],
		}, nil

	default:
		return nil, fmt.Errorf("unsupported query type")
	}
}
