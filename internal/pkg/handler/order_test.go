package handler

import (
	"OrderService/internal/models"
	"OrderService/internal/pkg/service"
	mockService "OrderService/internal/pkg/service/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"log"
	"net/http/httptest"
	"testing"
)

const orderJsonOutputById = `{"order_uid":"b563feb7b2b84b6test","track_number":"WBILMTESTTRACK","entry":"WBIL","delivery":{"name":"TestTestov","phone":"+9720000000","zip":"2639809","city":"KiryatMozkin","address":"PloshadMira15","region":"Kraiot","email":"test@gmail.com"},"payment":{"transaction":"b563feb7b2b84b6test","request_id":"","currency":"USD","provider":"wbpay","amount":1817,"payment_dt":1637907727,"bank":"alpha","delivery_cost":1500,"goods_total":317,"custom_fee":0},"items":[{"chrt_id":9934930,"track_number":"WBILMTESTTRACK","price":453,"rid":"ab4219087a764ae0btest","name":"Mascaras","sale":30,"size":"0","total_price":317,"nm_id":2389212,"brand":"VivienneSabo","status":202}],"locale":"en","internal_signature":"","customer_id":"test","delivery_service":"meest","shardkey":"9","sm_id":99,"date_created":"2021-11-26T06:22:19Z","oof_shard":"1"}`
const allOrdersJsonOutput = `[{"order_uid":"b563feb7b2b84b6test","track_number":"WBILMTESTTRACK","entry":"WBIL","delivery":{"name":"TestTestov","phone":"+9720000000","zip":"2639809","city":"KiryatMozkin","address":"PloshadMira15","region":"Kraiot","email":"test@gmail.com"},"payment":{"transaction":"b563feb7b2b84b6test","request_id":"","currency":"USD","provider":"wbpay","amount":1817,"payment_dt":1637907727,"bank":"alpha","delivery_cost":1500,"goods_total":317,"custom_fee":0},"items":[{"chrt_id":9934930,"track_number":"WBILMTESTTRACK","price":453,"rid":"ab4219087a764ae0btest","name":"Mascaras","sale":30,"size":"0","total_price":317,"nm_id":2389212,"brand":"VivienneSabo","status":202}],"locale":"en","internal_signature":"","customer_id":"test","delivery_service":"meest","shardkey":"9","sm_id":99,"date_created":"2021-11-26T06:22:19Z","oof_shard":"1"}]`

func TestOrder_GetAll(t *testing.T) {
	type mock func(s *mockService.MockOrders)

	orders := make([]models.Order, 1, 1)

	err := json.Unmarshal([]byte(orderJsonOutputById), &orders[0])
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name                 string
		mock                 mock
		expectedStatusCode   int
		expectedResponseData string
	}{
		{
			name: "Ok",
			mock: func(s *mockService.MockOrders) {
				s.EXPECT().GetAll().Return(orders, nil)
			},
			expectedStatusCode:   200,
			expectedResponseData: allOrdersJsonOutput,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockOrders(c)
			testCase.mock(repo)

			services := &service.Service{Orders: repo}
			handler := Handler{service: services}

			r := gin.New()
			r.GET("/api/orders/all", handler.getAll)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/orders/all", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseData, w.Body.String())
		})
	}
}

func TestOrder_GetOrderByUIDFromCache(t *testing.T) {
	type mock func(s *mockService.MockOrders, uid string)

	var order models.Order
	err := json.Unmarshal([]byte(orderJsonOutputById), &order)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name                 string
		uid                  string
		mock                 mock
		expectedStatusCode   int
		expectedResponseData string
	}{
		{
			name: "Ok",
			uid:  "b563feb7b2b84b6test",
			mock: func(s *mockService.MockOrders, uid string) {
				s.EXPECT().GetByUID(uid).Return(order, nil)
			},
			expectedStatusCode:   200,
			expectedResponseData: orderJsonOutputById,
		},
		{
			name: "Uid not found",
			uid:  "666testtesttesttest",
			mock: func(s *mockService.MockOrders, uid string) {
				s.EXPECT().GetByUID(uid).Return(models.Order{}, errors.New("no rows in result set"))
			},
			expectedStatusCode:   500,
			expectedResponseData: `{"message":"no rows in result set"}`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockOrders(c)
			testCase.mock(repo, testCase.uid)

			services := &service.Service{Orders: repo}
			handler := Handler{service: services}

			r := gin.New()
			r.GET("/api/orders/:uid", handler.GetOrderByUIDFromCache)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/orders/%s", testCase.uid), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseData, w.Body.String())
		})
	}
}
