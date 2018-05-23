package main // import "github.com/NexoMichael/jwt"

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var helpMsg = `jwt - command line JWT token parser

Usage:
    jwt [encoded token]
`

func main() {
	jwtcli(os.Stdout, os.Args)
}

func jwtcli(w io.Writer, args []string) {
	if len(args) != 2 {
		fmt.Fprintln(w, helpMsg)
		return
	}

	t, err := parseToken(args[1])
	if err != nil {
		fmt.Fprintf(w, "Token is not valid: %s\n", err.Error())
		return
	}

	t.print(w)
}

// token represents typical JWT token structure
type token struct {
	header    map[string]interface{}
	body      map[string]interface{}
	signature string
}

// escape character for terminal color change
const escape = "\x1b"

// Foreground text colors
const (
	fgRed int = iota + 31
	fgGreen
	fgYellow
	fgHiBlack = 90
	fgReset   = 0
)

// print converts token to human-readable representation and put it to io.Writer
func (t token) print(w io.Writer) {
	header, _ := json.MarshalIndent(t.header, "", "\t")
	body, _ := json.MarshalIndent(t.body, "", "\t")

	setColor(w, fgGreen)
	printPart(w, "Header", string(header))
	setColor(w, fgYellow)
	printPart(w, "Body", string(body))
	setColor(w, fgHiBlack)
	fmt.Fprintf(w, "Issued at: %s\n", t.tryParseTime("iat"))
	fmt.Fprintf(w, "Not before: %s\n", t.tryParseTime("nbf"))
	fmt.Fprintf(w, "Expires at: %s\n", t.tryParseTime("exp"))
	setColor(w, fgRed)
	printPart(w, "Signature", t.signature)
	setColor(w, fgReset)
}

// tryParseTime parses unix-time encoded claim from JWT token body
func (t token) tryParseTime(key string) string {
	v, found := t.body[key]
	if !found || v == nil {
		return "undefined"
	}
	val, ok := v.(float64)
	if !ok {
		return "invalid date"
	}
	return time.Unix(int64(val), 0).Format(time.RFC822)
}

// parseToken parses JWT token string into token object
func parseToken(tokenString string) (token, error) {
	var t token
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return t, errors.New("token should consist of 3 parts separated by dot symbol")
	}

	if err := parsePart(&t.header, parts[0]); err != nil {
		return t, errors.Wrap(err, "failed to parse token header")
	}

	if err := parsePart(&t.body, parts[1]); err != nil {
		return t, errors.Wrap(err, "failed to parse token body")
	}

	t.signature = parts[2]
	return t, nil
}

// decodePart decodes part of the JWT token using base64 algorithm and ignoring stripped padding
func decodePart(part string) ([]byte, error) {
	if l := len(part) % 4; l > 0 {
		part += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(part)
}

// parsePart parses part of the token and provide it in human-readable format
func parsePart(dst interface{}, part string) (err error) {
	var src []byte
	if src, err = decodePart(part); err != nil {
		return errors.Wrap(err, "failed to decode token part")
	}

	if err = json.Unmarshal(src, dst); err != nil {
		return errors.Wrap(err, "failed to unmarshal token part")
	}
	return nil
}

// printPart prints part of the token to io.Writer
func printPart(w io.Writer, caption, part string) {
	fmt.Fprintf(w, "\n✻ %s\n%s\n", caption, part)
}

func setColor(w io.Writer, color int) {
	fmt.Fprintf(w, "%s[%dm", escape, color)
}
