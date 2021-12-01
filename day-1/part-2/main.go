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


func sumSlice(slice []int) int {
  sum := 0

  for i := 0; i < len(slice); i++ {
    sum += slice[i]
  }

  return sum
}


func findDepthRate(depths []int) {
  if len(depths) < 4 {
    printRate(0)
    return
  }

  rate := 0
  previous := sumSlice(depths[0 : 3])

  for i := 4; i <= len(depths); i++ {
    current := sumSlice(depths[i-3 : i])

    if current > previous {
      rate++
    }

    previous = current
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
