package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
			tm := time.Now().Format(time.RFC1123)
			_, _ = w.Write([]byte("The time is " + tm))
	}

	return http.HandlerFunc(fn)
}

func main() {
	gin.New()
	mux := http.NewServeMux()

	th := timeHandler(time.RFC1123)
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
