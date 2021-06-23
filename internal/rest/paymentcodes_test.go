package rest_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julianjca/julian-golang-training-beginner/internal/rest"
	"github.com/julianjca/julian-golang-training-beginner/internal/rest/mocks"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockService
		starwars       *mocks.MockStarWarsClient
		expectedReturn error
		url            string
		body           io.Reader
		expectedCode   int
	}{
		{
			desc: "create payment codes - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)

				return m
			}(),
			starwars: func() *mocks.MockStarWarsClient {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockStarWarsClient(ctrl)

				return m
			}(),
			expectedReturn: nil,
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - failed because name not provided",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				return m
			}(),
			starwars: func() *mocks.MockStarWarsClient {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockStarWarsClient(ctrl)

				return m
			}(),
			expectedReturn: nil,
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": ""
				}
			`),
			expectedCode: http.StatusBadRequest,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - failed from service",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(errors.New("internal server error"))

				return m
			}(),
			starwars: func() *mocks.MockStarWarsClient {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockStarWarsClient(ctrl)

				return m
			}(),
			expectedReturn: nil,
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusInternalServerError,
			url:          "/payment-codes",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodeRESTHandler(r, tC.service, tC.starwars)

			req := httptest.NewRequest("POST", tC.url, tC.body)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			require.Equal(t, tC.expectedCode, rec.Code)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockService
		starwars       *mocks.MockStarWarsClient
		expectedReturn golangtraining.PaymentCode
		url            string
		expectedCode   int
	}{
		{
			desc: "get payment codes by ID - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				p := golangtraining.PaymentCode{
					Name:        "test",
					PaymentCode: "test-1",
				}

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(p, nil)

				return m
			}(),
			starwars: func() *mocks.MockStarWarsClient {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockStarWarsClient(ctrl)

				return m
			}(),
			expectedReturn: golangtraining.PaymentCode{
				Name:        "test",
				PaymentCode: "test-1",
			},
			expectedCode: http.StatusOK,
			url:          "/payment-codes/id-123",
		},
		{
			desc: "get payment codes by ID - failed",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("error from server"))

				return m
			}(),
			starwars: func() *mocks.MockStarWarsClient {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockStarWarsClient(ctrl)

				return m
			}(),
			expectedReturn: golangtraining.PaymentCode{
				Name:        "test",
				PaymentCode: "test-1",
			},
			expectedCode: http.StatusNotFound,
			url:          "/payment-codes/id-123",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodeRESTHandler(r, tC.service, tC.starwars)

			req := httptest.NewRequest("GET", tC.url, nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			require.Equal(t, tC.expectedCode, rec.Code)
		})
	}
}
