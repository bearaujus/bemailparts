# BEmailParts - Email Parsing Utilities in Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bearaujus/bemailparts/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bearaujus/bemailparts)](https://goreportcard.com/report/github.com/bearaujus/bemailparts)

BEmailParts is a simple and efficient utility for parsing and handling email addresses. It provides a structured way 
to extract different parts of an email address and reassemble it with ease. This library is useful for applications 
requiring email validation or detailed email manipulation.

## Installation

To install BEmailParts, run the following command:

```shell
go get github.com/bearaujus/bemailparts
```

## Import

```go
import "github.com/bearaujus/bemailparts"
```

## Features

- Parse an email address into its components (username, domain, domain name, and TLD).
- Validate the email format using a regular expression.
- Rebuild the email address from its components.

## Usage

### 1. Basic Parsing Workflow

Use the New function to parse an email into its parts:
```go
e, err := bemailparts.New("test@domain.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println(e.Email())               // Output: test@domain.com"
fmt.Println(e.Username())            // Output: test
fmt.Println(e.Domain())              // Output: domain.com
fmt.Println(e.DomainName())          // Output: domain
fmt.Println(e.DomainTLD())           // Output: .com
fmt.Println(e.DomainTLDWithoutDot()) // Output: com
```

### 2. Advanced Creation by Parts

You can also create an email from its components using helper functions:
```go
e, err := bemailparts.NewFromUsernameAndDomain("test", "domain.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
```
```go
e, err := bemailparts.NewFromFullParts("test", "domain", "com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
```

### 3. Rebuilding the Email

The Email() function reassembles the email from its parts:
```go
e, err := bemailparts.New("test@domain.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println("Email:", e.Email())   // Output: test@domain.com
```

## License

This project is licensed under the MIT License - see
the [LICENSE](https://github.com/bearaujus/bemailparts/blob/master/LICENSE) file for details.