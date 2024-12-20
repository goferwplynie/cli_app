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
	if len(args) > 1 {
		switch args[1] {
		case "add":
			add_task(args[2:])
		case "show":
			show_tasks(args[1:])
		case "end":
			end_task(args[2])
		}
	} else {
		setup()
	}
}

func setup() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		return
	}

	os.Mkdir(homeDir+"/.tasks", 0755)

	fmt.Println("created tasks directory")

	os.Create(homeDir + "/.tasks/tasks.csv")

	fmt.Println("added tasks.csv")

	file, err := os.OpenFile(homeDir+"/.tasks/tasks.csv", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening a file ", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll([][]string{{"id", "name"}})
	if err != nil {
		fmt.Println("Error writing to file ", err)
		return
	}
	fmt.Println("finished setup. app ready to use!")
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
		fmt.Println("Error while converting to int", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening a file", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading a file", err)
		return
	}

	records = records[1:]
	updatedRecords := records

	for i, row := range records {
		current_id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println("Error while converting to int", err)
			return
		}
		if current_id == rowToDelete {
			updatedRecords = [][]string{{"id", "name"}}
			updatedRecords = append(updatedRecords, records[:i]...)
			updatedRecords = append(updatedRecords, records[i+1:]...)

		}
	}

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		fmt.Println("Error writing to file ", err)
		return
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
