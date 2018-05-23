# jwt

[![BuildStatus Widget]][BuildStatus Result]
[![CodeCovWidget]][CodeCovResult]
[![GoReport Widget]][GoReport Status]
[![GoDoc Widget]][GoDoc Link]

[BuildStatus Result]: https://travis-ci.org/NexoMichael/jwt
[BuildStatus Widget]: https://travis-ci.org/NexoMichael/jwt.svg?branch=master

[GoReport Status]: https://goreportcard.com/report/github.com/NexoMichael/jwt
[GoReport Widget]: https://goreportcard.com/badge/github.com/NexoMichael/jwt

[CodeCovResult]: https://coveralls.io/github/NexoMichael/jwt
[CodeCovWidget]: https://coveralls.io/repos/github/NexoMichael/jwt/badge.svg?branch=master

[GoDoc Link]: http://godoc.org/github.com/NexoMichael/jwt
[GoDoc Widget]: http://godoc.org/github.com/NexoMichael/jwt?status.svg

Command line JWT token parser written in Golang

## What is this?

`jwt` is command line tool for parsing [JSON Web Tokens](http://jwt.io/) (JWT)

## Installation

Use the `go` command:

	$ go get -u github.com/NexoMichael/jwt-cli


## Usage

```
$ jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlRlc3QgVG9rZW4iLCJpYXQiOjE1MTYyMzkwMjIsIm5iZiI6MTUxNjI0OTAyMiwiZXhwIjoxNTE2MjU5MDIyfQ.DQJ8SA18nhH0Zh6HaxUAsFwsa37Fp82rVJvnWJfHgwU

✻ Header
{
	"alg": "HS256",
	"typ": "JWT"
}

✻ Body
{
	"exp": 1516259022,
	"iat": 1516239022,
	"name": "Test Token",
	"nbf": 1516249022,
	"sub": "1234567890"
}
Issued at: 18 Jan 18 04:30 GMT
Not before: 18 Jan 18 07:17 GMT
Expires at: 18 Jan 18 10:03 GMT

✻ Signature
DQJ8SA18nhH0Zh6HaxUAsFwsa37Fp82rVJvnWJfHgwU
```

## Requirements

jwt-cli package requires Go >= 1.5.


## Copyright

Copyright (C) 2018 by Mikhail Kochegarov <mikhail@kochegarov.biz>.

UUID package releaed under MIT License.
See [LICENSE](https://github.com/NexoMichael/shortuuid/blob/master/LICENSE) for details.