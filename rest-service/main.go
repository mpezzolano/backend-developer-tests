package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/common"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/people"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/router"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/storage"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	logger := common.NewLogger()
	httpAddress := flag.String("http", ": 3000", "http listen address")

	// define repository, service and endpoints

	peopleRepo := people.NewRepository(storage.NewStorage())
	peopleService := people.NewService(peopleRepo)
	peopleEndpoints := people.MakeEndpoints(peopleService, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	mwfs := make([]mux.MiddlewareFunc, 0)

	go func() {
		_ = logger.Log("Listening On Port", *httpAddress)
		errs <- http.ListenAndServe(*httpAddress, router.NewHandler(router.Endpoints{People: peopleEndpoints}, mwfs...))
	}()

	fmt.Println(<-errs)

}
