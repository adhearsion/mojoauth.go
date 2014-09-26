[![Build Status](https://travis-ci.org/mojolingo/mojoauth.go.svg?branch=develop)](http://travis-ci.org/mojolingo/mojoauth.gp)

# mojoauth

[MojoAuth](http://mojolingo.com/mojoauth) is a set of standard approaches to cross-app authentication based on [Hash-based Message Authentication Codes](http://en.wikipedia.org/wiki/Hash-based_message_authentication_code) (HMAC), inspired by ["A REST API For Access To TURN Services"](http://tools.ietf.org/html/draft-uberti-behave-turn-rest).

## Usage

```go
package main

import (
  "fmt"
  "github.com/mojolingo/mojoauth.go"
)

func main() {
  // Generate a shared secret
  secret = mojoauth.CreateSecret
    // => "XyD+xeJHivzbOUe3vwdU6Z5vDe/vio34MxKX8HYViR0+p4t/NzaIpbK+9VwX\n5qHCj7m4f7UNRXgOJPXzn6MT0Q==\n"

  // Create temporary credentials
  credentials = mojoauth.CreateCredentials(id: 'foobar', secret: secret)
    // => {:username=>"1411837760:foobar", :password=>"wb6KxLj6NXcUaqNb1SlHH1V3QHw=\n"}

  // Test credentials
  mojoauth.TestCredentials({username: "1411837760:foobar", password: "wb6KxLj6NXcUaqNb1SlHH1V3QHw=\n"}, secret: secret)
    // => "foobar"
  mojoauth.TestCredentials({username: "1411837760:foobar", password: "wrongpassword"}, secret: secret)
    // => false

  // 1 day later
  mojoauth.TestCredentials({username: "1411837760:foobar", password: "wb6KxLj6NXcUaqNb1SlHH1V3QHw=\n"}, secret: secret)
    // => false
}
```

## Contributing

1. [Fork it](https://github.com/mojolingo/mojoauth.go/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
