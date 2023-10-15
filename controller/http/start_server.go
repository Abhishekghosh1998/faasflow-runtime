package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
)

// StartServer starts the flow function
func StartServer(runtime runtime.Runtime, port int, readTimeout time.Duration, writeTimeout time.Duration) error {

	fmt.Println("runtime/controller/http/start_server::StartServer start")
	err := runtime.Init()
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		Handler:        router(runtime),
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	fmt.Println("runtime/controller/http/start_server::StartServer end")
	return s.ListenAndServe()
}
