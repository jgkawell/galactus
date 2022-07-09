package util

import (
	"encoding/base64"

	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
)

func GenerateBase64URLEncodedId(logger l.Logger) (string, l.Error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", logger.WrapError(l.NewError(err, "failed to generate random uuid"))
	}
	b, err := u.MarshalBinary()
	if err != nil {
		return "", logger.WrapError(l.NewError(err, "failed to marshal uuid to bytes"))
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
