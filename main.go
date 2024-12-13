package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	args := os.Args
	switch args[1] {
	case "add":
		add_task(args[2:])
	case "show":
		show_tasks(args[1:])
	case "end":
		end_task(args[2])
	}
}

func end_task(id string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		return
	}

	// Define the global file path
	filePath := filepath.Join(homeDir, ".tasks", "tasks.csv")

	rowToDelete, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	records = records[1:]
	updatedRecords := records

	for i, row := range records {
		current_id, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}
		if current_id == rowToDelete {
			updatedRecords = [][]string{{"id", "name"}}
			updatedRecords = append(updatedRecords, records[:i]...)
			updatedRecords = append(updatedRecords, records[i+1:]...)

		}
	}

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		panic(err)
	}

}

func show_tasks(args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		return
	}

	// Define the global file path
	filePath := filepath.Join(homeDir, ".tasks", "tasks.csv")

	flags := flag.NewFlagSet("show", flag.ExitOnError)

	id := flags.Int("id", 0, "search task by id")

	err = flags.Parse(args[1:])

	if err != nil {
		panic(err)
	}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	rows = rows[1:]
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"id", "name"})

	if *id != 0 {
		for _, row := range rows {
			current_id, err := strconv.Atoi(row[0])
			if err != nil {
				panic(err)
			}
			if current_id == *id {
				table.Append(row)
				table.Render()
			}
		}
	} else {

		for _, row := range rows {
			table.Append(row)
		}

		table.Render()

	}

}

func add_task(args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		return
	}

	name := strings.Join(args, " ")

	// Define the global file path
	filePath := filepath.Join(homeDir, ".tasks", "tasks.csv")

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	last_id := 0
	last_id, err = strconv.Atoi(rows[len(rows)-1][0])

	new_id := last_id + 1
	data := []string{strconv.Itoa(new_id), name}

	fmt.Println("Adding task:", data)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(data); err != nil {
		panic(err)
	}

	fmt.Println("Task added successfully!")
}
