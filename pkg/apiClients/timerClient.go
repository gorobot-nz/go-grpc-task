package apiclients

import (
	"encoding/json"
	"errors"
	"fmt"
	challenge "github.com/gorobot-nz/go-grpc-task/pkg/gen/pkg/proto"
	"github.com/gorobot-nz/go-grpc-task/pkg/responses"
	"io/ioutil"
	"log"
	"net/http"
)

type TimerClient struct {
	client *http.Client
}

func NewTimerClient() *TimerClient {
	return &TimerClient{client: &http.Client{}}
}

func (c *TimerClient) CreateTimer(timer *challenge.Timer) (bool, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s", timer.Name), nil)
	res, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode == http.StatusGatewayTimeout {
		createReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s/%d", timer.Name, timer.Seconds), nil)
		createRes, err := c.client.Do(createReq)
		if err != nil {
			return false, err
		}
		if createRes.StatusCode != 200 {
			return false, errors.New("Something goes wrong")
		}
		return false, nil
	}
	return true, nil
}

func (c *TimerClient) GetRemainingSeconds(timer *challenge.Timer) (int, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s", timer.Name), nil)
	res, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	var timerResponse responses.TimerResponse
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyText, &timerResponse)
	if err != nil {
		return 0, err
	}
	return timerResponse.SecondsRemaining, nil
}
