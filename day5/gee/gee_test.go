package gee

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	// 待写
	v1 := r.Group("/v1")
	v1 := r.Group("/v1")
	v2 := r.Group("/v2")
	v3 := r.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2 ")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
