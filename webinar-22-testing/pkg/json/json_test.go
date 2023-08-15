package json_test

import (
	"testing"
	"webinar-22/pkg/json"
)

func BenchmarkMapUnmarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := json.UnmarshalToMap()
		if err != nil {
			b.Log(err.Error())
			b.Fail()
		}
	}
}

func BenchmarkStructUnmarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := json.UnmarshalToStruct()
		if err != nil {
			b.Log(err.Error())
			b.Fail()
		}
	}
}
