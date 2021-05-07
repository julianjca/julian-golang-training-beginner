package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

type paymentCodeServiceHandler struct {
	service golangtraining.IPaymentCodeService
}

type GetByIDRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InitVirtualAccountsRESTHandler will initialize the REST handler for Virtual Account Settings
func InitPaymentCodeRESTHandler(r *httprouter.Router, service golangtraining.IPaymentCodeService) {
	h := paymentCodeServiceHandler{
		service: service,
	}

	r.POST("/payment-codes", h.Create)
	r.GET("/payment-codes/:id", h.GetByID)
}

func (s paymentCodeServiceHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
	}

	var p *golangtraining.PaymentCode
	if err = json.Unmarshal(b, &p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
	}

	if p.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"bad request"}`))
		return
	}

	err = s.service.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error creating"}`))
	}

	e, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}

func (s paymentCodeServiceHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	pID := ps.ByName("id")

	res, err := s.service.GetByID(pID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"not found"}`))
		return
	}

	e, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(e)
}
