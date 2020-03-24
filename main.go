package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "bufio"
  "os"
  "os/signal"
  "syscall"
  "strconv"
)

var (
  Info = Teal
  Warn = Yellow
  Fata = Red
)

var (
  Black = color("\033[1;30m%s\033[0m") 
  Red = color("\033[1;31m%s\033[0m") 
  Green = color("\033[1;32m%s\033[0m") 
  Yellow = color("\033[1;33m%s\033[0m") 
  Purple = color("\033[1;34m%s\033[0m") 
  Magenta = color("\033[1;35m%s\033[0m") 
  Teal = color("\033[1;36m%s\033[0m") 
  White = color("\033[1;37m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

func banner() {
  b, err := ioutil.ReadFile("assets/goservelogo.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(string(b))
  fmt.Println("-----------------------------------------------------------------------------------------------------------------")
}

func clean() {
  os.Remove("config/goserve.conf")
  os.Create("config/goserve.conf")
  fmt.Println(Info("GoServe cleaned, run make to start again"))
}

func setupCloseHandler() {
  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() { 
	<-c
    fmt.Println(Fata("\n\nERROR: Ctrl+C detected, exiting."))
    os.Exit(0)
  }()
}

func setup() {

  reader := bufio.NewReader(os.Stdin)
  banner()

  steps := [] string {Info("Port"), Info("Source (starting from /)")}

    for i := 0;i < len(steps);i++{
    fmt.Print(steps[i] + " $ ")
    if i != 0 {
      fmt.Print("/")
    }
    text, _ := reader.ReadString('\n')
      // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    _, err := strconv.Atoi(text)

    // At this point, we consider all is good
    settings, err := os.OpenFile("config/goserve.conf", os.O_APPEND | os.O_WRONLY, 0644)

    if err != nil {
      fmt.Println(err)
      return
    }

    write, err := settings.WriteString(text + ";")

    if err != nil {
      fmt.Println(err)
      settings.Close()
      fmt.Println(write)
      return
    }

    err = settings.Close()

    if err != nil {
      fmt.Println(err)
      return
    }

  }

  fmt.Println("All done setting up!")
  serve()

}

func serve() {
  banner()

  data, err := ioutil.ReadFile("config/goserve.conf")
  if err != nil {
    fmt.Println(err)
    return
  }
  parsed := strings.Split(string(data), ";")
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fmt.Println(parsed)

  port := parsed[0]
  srcDir := "/" + parsed[1]

  http.Handle("/", http.FileServer(http.Dir(srcDir)))
  fmt.Println("Server is up and running, serving " + srcDir + ".")
  if err := http.ListenAndServe(":" + port, nil);
  err != nil {
    panic(err)
  }
}

func main() {
  setupCloseHandler()

  confLen, _ := os.Stat("config/goserve.conf")
  conf := confLen.Size()
  nilFile, _ := os.Stat("config/nilFile.goserve")
  nilLen := nilFile.Size()

  if conf == nilLen {
    setup()
  } else {
    fmt.Print(Info("Would you like to restart? "))
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    if text == "y" {
      fmt.Print("Are you sure? ")
      reader2 := bufio.NewReader(os.Stdin)
      text2, _ := reader2.ReadString('\n')
      text2 = strings.Replace(text2, "\n", "", -1)
      if text2 == "y" {
        clean()
      }
    } else if text == "n" {
      serve()
    }
  }
}