package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
  "strings"
)
 

/*

Features:

1. Show number of lines in log [X]
2. Show top 3 messages (errors) that we have and count of appears [X]
3. Show top IPs and count of appears in log []
4. Show top paths and count of appears []

*/

func main() {
  allLines := make(map[int][]string)


  file, err := os.Open("example.log")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  totalLines := 0
  for scanner.Scan() {
   totalLines += 1


   check := strings.Fields(scanner.Text())
   allLines[totalLines] = check

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  // Print all messages
  //for i := 0; i < len(allLines); i++ {
  //  if len(allLines[i]) > 0 {
  //  fmt.Println(allLines[i][8])
  //  }
  //}

  // 1. Show number of lines in log [X]
  fmt.Printf("Total lines: %v\n", totalLines)
  // fmt.Println(allLines)


  

  // 2. Show top 3 messages (errors) that we have and count of appears []
  httpMessages := showErr(allLines)
  fmt.Println(httpMessages)
  //httpMap := make(map[string]int)
  //for i := 0; i < len(httpMessages); i++ {
   // Which is 1 - 100, 2 - 200 etc..
  // mesg := ""
  // switch i {
  // case 0:
  //  mesg = "1xx"
  // case 1:
  //  mesg = "2xx"
  // case 2:
  //  mesg = "3xx"
  // case 3:
  //  mesg = "4xx"
  // case 4:
  //  mesg = "5xx"
  // }
  // httpMap[mesg] = httpMessages[i]
  //}

  fmt.Println("Status codes:")
  mesg := ""
  for i := 0; i < len(httpMessages); i++ {
    switch i {
    case 0:
      continue
     // mesg = "1xx"
    case 1:
      mesg = "2xx"
    case 2:
      continue
     // mesg = "3xx"
    case 3:
      mesg = "4xx"
    case 4:
      mesg = "5xx"
    }
    fmt.Printf("%s --> %v\n", mesg, httpMessages[mesg])
  }





  //top := []int

  // Bubble sort 
  // for i := 0; i < len(httpMessages); i++ {
  //  for j := 0; j < len(httpMessages)-i; j++ {
  //    if httpMessage[j] > httpMessage[j+1] {
  //      httpMessage[j], httpMessage[j+1] = httpMessage[j+1], httpMessage[j]
  //    }
  //  }
  // }

  //for i := range 4 {
  //   fmt.Println("Error
  // }



  // 3. Show top IPs and count of appears in log []
 


  fmt.Println(" ")
  topIPs := topIps(allLines)
  fmt.Println(topIPs)



  // 4. Show top paths and count of appears []
  fmt.Println(" ")
  topPaths := topPaths(allLines)
  fmt.Println(topPaths)
}


func showErr(list map[int][]string) map[string]int {
  oneh, twoh, threeh, fourh, fiveh := 0, 0, 0, 0, 0
  for i := 0; i < len(list); i++ {

    if len(list[i]) > 0 {
    switch list[i][8]{
    case "100", "101", "102", "103":
      oneh += 1
    case "200", "201", "202", "204":
      twoh += 1
    case "301", "302", "303", "307":
      threeh += 1
    case "400", "401", "403", "404", "405", "429":
      fourh += 1
    case "500", "501", "502", "503", "504":
      fiveh += 1
    }
   } 
  }


  return map[string]int{"1xx": oneh, "2xx": twoh, "3xx": threeh, "4xx": fourh, "5xx": fiveh}
}

func topIps(list map[int][]string) map[string]int {
  remember := make(map[string]int)
  for i := 0; i < len(list); i++ {
    if len(list[i]) < 1 {
      continue
    }
    if remember[list[i][0]] != 0 {
        remember[list[i][0]] += 1
    } else {
        remember[list[i][0]] = 1
    }
}

  return remember
}


func topPaths(list map[int][]string) map[string]int {
  remember := make(map[string]int)
  for i := 0; i < len(list); i++ {
    if len(list[i]) < 1 {
      continue
    }
    if remember[list[i][6]] != 0 {
        remember[list[i][6]] += 1
    } else {
        remember[list[i][6]] = 1
    }
}

  return remember
}
