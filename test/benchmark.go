package main

import (
"testing"
)

type Request struct {
	Field string
}

var r = Request{Field: "a"}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Request{Field: "b"}
	}
}

var p = &Request{Field: "a"}

func BenchmarkStructPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p = &Request{Field: "b"}
	}
}