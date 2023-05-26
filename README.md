# Govalid

A small, simple and dependency-free Go library to validate strings.

The primary use of this library is to validate and parse settings coming from environment variables and URL query parameters, where the only type defined is the string.

💁 This project is still sub `v1.0.0` so is subject to breaking changes with future releases.

## Features

- [x] Validate and parse binary choices such as `yes` / `no`
  - [x] Set string values signaling an enabled status
  - [x] Set string values signaling a disabled status
- [x] Validate and parse separated values such as CSV
  - [x] Custom separator option
  - [x] List of accepted values option
  - [x] Lowercase all values option
- [x] Validate and parse ports
  - [x] Listening port option
  - [x] Non root privileged ports allowed option
- [x] Validate addresses
  - [x] Listening address option
- [x] Validate and parse root URLs
- [x] Validate and parse URLs
  - [x] Set schemes allowed options
- [x] Validate email addresses
  - [x] MX Lookup option
- [x] Validate and parse duration strings
  - [x] Allow negative durations option
  - [x] Allow zero durations option
- [x] Validate and parse integers
  - [x] Range option
  - [x] Minimum option
  - [x] Maximum option
- [x] Validate common digest strings such as `sha256hex` or `md5hex`

## Usage

### With subpackages

There are subpackages for each validation required.
For example using the `github.com/qdm12/govalid/address` package:

```go
package main

import (
    "fmt"

    "github.com/qdm12/govalid/address"
)

func main() {
    const s = ":8000"
    const uid = 1000

    addr, err := address.Validate(s, address.OptionListening(uid))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Address: ", addr)
}

```

### With `Validator`

If you prefer, I also defined a `Validator` object at the root of this repository, for example:

```go
package main

import (
    "fmt"

    "github.com/qdm12/govalid"
    "github.com/qdm12/govalid/address"
)

func main() {
    const s = ":8000"
    const uid = 1000

    validator := govalid.New()

    addr, err := validator.ValidateAddress(s, address.OptionListening(uid))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Address: ", addr)
}

```

## TODOs

### More validations

- [ ] Validate phone numbers
  - [ ] International number option
- [ ] Validate domain names and/or hostnames
