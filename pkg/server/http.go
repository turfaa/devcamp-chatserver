package server

import (
	"encoding/json"
	"fmt"

	"chatserver/pkg/usecase"
	"github.com/tokopedia/tdk/go/app/http"
	"github.com/tokopedia/tdk/go/log"
)

type HttpService struct {
}

func NewHttpServer() HttpService {
	return HttpService{}
}

func (s HttpService) RegisterHandler(r *http.Router) {
	r.HandleFunc("/", index, "GET")

	r.HandleFunc("/new_order", handleNewOrder, "POST")

}

func index(ctx http.TdkContext) error {
	ctx.Writer().Write([]byte("Hello world"))
	return nil
}

// we gonna create new order via http API
func handleNewOrder(ctx http.TdkContext) error {
	order := new(usecase.Order)
	err := json.Unmarshal(ctx.Body(), order)
	if err != nil {
		return err
	}

	invoice, err := orderUsecase.PutNewOrder(*order)
	if err != nil {
		log.Error(err)
		return err
	}

	txt := fmt.Sprintf("invoice created: %s", invoice)
	ctx.Write([]byte(txt))
	return nil
}
