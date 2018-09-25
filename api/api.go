package api

import (
	"github.com/djumpen/go-seat-distribution/pkg/generator"
	"github.com/djumpen/go-seat-distribution/pkg/salon"
	"github.com/djumpen/go-seat-distribution/pkg/seat"
	"github.com/djumpen/go-seat-distribution/pkg/storage"

	"encoding/json"
	"net/http"

	"io/ioutil"
)

type API struct {
	SalonFactory  salon.SalonFactory
	Storage       storage.Storage
	UUIDGenerator generator.Generator
}

func NewAPI(factory salon.SalonFactory, storage storage.Storage, generator generator.Generator) *API {
	return &API{
		SalonFactory:  factory,
		Storage:       storage,
		UUIDGenerator: generator,
	}
}

type GetNewSalonReq struct {
	Rows              int   `json:"Rows"`
	Blocks            int   `json:"Blocks"`
	SeatCountPerBlock []int `json:"SeatCount"`
}

type GetNewSalonResp struct {
	SalonID string
	Salon   salon.Salon
}

func (api *API) GetNewSalon(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	var req GetNewSalonReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	s, err := api.SalonFactory.NewSalon(req.Rows, req.Blocks, req.SeatCountPerBlock)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	salonID := api.UUIDGenerator.GenerateUUID()
	err = api.Storage.SaveSalon(salonID, *s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	resp := GetNewSalonResp{
		SalonID: salonID,
		Salon:   *s,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type GetSalonReq struct {
	SalonID string `json:"SalonID"`
}

type GetSalonResp struct {
	Salon salon.Salon
}

func (api *API) GetSalon(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	var req GetSalonReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	resp := GetSalonResp{
		Salon: s,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type AssignSeatReq struct {
	SalonID string `json:"SalonID"`
}

type AssignSeatResp struct {
	AssignedSeat seat.Seat
}

func (api *API) AssignSeat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	var req AssignSeatReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	assignedSeat, err := salon.DefaultSeatAssign(&s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	resp := AssignSeatResp{
		AssignedSeat: assignedSeat,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type SeatInfoReq struct {
	SalonID    string `json:"SalonID"`
	SeatNumber int    `json:"SeatNumber"`
}

type SeatInfoResp struct {
	Seat seat.Seat
}

func (api *API) SeatInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	var req SeatInfoReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	seatOnIndex, err := salon.GetSeatByIndex(&s, req.SeatNumber)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	resp := AssignSeatResp{
		AssignedSeat: seatOnIndex,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
