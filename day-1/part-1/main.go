package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)


func printRate(rate int) {
  fmt.Printf("Rate: %v\n", rate)
  return
}


func findDepthRate(depths []int) {
  if len(depths) < 1 {
    printRate(0)
    return
  }

  rate := 0
  previous := depths[0]

  for i := 1; i < len(depths); i++ {
    if depths[i] > previous {
      rate++
    }

    previous = depths[i]
  }

  printRate(rate)
  return
}


func parseAndExecute(file * os.File) {
  scanner := bufio.NewScanner(file)
  depths := make([]int, 0)

  for scanner.Scan() {
    value, err := strconv.Atoi(scanner.Text())

    if err != nil {
      log.Fatal(err)
    }

    depths = append(depths, value)
  }

  findDepthRate(depths)
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
