package main

import (
	"testing"
)

func BenchmarkSerlializeToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToJSON(metadata)
	}
}

func BenchmarkSerializeToXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToXML(metadata)
	}
}

func BenchmarkToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToProto(genMetadata)
	}
}