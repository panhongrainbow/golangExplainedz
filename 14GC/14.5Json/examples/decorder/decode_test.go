package decorder

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var jsonData = []byte(`{"Name":"john", "Age":30}`)

func Test_decoder(t *testing.T) {
	for i := 0; i < 10; i++ {
		var u User
		err := Unmarshal(jsonData, &u)
		require.NoError(t, err)

		fmt.Printf("%+v", u)
	}
}

func Benchmark_My_Unmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var u User
		_ = Unmarshal(jsonData, &u)
	}
}

func Benchmark_Json_Unmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var u User
		_ = json.Unmarshal(jsonData, &u)
	}
}

func Benchmark_Parser_Unmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = U2(jsonData)
	}
}
