package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkFr2ydznpmnuxhd7PxJ5
jxrabx0fW92x0dMML88UaRqF/027GhUYGkZwl/X6IED/cfhy4Zwzyc4cyiCKqg4H
vvZu0qcWmHUTHGLRFfydmH3Yg+A7tOQwEzAgsAUlPmQt01JaUpCLWhX0LPaSePj4
TDDzsBeoBumptpO2CCHaOHmq9MozpnLooGX4PQ8rSLAPCHKJm7FnFnPDRcm0A9Os
EMLZzTDXvBoMp30najyOPJACbQ4dScxKJu0WRehD38B3e9UJPX7Wl20taZtNLhYO
UdtAvcllnkzWfhdbSqGT5ydrUPcHaMo3w823YzjFU72d7Q+ZzWx6qw/cOaV1dNkQ
ZQIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{PublicKey: key}
	tkn := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjMzMzkzNzgsImlzcyI6ImNvb2xjYXIvYXV0aCIsInN1YiI6IjEyMzQ1Njc4OTAwODk3OTg3ODk4eTg3OTh5ODl5dGZ0eSJ9.DJtkcsFLGkMW7yOF87iOIHP9zOi7LICFRFSf5N9KY5TqrOKX5owa9PLvgPaFiG2_0V6Cmt_RMmjlSzq9sWSbS7m3GrY05VXFAj-Sr-6Fs9oxO-dOdLZXpPJNVrOGIQjlcH5xcqInePev8_MhvYgK6_su1gACGK_GflrL1OCPnDzDsBzs_UIHtecLpWnujHEihAry1iAALx1h8n_E2vud-UxjU8BP_rfcIcyMVznzA40b4MDY0CnYJ0lev35DSBHdoLLINFlyacCt5c0Cyyt5sUdO4WB0pLj-3-t-Wxd2a5LQe8clYKt2hztj5KdukIlXH8IKeVS5p0ofJvvrrYX2jQ"
	jwt.TimeFunc = func() time.Time {
		return time.Now()
	}
	accountID, err := v.Verify(tkn)
	if err != nil {
		t.Fatalf("verify failed: %v", err)
	}
	want := "12345678900897987898y8798y89ytfty"
	if accountID != want {
		t.Errorf("wrong id, want:%q,got:%q", want, accountID)
	}
}
