package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "bufio"
  "os"
  "strconv"
)

func ignore(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(""))
}

func banner() {
    b, err := ioutil.ReadFile("assets/goservelogo.txt")
    if err != nil {
        panic(err)
    }
	fmt.Println(string(b))
	fmt.Println("-----------------------------------------------------------------------------------------------------------------")
}

func clean(){
	os.Remove("conifg/goserve.conf")
	os.Create("config/goserve.conf")
	fmt.Println("GoServe cleaned, run make to start again")
}

func setup() {

  reader := bufio.NewReader(os.Stdin)
  banner()

  steps := []string{"Port", "Source (starting from /)", "Ignore Dir (starting from /)"}

  for i:=0; i<len(steps);i++{
	fmt.Print(steps[i] + " $ ")
	if i != 0 {
		fmt.Print("/")
	}
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	num, err := strconv.Atoi(text)
	if err == nil && num <= 1024 && i == 0 {
		fmt.Println("We don't support low ports yet")
		os.Exit(1)
	}

	// At this point, we consider all is good
	settings, err := os.OpenFile("config/goserve.conf", os.O_APPEND|os.O_WRONLY, 0644)
	
	if err != nil {
        fmt.Println(err)
        return
    }
	
	write, err := settings.WriteString(text + ";")
	fmt.Println(write)

	if err != nil {
        fmt.Println(err)
        settings.Close()
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
  parsed := strings.Split(string(data), ";")

  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fmt.Println(parsed)

  port := parsed[0]
  srcDir := "/" + parsed[1]
  ignoreDir := "/" + parsed[2]

  http.Handle("/", http.FileServer(http.Dir(srcDir)))
  http.HandleFunc(ignoreDir, ignore)
  fmt.Println("Server is up and running, serving " + srcDir + ", ignoring " + ignoreDir + ".")
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}

func main(){
  confLen, _ := os.Stat("config/goserve.conf")
  conf := confLen.Size()
  nilFile, _ := os.Stat("config/nilFile.goserve")
  nilLen := nilFile.Size()

  if conf == nilLen {
    setup()
  } else {
	fmt.Print("Would you like to restart? ")
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