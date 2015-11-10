package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ninnemana/boomer/api"
	"golang.org/x/net/http2"
)

const (
	HEADER = `
    _------_
   /        \
  |          |
  |          |
  |     __  __)
  |    /  \/  \
 /\/\ (o   )o  )
 /c    \__/ --.    ______________
 \_   _-------'   /              \\
 |  /         \ <   HEY BOOMER!!! \\
 | | '\_______)   \_______________\\
 |  \_____)
 |_____ |
|_____/\/\
/        \
	`
)

func main() {
	router := httprouter.New()

	router.GET("/", api.Index)
	router.POST("/bench", api.Bench)

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	http2.ConfigureServer(&srv, &http2.Server{})

	log.Printf("%s\n", HEADER)
	log.Printf("Starting server on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
