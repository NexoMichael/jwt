package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHelpMessage(t *testing.T) {
	var b bytes.Buffer
	jwtcli(&b, []string{})
	require.Equal(t, "jwt - command line JWT token parser\n\nUsage:\n    jwt [encoded token]\n\n", b.String())
}

func TestBadToken(t *testing.T) {
	var b bytes.Buffer
	jwtcli(&b, []string{"", "something"})
	require.Equal(t, "Token is not valid: token should consist of 3 parts separated by dot symbol\n", b.String())
}

func TestNonParseableTokenHeader(t *testing.T) {
	var b bytes.Buffer
	jwtcli(&b, []string{"", "a.b.c"})
	require.Equal(t, "Token is not valid: failed to parse token header: failed to decode token part: illegal base64 data at input byte 1\n", b.String())
}

func TestNonParseableTokenBody(t *testing.T) {
	var b bytes.Buffer
	jwtcli(&b, []string{"", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.b.c"})
	require.Equal(t, "Token is not valid: failed to parse token body: failed to decode token part: illegal base64 data at input byte 1\n", b.String())
}

func TestParseToken(t *testing.T) {
	var b bytes.Buffer
	iat := time.Unix(1516239022, 0).Format(time.RFC822)
	nbf := time.Unix(1516249022, 0).Format(time.RFC822)
	exp := time.Unix(1516259022, 0).Format(time.RFC822)
	jwtcli(&b, []string{"", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlRlc3QgVG9rZW4iLCJpYXQiOjE1MTYyMzkwMjIsIm5iZiI6MTUxNjI0OTAyMiwiZXhwIjoxNTE2MjU5MDIyfQ.DQJ8SA18nhH0Zh6HaxUAsFwsa37Fp82rVJvnWJfHgwU"})
	require.Equal(t, fmt.Sprintf("\x1b[32m\n✻ Header\n{\n\t\"alg\": \"HS256\",\n\t\"typ\": \"JWT\"\n}\n\x1b[33m\n✻ Body\n{\n\t\"exp\": 1516259022,\n\t\"iat\": 1516239022,\n\t\"name\": \"Test Token\",\n\t\"nbf\": 1516249022,\n\t\"sub\": \"1234567890\"\n}\n\x1b[90mIssued at: %s\nNot before: %s\nExpires at: %s\n\x1b[31m\n✻ Signature\nDQJ8SA18nhH0Zh6HaxUAsFwsa37Fp82rVJvnWJfHgwU\n\x1b[0m", iat, nbf, exp), b.String())
}

func TestBadTokenPart(t *testing.T) {
	var dst map[string]interface{}
	require.NotNil(t, parsePart(&dst, "bm90IGpzb24"))
}

func TestTryParseTime(t *testing.T) {
	now := time.Now().Unix()
	tok := token{
		body: map[string]interface{}{
			"1": "not a number",
			"2": nil,
			"3": -100,
			"4": float64(now),
		},
	}
	require.Equal(t, "undefined", tok.tryParseTime("0"))
	require.Equal(t, "invalid date", tok.tryParseTime("1"))
	require.Equal(t, "undefined", tok.tryParseTime("2"))
	require.Equal(t, "invalid date", tok.tryParseTime("3"))
	require.Equal(t, time.Unix(now, 0).Format(time.RFC822), tok.tryParseTime("4"))
}

func TestMain(t *testing.T) {
	defer func() {
		require.Nil(t, recover())
	}()
	main()
}
