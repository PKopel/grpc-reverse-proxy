package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	defaultPort = ":50051"
)

type reverseProxy struct {
	proxies []*httputil.ReverseProxy
	current int
}

func (proxy *reverseProxy) handle(w http.ResponseWriter, r *http.Request) {
	p := proxy.proxies[proxy.current]
	proxy.current = (proxy.current + 1) % len(proxy.proxies)

	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	p.ServeHTTP(w, r)
}

func main() {
	port := defaultPort
	servers := []string{}
	switch argsLen := len(os.Args); {
	case argsLen == 2:
		port = os.Args[1]
	case argsLen > 2:
		port = os.Args[1]
		servers = os.Args[2:]
	}
	log.Printf("port: %v forwarded to: %v", port, servers)

	proxies := []*httputil.ReverseProxy{}
	for _, server := range servers {
		remote, err := url.Parse(server)
		if err != nil {
			panic(err)
		}

		proxies = append(proxies, httputil.NewSingleHostReverseProxy(remote))
	}

	proxy := reverseProxy{proxies: proxies, current: 0}
	http.HandleFunc("/", proxy.handle)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
