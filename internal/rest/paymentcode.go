package rest

import (
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

type virtualAccountsServiceHandler struct {
	service golangtraining.IPaymentCodeService
}

// InitVirtualAccountsRESTHandler will initialize the REST handler for Virtual Account Settings
func InitPaymentCodeRESTHandler(r *httprouter.Router, service golangtraining.IPaymentCodeService) {
	r.POST("/payment-codes", service.Create)
	r.GET("/payment-codes/:id", service.GetByID)
}
