package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)


func runDiagnostics(diagnostics []string) {
  ones := make([]int, len(diagnostics[0]))

  for i := 0; i < len(diagnostics); i++ {
    for j := 0; j < len(diagnostics[i]); j++ {
      if diagnostics[i][j] == '1' {
        ones[j] += 1
      }
    }
  }

  gamma := 0
  epsilon := 0
  majority := len(diagnostics) / 2
  digits := len(ones)

  for i := 0; i < digits; i++ {
    value := 1 << (digits - 1 - i)
    if ones[i] >= majority {
      gamma += value
    } else {
      epsilon += value
    }
  }

  fmt.Printf("Power consumption: %v\n", gamma * epsilon)
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
