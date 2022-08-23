package ga_decoder

import (
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/luanruisong/ga-decoder/pb"
	"net/url"
)

func DecoderMigrationPayload(strUrl string) (*pb.MigrationPayload, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}
	data := u.Query().Get("data")
	if len(data) == 0 {
		return nil, errors.New("can not get data")
	}
	return DecoderMigrationPayloadOnlyData(data)
}

func DecoderMigrationPayloadOnlyData(data string) (*pb.MigrationPayload, error) {
	b, _ := base64.StdEncoding.DecodeString(data)
	res := new(pb.MigrationPayload)
	if err := proto.Unmarshal(b, res); err != nil {
		return nil, err
	}
	return res, nil
}

func DecodeSecret(secret []byte) string {
	return base32.StdEncoding.EncodeToString(secret)
}
