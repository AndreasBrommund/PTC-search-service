package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"strconv"
)

func TestNoneExistingGroupParameter(t *testing.T) {
	if ok,msg := testEndpointParameter(
		"GET",
		"http://localhost:8080/api/1/tags?&startDate=2012&endDate=2013&limit=12",
		`{"error_msg":"No parameter found for key: group"}`,
		400,
		Tags);!ok{
		t.Error(msg)
	}
}
func TestNoneExistingStartDateParameter(t *testing.T) {
	if ok,msg := testEndpointParameter(
		"GET",
		"http://localhost:8080/api/1/tags?group=a&endDate=2013&limit=12",
		`{"error_msg":"No parameter found for key: startDate"}`,
		400,
		Tags);!ok{
		t.Error(msg)
	}
}
func TestNoneExistingEndDateParameter(t *testing.T) {
	if ok,msg := testEndpointParameter(
		"GET",
		"http://localhost:8080/api/1/tags?group=a&startDate=12&limit=12",
		`{"error_msg":"No parameter found for key: endDate"}`,
		400,
		Tags);!ok{
		t.Error(msg)
	}
}
func TestNoneExistingLimitParameter(t *testing.T) {
	if ok,msg := testEndpointParameter(
		"GET",
		"http://localhost:8080/api/1/tags?group=a&startDate=12&endDate=2013",
		`{"error_msg":"No parameter found for key: limit"}`,
		400,
		Tags);!ok{
		t.Error(msg)
	}
}
func TestCastLimitFromStringToInt(t *testing.T){
	if ok,msg := testEndpointParameter(
		"GET",
		"http://localhost:8080/api/1/tags?group=a&startDate=12&endDate=2013&limit=a",
		`{"error_msg":"Could not convert limit to string."}`,
		400,
		Tags);!ok{
		t.Error(msg)
	}
}

func testEndpointParameter(method,url,expectedBody string,statusCode int,
	handler func(w http.ResponseWriter,r *http.Request)) (bool,string) {
	req, _ := http.NewRequest(method,url, nil)
	response := httptest.NewRecorder()

	handler(response,req)

	if ok,msg := checkResponse(expectedBody,statusCode,response);!ok{
		return false,msg
	}
	return true,""
}

func checkResponse(expectedBody string,expectedCode int,response *httptest.ResponseRecorder) (bool,string){
	if response.Code != expectedCode {
		return false,"Expected: "+strconv.Itoa(expectedCode)+", got: "+strconv.Itoa(response.Code)
	}
	if strings.TrimSpace(response.Body.String()) != expectedBody {
		return false,"Expected: "+expectedBody+", got: "+strings.TrimSpace(response.Body.String())
	}
	return true,""
}

