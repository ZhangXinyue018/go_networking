package reverse

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// todo: why always direct to 9090 not 8080 ????
func Start() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:8080"
		req.URL.Path = "/auth"
	}
	proxy := &httputil.ReverseProxy{Director: director}
	fmt.Println("Serve on :9090")
	log.Fatal(http.ListenAndServe(":9090", proxy))
}
