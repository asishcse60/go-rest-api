package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil{
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T){
	client := resty.New()
	resp, err := client.R().SetBody(`{"slug":"/","author":"A-12345","body":"hello world"}`).Post(BASE_URL + "/api/comment")
	if err != nil{
		t.Fail()
	}
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}