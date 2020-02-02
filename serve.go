import (
    "flag"
    "github.com/gorilla/mux"
    "io"
    "log"
    "net"
    "net/http"
    "net/http/fcgi"
    "runtime"
)

var local = flag.String("local", "", "serve as webserver, example: 0.0.0.0:8000")

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func homeView(w http.ResponseWriter, r *http.Request) {
    headers := w.Header()
    headers.Add("Content-Type", "text/html")
    io.WriteString(w, "<html><head></head><body><p>It works!</p></body></html>")
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", homeView)

    flag.Parse()
    var err error

    if *local != "" { // Run as a local web server
        err = http.ListenAndServe(*local, r)
    } else { // Run as FCGI via standard I/O
        err = fcgi.Serve(nil, r)
    }
    if err != nil {
        log.Fatal(err)
    }
}
