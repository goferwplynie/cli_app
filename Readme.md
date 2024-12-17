# Task app

cli task app written in go. Data is stored in csv file

## usage

`go run main.go add --name=example_name` - add new task

`go run main.go show` - show all tasks in form of table

`go run main.go show --id=1` - search task by id

`go run main.go end id` - mark task as completed and delete csv record
