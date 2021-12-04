package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)


type filter func(bool) bool


func filterDiagnostics(diagnostics []string, digit int, fn filter) []string {
  filtered := make([]string, 0)

  ones := 0
  for _, value := range diagnostics {
    if value[digit] == '1' { ones += 1 }
  }

  var majority byte = '0'
  if ones >= (len(diagnostics) - ones) {
    majority = '1'
  }

  for _, value := range diagnostics {
    if fn(value[digit] == majority) {
      filtered = append(filtered, value)
    }
  }

  return filtered
}


func keepMajority(majority bool) bool {
  return majority
}


func keepMinority(majority bool) bool {
  return !majority
}


func getCo2Scrubber(diagnostics []string) int64 {
  digits := len(diagnostics[0])

  for i := 0; i < digits; i++ {
    diagnostics = filterDiagnostics(diagnostics, i, keepMinority)

    if len(diagnostics) <= 1 { break }
  }

  val, err := strconv.ParseInt(diagnostics[0], 2, 0)

  if err != nil {
    log.Fatal("Could not convert string to int")
    return -1
  }

  return val
}


func getOxygenGenerator(diagnostics []string) int64 {
  digits := len(diagnostics[0])

  for i := 0; i < digits; i++ {
    diagnostics = filterDiagnostics(diagnostics, i, keepMajority)

    if len(diagnostics) <= 1 { break }
  }

  val, err := strconv.ParseInt(diagnostics[0], 2, 0)

  if err != nil {
    log.Fatal("Could not convert string to int")
    return -1
  }

  return val
}


func runDiagnostics(diagnostics []string) {
  o2 := getOxygenGenerator(diagnostics)
  co2 := getCo2Scrubber(diagnostics)

  fmt.Printf("Life support: %v\n", o2 * co2)
}


func parseAndExecute(file * os.File) {
  scanner := bufio.NewScanner(file)
  diagnostics := make([]string, 0)

  for scanner.Scan() {
    diagnostics = append(diagnostics, scanner.Text())
  }

  runDiagnostics(diagnostics)
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
