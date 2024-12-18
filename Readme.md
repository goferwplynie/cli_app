# Task app

cli task app written in go. Data is stored in csv file

## instalation

- **Linux**:
  
  clone the repository
  
  ```
  git clone https://github.com/goferwplynie/cli_tasks.git
  cd cli_tasks
  ```
  
  build app
  
  ``` go build -o tasks```
  
  move to the PATH
  
  ```sudo mv tasks /usr/local/bin```

## setup

just run `tasks` command in terminal

## usage

```tasks show``` - show all tasks

```tasks show --id=id``` - show task on specified id

```tasks add task name``` - add task (suports spaces)

```tasks end id``` - end task with specified id


  
  
