package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAkFr2ydznpmnuxhd7PxJ5jxrabx0fW92x0dMML88UaRqF/027
GhUYGkZwl/X6IED/cfhy4Zwzyc4cyiCKqg4HvvZu0qcWmHUTHGLRFfydmH3Yg+A7
tOQwEzAgsAUlPmQt01JaUpCLWhX0LPaSePj4TDDzsBeoBumptpO2CCHaOHmq9Moz
pnLooGX4PQ8rSLAPCHKJm7FnFnPDRcm0A9OsEMLZzTDXvBoMp30najyOPJACbQ4d
ScxKJu0WRehD38B3e9UJPX7Wl20taZtNLhYOUdtAvcllnkzWfhdbSqGT5ydrUPcH
aMo3w823YzjFU72d7Q+ZzWx6qw/cOaV1dNkQZQIDAQABAoIBACKng7lgK8hag/TQ
LEku8Tm2k5u7HQ0bwpWBmdpcVyUINgTdLz6Ks9eS83K3nU1i7S/6GfIpYpFexrCL
cV5zsc5ZMK+nZxhAes8EtfcWRusVwwWdrhe19AiXNDGteoxo1kt16LMLejitLoog
w3cBJfJ8ifPLKg5Gx/RJr5hyMAYAarCWy65YoRVIIhOrGavyqZi63XiTEp7fyVMF
TkDVPWftynRJy0sz3vp+Ok7A3eclOKBD4Vz/q3NXejGozkth84wRlcAkLG5wKsEq
cPWZ14oSw777b3T3hV4ljXmtUKj1NbvrnUx2TXtJxrThHRSDwV9lfgVI2Y4iXVF4
TVFKuIECgYEA0b+2/2P4o5pP2S6/INWNBlCXA/ibl1Wf2HD2RbHzCNn1YD2FYuh6
4zwOIQADJf7lz6r4SYj2k2Y59VOVCSJYQ3OoOMhKOXTtBR6lc6Vm/6cmJoqB1vbl
aJn/Ld7yhAAM8pzLkyvfEmFhvdBU9DfMFelEyqcphn78Xdin6ueCh6ECgYEAsC/L
jmKoXCF78VYNC57aX3Bm1enj2tGY3VyV1rc5jr3ivEoa26mryRpAbTgB6Rx4CW/T
Dgoz3rL75wzdLiSEB1mYfGdPlSsEZJsggvwmW53EQQtpTBjIzxYU9jlPgpzTc4b+
ioMvfMTs6RzTnG2QbLVuRCU45o2b6CWCAt7TQkUCgYEAvwayvKAo5593j14Ste2n
9ZNaJkS6N7bE0JP4xvrNVEdlQZRmMfF3UhL07zsaovUmCd81J4u0vgPBT1wjBOGh
rzTbhXNsni2OXDZQCyYdy0JI7ZsBq2zK/FwcWoONLYj6Qc9pXIz9KblFEmF1rcJP
fbkobMSXfiWS5EmYjMjySWECgYBUEJoJeB1oyDlBL5PN5Z/ARftrOcwUTkmn5VNB
Pe9iokubF6i1AsIKlFIFSuHufjzwE8EaQ9f3/GKhHcwzBg1RDHjrcsfQHtRbxIDA
vtr2f9JyTqWRP4og9SJPUY0UfwuNZe3x3SI9YCDCIZT+YHC2zeKs9S2vJAYtwCfG
gtc+GQKBgDWZRH9mIs7NqciahKNNRjDMU1FTzveu6jdFLKaRDo8VvmT8yAGg9SNS
mkMWvLXoZ7LJBV8WRf6Qa05kK/BCvxu03Dwys9ilETNT6xkiD1HThumESnYLp5Dl
dfsJSCKrYj/qlq0ENlYSbSce9+TllkWLO3vL49G/8JPcCcBI55Zj
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	pem, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", pem)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	token, err := g.GenerateToken("12345678900897987898y8798y89ytfty", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token %v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MDA4OTc5ODc4OTh5ODc5OHk4OXl0ZnR5In0.fDFxhZPVI8hdnRc81h71fe7mMVwVfrHYae_-MDSv7jbI0qwH7Z8P_xae1a72NTjzrIR-RvrReQ8KjPYuMY7CzMlf8q0f6lMPCC-0QnCRYONSRqggA7N3adMFxuvqvfGrVvYf9Xhwo0PdPG_uTVExD6JI3Q2vnUPBPFuGvl_nTxm_Yesb94wiqi2rYfxLzNpfSEQ6utSIbimE21mxhqGz2m12i6wXKh_SbhZkuSGyjhxmjdytN_SsKR1JRnhLcvJM8DKFQ9KJMMZ0hTvzbv0MwBEGagyoLXUyzsOLyMzf_OYQSzQFd12bnGXFF8xN7-zxC-5kiU9NDCHlYhlV95GiPw"
	if token != want {
		t.Errorf("wrong token generated,\n want:%q,\n got:%q", want, token)
	}
}
