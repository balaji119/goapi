package main

import(
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/balaji119/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

	fmt.Println("Starting Go API service...")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
