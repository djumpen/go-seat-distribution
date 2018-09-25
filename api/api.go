package api

import (
	"github.com/djumpen/go-seat-distribution/pkg/generator"
	"github.com/djumpen/go-seat-distribution/pkg/salon"
	"github.com/djumpen/go-seat-distribution/pkg/seat"
	"github.com/djumpen/go-seat-distribution/pkg/storage"

	"encoding/json"
	"net/http"
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

	var req GetNewSalonReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	s, err := api.SalonFactory.NewSalon(req.Rows, req.Blocks, req.SeatCountPerBlock)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	salonID := api.UUIDGenerator.GenerateUUID()
	err = api.Storage.SaveSalon(salonID, *s)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	resp := GetNewSalonResp{
		SalonID: salonID,
		Salon:   *s,
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		newAPIErr(w, err)
		return
	}
}

type GetSalonReq struct {
	SalonID string `json:"SalonID"`
}

type GetSalonResp struct {
	Salon salon.Salon
}

func (api *API) GetSalon(w http.ResponseWriter, r *http.Request) {
	var req GetSalonReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	resp := GetSalonResp{
		Salon: s,
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		newAPIErr(w, err)
		return
	}
}

type AssignSeatReq struct {
	SalonID string `json:"SalonID"`
}

type AssignSeatResp struct {
	AssignedSeat seat.Seat
}

func (api *API) AssignSeat(w http.ResponseWriter, r *http.Request) {
	var req AssignSeatReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	assignedSeat, err := salon.DefaultSeatAssign(&s)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	resp := AssignSeatResp{
		AssignedSeat: assignedSeat,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		newAPIErr(w, err)
		return
	}
}

type SeatInfoReq struct {
	SalonID    string `json:"SalonID"`
	SeatNumber int    `json:"SeatNumber"`
}

type SeatInfoResp struct {
	Seat seat.Seat
}

func (api *API) SeatInfo(w http.ResponseWriter, r *http.Request) {
	var req SeatInfoReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	s, err := api.Storage.GetSalon(req.SalonID)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	seatOnIndex, err := salon.GetSeatByIndex(&s, req.SeatNumber)
	if err != nil {
		newAPIErr(w, err)
		return
	}
	resp := AssignSeatResp{
		AssignedSeat: seatOnIndex,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		newAPIErr(w, err)
		return
	}
}

func newAPIErr(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	w.Write([]byte(err.Error()))
}
