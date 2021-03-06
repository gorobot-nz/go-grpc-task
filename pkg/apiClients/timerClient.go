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

func (c *TimerClient) CreateTimer(timer *challenge.Timer) error {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s", timer.Name), nil)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusGatewayTimeout {
		createReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s/%d", timer.Name, timer.Seconds), nil)
		createRes, err := c.client.Do(createReq)
		if err != nil {
			return err
		}
		if createRes.StatusCode != 200 {
			return errors.New("Something goes wrong")
		}
		return nil
	}
	return nil
}

func (c *TimerClient) GetRemainingSeconds(timer *challenge.Timer) (json.Number, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://timercheck.io/%s", timer.Name), nil)
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("Timer off")
	}
	var timerResponse responses.TimerResponse
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyText, &timerResponse)
	if err != nil {
		return "", err
	}
	return timerResponse.SecondsRemaining, nil
}
