package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil{
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}