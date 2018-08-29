package server

import (
	"encoding/json"
	"fmt"

	"chatserver/pkg/domain"
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

	r.HandleFunc("/get_messages/{username}", handleGetMessage, "GET")
	r.HandleFunc("/post_messages", handlePostMessage, "POST")
	r.HandleFunc("/rooms", handleGetRooms, "GET")
	r.HandleFunc("/current_rooms", handleGetCurrentRooms, "GET")
	r.HandleFunc("/rooms/join", handleJoinRooms, "POST")
}

func handleGetMessage(ctx http.TdkContext) error {
	username := ctx.Vars()["username"]
	if messages, err := chatUsecase.GetMessages(username); err != nil {
		log.Error(err.Error())
		ctx.JSON(err.Error())
	} else {
		ctx.JSON(messages)
	}
	return nil
}

func handlePostMessage(ctx http.TdkContext) error {
	message := new(domain.Message)
	if err := json.Unmarshal(ctx.Body(), message); err != nil {
		return err
	}

	if err := chatUsecase.SendMessage(message); err != nil {
		return err
	}

	ctx.JSON(message)

	return nil
}

func handleGetRooms(ctx http.TdkContext) error {
	if rooms, err := chatUsecase.GetAllRooms(); err != nil {
		return err
	} else {
		ctx.JSON(rooms)
	}
	return nil
}

func handleGetCurrentRooms(ctx http.TdkContext) error {
	return nil
}

func handleJoinRooms(ctx http.TdkContext) error {
	return nil
}

func dummyFunc(ctx http.TdkContext) error {
	return nil
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
