package tests

import (
	"archive/zip"
	"bytes"
	"flag"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/gen/mocks"
	"test/internal/app/module1/domain/services"
	"testing"
)

func TestUserDetailApi(t *testing.T) {
	flag.Parse()
	us := new(mocks.UserDetailService)
	us.On("GetUserDetail", mock.AnythingOfType("uint64")).Return(&services.UserDetail{}, nil)
	api, err := CreateUserDetailAPI(*resourcesPath, us)
	if err != nil {
		t.Fatalf("get userDetail api  error,%+v", err)
	}
	resp := callAPI(api.API, "GET", "/detail?id=1", nil)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHttp(t *testing.T) {
	var b bytes.Buffer
	zipwriter := zip.NewWriter(&b)
	filewriter, err := zipwriter.Create("testdata.txt")
	if err != nil {
		panic(err)
	}
	_, err = filewriter.Write([]byte("testdata"))
	if err != nil {
		panic(err)
	}
	err = zipwriter.Close()
	if err != nil {
		panic(err)
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles",
		httpmock.NewBytesResponder(200, b.Bytes()))

	br := bytes.NewReader(b.Bytes())
	reader, err := zip.NewReader(br, br.Size())
	if err != nil {
		panic(err)
	}
	_, err = reader.Open("testdata.txt")
	if err != nil {
		panic(err)
	}
}
