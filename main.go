package main

import (
	"bufio"
	"fmt"
	"os"
	"sojebsql/execution"
	"sojebsql/parser"
	"sojebsql/storage"
)

func main() {
	engine := storage.NewStorageEngine()
	executor := execution.NewExecutor(engine)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple SQL Database Engine (Type 'exit' to quit)")

	for {
		fmt.Print("SQL> ")
		scanner.Scan()
		queryStr := scanner.Text()

		if queryStr == "exit" {
			break
		}

		query, err := parser.Parse(queryStr)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		executor.Execute(query)
	}
}
