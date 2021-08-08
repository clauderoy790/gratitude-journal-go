package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/server"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

var baseTestUrl = "http://localhost:8080/"

type ApiSuite struct {
	suite.Suite
	client http.Client
}

func (s *ApiSuite) SetupSuite() {
	go func() {
		server := server.New(context.Background())
		server.Start()
	}()
	time.Sleep(150 * time.Millisecond)
	s.client = http.Client{}
}

func (s *ApiSuite) TearDownSuite() {

}

func (s *ApiSuite) TestHome() {
	res, err := DoTestRequest(s, baseTestUrl, http.MethodGet, "", http.Header{})
	s.NoError(err, "Got an error in home: %v", err)
	if err != nil {
		fmt.Println("error in request: ", err)
		return
	}
	defer res.Body.Close()

	m := make(server.M)
	json.NewDecoder(res.Body).Decode(&m)
	exp := server.M{"message": "welcome to daily gratitude"}
	s.Equalf(exp, m, "home response error, expected %v, got %\n", exp, m)
}

func TestApiEndpoints(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}

func DoTestRequest(a *ApiSuite, url, method string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return a.client.Do(request)
}