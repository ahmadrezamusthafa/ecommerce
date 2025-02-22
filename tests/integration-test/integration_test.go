package integrationtest

import (
	"encoding/json"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity/requests"
	"github.com/mitchellh/mapstructure"
	"strings"
	"testing"
)

const (
	Url = "http://localhost:8005/api/v1"
)

func TestSubmitOrder(t *testing.T) {
	// register user
	registerUserReq, _ := json.Marshal(requests.CreateUserRequest{
		Email:    "int-user@test.com",
		Username: "int-user",
		Name:     "user integration",
		Password: "123",
	})
	resp := httpPost(Url+"/user/register", registerUserReq, map[string]string{}, t)
	if resp.Status == "error" && !strings.Contains(resp.Message, "already in use") {
		t.Error("Unexpected error when registering user", resp.Message)
		return
	}

	// login user
	loginUserReq, _ := json.Marshal(requests.UserLoginRequest{
		Email:    "int-user@test.com",
		Password: "123",
	})
	resp = httpPost(Url+"/user/login", loginUserReq, map[string]string{}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when login:", resp.Message)
		return
	}
	var token string
	if dataMap, ok := resp.Data.(map[string]interface{}); ok {
		err := mapstructure.Decode(dataMap["token"], &token)
		if err != nil {
			t.Error(err.Error())
		}
	}

	// add to cart
	addToCartReq, _ := json.Marshal(requests.AddToCartRequest{
		ProductID: 1,
		Quantity:  1,
	})
	resp = httpPost(Url+"/cart/items", addToCartReq, map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when add to cart:", resp.Message)
		return
	}
}
