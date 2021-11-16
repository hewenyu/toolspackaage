package httpc

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ForwardHandler(writer http.ResponseWriter, request *http.Request) {
	u, err := url.Parse("https://test.com/path/to/uri")
	if nil != err {
		log.Println(err)
		return
	}

	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
		},
	}

	proxy.ServeHTTP(writer, request)
}

func ForwardHandler2(writer http.ResponseWriter, request *http.Request) {
	u := &url.URL{
		Scheme: "https",
		Host:   "target.com",
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	request.URL.Path = "/path/to/uri"
	proxy.ServeHTTP(writer, request)
}
