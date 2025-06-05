package handlers

import (
	"github.com/Pallinder/go-randomdata"
	"light-ms/order/internal/testutils"
	"light-ms/order/internal/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrdersHandler_CreateOrder(t *testing.T) {
	mockHandler := OrdersHandler{
		ucase: usecase.NewOrdersUcase(testutils.MockOrdersRepository{}),
	}

	router := testutils.GetRouter()
	router.POST("/orders", mockHandler.CreateOrder)

	payload := map[string]any{
		"user_id": randomdata.Number(10),
		"amount":  randomdata.Number(100),
	}

	r := httptest.NewRequest("POST", "/orders", testutils.ToReader(payload))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Error("Incorrect code", testutils.RespBody(resp))
	}
}
