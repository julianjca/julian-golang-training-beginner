package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

type paymentCodeServiceHandler struct {
	service golangtraining.IPaymentCodeService
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
		fmt.Print(err)
	}

	var p *golangtraining.PaymentCode
	if err = json.Unmarshal(b, &p); err != nil {
		fmt.Print(err)
	}

	err = s.service.Create(p)
	if err != nil {
		fmt.Print(err)
	}

	e, err := json.Marshal(p)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}

func (s paymentCodeServiceHandler) GetByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusOK)
	res, err := s.service.GetByID(params.ByName("id"))
	jsonData := res

	e, err := json.Marshal(jsonData)

	if err != nil {
		return
	}

	w.Write(e)
}
