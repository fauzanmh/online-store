package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/fauzanmh/online-store/schema/request"
	"github.com/fauzanmh/online-store/schema/response"
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const (
	baseUrl = "http://localhost:8099/api/v1"
)

func parseRequestBody(data interface{}) io.Reader {
	body, _ := json.Marshal(data)
	return bytes.NewBuffer(body)
}
func TestCheckoutOrder(t *testing.T) {
	type testCase struct {
		name    string
		request request.CheckoutRequest
	}

	type responseBody struct {
		response.Base
	}

	// case
	testCases := []testCase{
		{
			name: "Checkout Order by Customer 1",
			request: request.CheckoutRequest{
				Items: []request.ItemsRequest{
					{
						ProductID: 5,
						Qty:       1,
					},
					{
						ProductID: 2,
						Qty:       5,
					},
				},
			},
		},
		{
			name: "Checkout Order by Customer 2",
			request: request.CheckoutRequest{
				Items: []request.ItemsRequest{
					{
						ProductID: 5,
						Qty:       1,
					},
					{
						ProductID: 4,
						Qty:       10,
					},
				},
			},
		},
		{
			name: "Checkout Order by Customer 3",
			request: request.CheckoutRequest{
				Items: []request.ItemsRequest{
					{
						ProductID: 5,
						Qty:       1,
					},
					{
						ProductID: 1,
						Qty:       7,
					},
				},
			},
		},
		{
			name: "Checkout Order by Customer 4",
			request: request.CheckoutRequest{
				Items: []request.ItemsRequest{
					{
						ProductID: 5,
						Qty:       1,
					},
					{
						ProductID: 3,
						Qty:       15,
					},
				},
			},
		},
	}

	// init go routines
	var wg sync.WaitGroup
	result := make(chan bool, len(testCases))

	for idx := range testCases {
		wg.Add(1)
		go func(idx int) {
			defer func() {
				if x := recover(); x != nil {
					errPanic := fmt.Errorf("name: %s ,run time panic: %v", testCases[idx].name, x)
					zap.S().Error(errPanic)
				}
				wg.Done()
			}()

			t.Run(fmt.Sprintf(testCases[idx].name, idx), func(t *testing.T) {
				httpClient := httpclient.NewClient()

				headers := http.Header{}
				headers.Set("Content-Type", "application/json")
				request := testCases[idx].request

				url := baseUrl + "/orders/checkout"
				res, err := httpClient.Post(url, parseRequestBody(request), headers)
				if err != nil {
					assert.NoError(t, err)
					return
				}

				defer res.Body.Close()
				respBody, err := ioutil.ReadAll(res.Body)
				if err != nil {
					assert.NoError(t, err)
					return
				}

				resp := responseBody{}

				err = json.Unmarshal([]byte(respBody), &resp)
				if err != nil {
					assert.NoError(t, err)
					return
				}

				if res.StatusCode != 200 {
					result <- true
				}
			})
		}(idx)
	}

	wg.Wait()

	successTest := 0
	for range testCases {
		select {
		case <-result:
			successTest += 1
			continue
		default:
			continue
		}
	}

	if successTest > 0 {
		assert.Equal(t, "PASS", "PASS")
	} else {
		assert.Equal(t, "PASS", "FAIL")
	}
}
