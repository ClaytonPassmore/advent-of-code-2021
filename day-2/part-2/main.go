package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)


type Command struct {
  instruction string
  value int
}


type Position struct {
  aim int
  depth int
  horizontal int
}


func updatePosition(position * Position, command * Command) {
  switch command.instruction {
  case "forward":
    position.horizontal += command.value
    position.depth += command.value * position.aim
  case "up":
    position.aim -= command.value
  case "down":
    position.aim += command.value
  }

  return
}


func executeCommands(commands []Command) {
  position := Position {
    aim: 0,
    depth: 0,
    horizontal: 0,
  }

  for i := 0; i < len(commands); i++ {
    updatePosition(&position, &commands[i])
  }

  fmt.Printf("Depth: %v\n", position.depth)
  fmt.Printf("Horizontal: %v\n", position.horizontal)
  fmt.Printf("Product: %v\n", position.horizontal * position.depth)

  return
}


func parseAndExecute(file * os.File) {
  scanner := bufio.NewScanner(file)
  commands := make([]Command, 0)

  for scanner.Scan() {
    line := scanner.Text()
    line_tokens := strings.Split(line, " ")
    value, err := strconv.Atoi(line_tokens[1])

    if err != nil {
      log.Fatal(err)
    }

    command := Command {
      instruction: line_tokens[0],
      value: value,
    }

    commands = append(commands, command)
  }

  executeCommands(commands)
  return
}


func main() {
  file, err := os.Open(os.Args[1])

  if err != nil {
    log.Fatal(err)
    return
  }

  defer file.Close()

  parseAndExecute(file)
  return
}
