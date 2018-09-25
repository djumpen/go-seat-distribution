package main

import (
	"log"
	"net/http"

	"github.com/djumpen/go-seat-distribution/api"
	"github.com/djumpen/go-seat-distribution/cfg"
	"github.com/djumpen/go-seat-distribution/pkg/generator"
	"github.com/djumpen/go-seat-distribution/pkg/salon"
	"github.com/djumpen/go-seat-distribution/pkg/storage"

	"github.com/bluele/gcache"
)

func main() {

	salonCache := gcache.New(cfg.CacheSize).Build()

	storage := storage.NewCacheStorage(salonCache)
	salonFactory := salon.NewDefaultSalonFactory(salon.DefaultSeatClassifier, salon.DefaultSeatNaming)
	generator := generator.NewUUIDGenerator()

	api := api.NewAPI(salonFactory, storage, generator)

	http.HandleFunc("/salon/new", api.GetNewSalon)
	http.HandleFunc("/salon/get", api.GetSalon)
	http.HandleFunc("/seat/assign", api.AssignSeat)
	http.HandleFunc("/seat/info", api.SeatInfo)

	log.Fatal(http.ListenAndServe(cfg.HTTPListenPort, nil))
}
