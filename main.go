package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func main() {
	args := os.Args
	switch args[1] {
	case "add":
		add_task(args[1:])
	case "show":
		show_tasks(args[1:])
	case "end":
		end_task(args[1:])
	}
}

func end_task(args []string) {
	flags := flag.NewFlagSet("end", flag.ExitOnError)

	rowToDelete := flags.Int("id", 0, "search task by id")

	err := flags.Parse(args[1:])

	if err != nil {
		panic(err)
	}

	// Open the file for reading
	file, err := os.Open("tasks.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Check if the rowToDelete is valid
	if *rowToDelete < 0 || *rowToDelete >= len(records) {
		panic("wrong id")
	}

	records = records[1:]

	// Remove the specific row
	updatedRecords := append(records[:*rowToDelete], records[*rowToDelete+1:]...)

	// Reopen the file for writing (truncate and overwrite)
	file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the updated records back to the file
	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		panic(err)
	}
}

func show_tasks(args []string) {
	flags := flag.NewFlagSet("show", flag.ExitOnError)

	id := flags.Int("id", 0, "search task by id")

	err := flags.Parse(args[1:])

	if err != nil {
		panic(err)
	}

	file, err := os.Open("tasks.csv")

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
	flags := flag.NewFlagSet("add", flag.ExitOnError)

	name := flags.String("name", "example task", "add new task")

	err := flags.Parse(args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	file, err := os.OpenFile("tasks.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
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
	data := []string{strconv.Itoa(new_id), *name}

	fmt.Println("Adding task:", data)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(data); err != nil {
		panic(err)
	}

	fmt.Println("Task added successfully!")
}
