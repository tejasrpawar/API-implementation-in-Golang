package main
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"go_project/internal/handlers"
)


func main(){
	log.SetReportCaller(true)
	var r = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service..")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}

}

