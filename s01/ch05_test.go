package s01

import (
	"reflect"
	"testing"
)

func TestEncryptStringWithRepeatedXor(t *testing.T) {
	tests := []struct {
		src  string
		xor  string
		want string
	}{
		{
			"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
			"ICE",
			"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" +
				"a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
		},
	}
	for _, tt := range tests {
		got := EncryptStringWithRepeatedXor(tt.src, tt.xor)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("EncryptStringWithRepeatedXor() = %v, want %v", got, tt.want)
		}
	}
}
