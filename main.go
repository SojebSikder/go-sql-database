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

	if len(os.Args) < 2 {

	} else {
		arg := os.Args[1]
		if arg == "cli" {
			cli()
		} else if arg == "run" {
			if len(os.Args) < 3 {
				fmt.Println("Please provide a file name")
			} else {
				fileName := os.Args[2]
				executeFromFile(fileName)
			}
		} else {
			fmt.Println("Invalid argument")
		}
	}
}

func cli() {
	engine := storage.NewStorageEngine()
	executor := execution.NewExecutor(engine)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("SQL Database Engine (Type 'exit' to quit)")

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

func executeFromFile(fileName string) {
	engine := storage.NewStorageEngine()
	executor := execution.NewExecutor(engine)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		queryStr := scanner.Text()

		queries := parser.SplitQueries(queryStr)
		for _, queryStr := range queries {

			query, err := parser.Parse(queryStr)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			executor.Execute(query)
		}
	}
}
