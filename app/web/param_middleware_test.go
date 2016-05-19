package web

import (
	"net/http"
	"testing"
)

func testParameterValue(t *testing.T) {
	expected := "123aBc"
	request, err := http.NewRequest("GET", "http://example.com?token=123Abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	value, err := Param(request, "token")
	if err != nil {
		t.Fatal(err)
	}

	if value != expected {
		t.Fatalf("Expected %s but got %s", expected, value)
	}
}

func testErrorOnAccessingNoneExistingParameterValue(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Param(req, "token")
	if err == nil {
		t.Fatal("Not handling none existing parameter value")
	}
}

func testMultipleParameters(t *testing.T) {
	e1 := "123Abc"
	e2 := "bca"
	request, err := http.NewRequest("GET", "http://example.com?token=123Abc&tok=bca", nil)
	if err != nil {
		t.Fatal(err)
	}
	v1, err := Param(request, "token")
	if err != nil {
		t.Fatal(err)
	}
	v2, err := Param(request, "tok")
	if err != nil {
		t.Fatal(err)
	}

	if v1 != e1 {
		t.Logf("Expected %s but got %s", e1, v1)
	}

	if v2 != e2 {
		t.Logf("Expected %s but got %s", e1, v1)
	}
}
