package integrationtest

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity/requests"
	"github.com/go-playground/assert/v2"
	"github.com/mitchellh/mapstructure"
	"strings"
	"testing"
	"time"
)

const (
	Url = "http://localhost:8005/api/v1"
)

func TestSubmitOrder(t *testing.T) {
	unix := time.Now().Unix()

	email := fmt.Sprintf("int-user-%d@test.com", unix)
	username := fmt.Sprintf("int-user-%d", unix)
	password := "123"

	// register user
	registerUserReq, _ := json.Marshal(requests.CreateUserRequest{
		Email:    email,
		Username: username,
		Name:     "user integration",
		Password: password,
	})
	resp := httpPost(Url+"/user/register", registerUserReq, map[string]string{}, t)
	if resp.Status == "error" && !strings.Contains(resp.Message, "already in use") {
		t.Error("Unexpected error when registering user", resp.Message)
		return
	}

	// login user
	loginUserReq, _ := json.Marshal(requests.UserLoginRequest{
		Email:    email,
		Password: password,
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

	// deposit balance
	depositReq, _ := json.Marshal(requests.DepositBalanceRequest{
		Amount: 5000000,
	})
	resp = httpPost(Url+"/account/deposit", depositReq, map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when deposit balance:", resp.Message)
		return
	}

	// withdraw balance
	withdrawReq, _ := json.Marshal(requests.WithdrawBalanceRequest{
		Amount: 50000,
	})
	resp = httpPost(Url+"/account/withdraw", withdrawReq, map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when withdraw balance:", resp.Message)
		return
	}

	// get balance
	resp = httpGet(Url+"/account", map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when get product:", resp.Message)
		return
	}
	var account entity.Account
	err := mapstructure.Decode(resp.Data, &account)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equal(t, float64(4950000), account.Balance)

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

	// submit order
	resp = httpPost(Url+"/orders", nil, map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when add to cart:", resp.Message)
		return
	}

	// get balance
	resp = httpGet(Url+"/account", map[string]string{"Authorization": token}, t)
	if resp.Status == "error" {
		t.Error("Unexpected error when get product:", resp.Message)
		return
	}
	err = mapstructure.Decode(resp.Data, &account)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.NotEqual(t, float64(4950000), account.Balance)
}
