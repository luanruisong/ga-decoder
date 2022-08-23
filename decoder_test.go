package ga_decoder

import (
	"testing"
)

func TestDecoderMigrationPayload(t *testing.T) {
	url := `otpauth-migration://offline?data={data str}`
	res, err := DecoderMigrationPayload(url)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range res.OtpParameters {
		t.Log(DecodeSecret(v.Secret))
	}
}
