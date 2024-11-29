package bemailparts

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	emailSeparator  = "@"
	domainSeparator = "."
)

const (
	usernamePattern   = `^[a-zA-Z0-9._%+-]+`
	domainNamePattern = `[a-zA-Z0-9.-]+`
	domainTLDPattern  = `[a-zA-Z]+$`
	domainPattern     = domainNamePattern + `\` + domainSeparator + domainTLDPattern
	emailPattern      = usernamePattern + emailSeparator + domainPattern
)

var (
	usernameRegex   = regexp.MustCompile(usernamePattern)
	domainNameRegex = regexp.MustCompile(domainNamePattern)
	domainTLDRegex  = regexp.MustCompile(domainTLDPattern)
	domainRegex     = regexp.MustCompile(domainPattern)
	emailRegex      = regexp.MustCompile(emailPattern)
)

// BEmailParts defines an interface for managing and manipulating email components.
// Example email for reference: "john.doe@example.com".
type BEmailParts interface {
	// Email returns the full email address in the format "username@domain".
	// Example: "john.doe@example.com".
	Email() string

	// Username returns the username part of the email (before the '@').
	// Example: "john.doe" from "john.doe@example.com".
	Username() string

	// Domain returns the domain part of the email (after the '@').
	// Example: "example.com" from "john.doe@example.com".
	Domain() string

	// DomainName returns the domain name (e.g., "example" in "example.com").
	// Example: "example" from "john.doe@example.com".
	DomainName() string

	// DomainTLD returns the top-level domain (TLD) of the email (e.g., "com" in "example.com").
	// Example: ".com" from "john.doe@example.com".
	// Example 2: ".co.id" from "john.doe@example.co.id".
	DomainTLD() string

	// DomainTLDWithoutDot returns the top-level domain (TLD)
	// of the email (e.g., "com" in "example.com") without a leading dot.
	// Example: "com" from "john.doe@example.com".
	// Example 2: "co.id" from "john.doe@example.co.id".
	DomainTLDWithoutDot() string

	// SetUsername updates the username part of the email.
	// Example: If called with "jane.doe", the updated email will be "jane.doe@example.com".
	// Returns an error if the provided username is invalid.
	SetUsername(username string) error

	// SetDomain updates the domain part of the email.
	// Example: If called with "newdomain.org", the updated email will be "john.doe@newdomain.org".
	// Returns an error if the provided domain is invalid.
	SetDomain(domain string) error

	// SetDomainName updates the domain name part of the email.
	// Example: If called with "newexample", the updated email will be "john.doe@newexample.com".
	// Returns an error if the provided domain name is invalid.
	SetDomainName(domainName string) error

	// SetDomainTLD updates the top-level domain (TLD) of the email.
	// Example: If called with "org", the updated email will be "john.doe@example.org".
	// Returns an error if the provided TLD is invalid.
	SetDomainTLD(domainTLD string) error

	// String returns the string representation of the email address.
	// Example: "john.doe@example.com"
	String() string
}

type bEmailParts struct {
	username string
	domain   string
}

// New creates a new instance of BEmailParts by parsing a full email address.
//
// Parameters:
//
//	email: A valid email address in the format "username@domain".
//
// Returns:
//   - A BEmailParts instance representing the parsed email.
//   - An error if the email format is invalid (e.g., missing '@' or invalid characters).
//
// Example:
//
//	emailParts, err := New("john.doe@example.com")
//	if err != nil {
//	    log.Fatalf("Invalid email: %v", err)
//	}
//
//	fmt.Println(emailParts.Email())               // Output: john.doe@example.com
//	fmt.Println(emailParts.Username())            // Output: john.doe
//	fmt.Println(emailParts.Domain())              // Output: example.com
//	fmt.Println(emailParts.DomainName())          // Output: example
//	fmt.Println(emailParts.DomainTLD())           // Output: .com
//	fmt.Println(emailParts.DomainTLDWithoutDot()) // Output: com
func New(email string) (BEmailParts, error) {
	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmailFormat
	}

	parts := strings.Split(email, emailSeparator)
	username := parts[0]
	domain := parts[1]

	return &bEmailParts{
		username: username,
		domain:   domain,
	}, nil
}

// NewFromUsernameAndDomain creates a new instance of BEmailParts from a username and domain.
//
// Parameters:
//
//	username: The username part of the email (before the '@').
//	domain: The domain part of the email (after the '@').
//
// Returns:
//   - A BEmailParts instance representing the constructed email.
//   - An error if either the username or domain is invalid.
//
// Example:
//
//	emailParts, err := NewFromUsernameAndDomain("john.doe", "example.com")
//	if err != nil {
//	    log.Fatalf("Invalid username or domain: %v", err)
//	}
//
//	fmt.Println(emailParts.Email())               // Output: john.doe@example.com
//	fmt.Println(emailParts.Username())            // Output: john.doe
//	fmt.Println(emailParts.Domain())              // Output: example.com
//	fmt.Println(emailParts.DomainName())          // Output: example
//	fmt.Println(emailParts.DomainTLD())           // Output: .com
//	fmt.Println(emailParts.DomainTLDWithoutDot()) // Output: com
func NewFromUsernameAndDomain(username, domain string) (BEmailParts, error) {
	if !usernameRegex.MatchString(username) {
		return nil, ErrInvalidEmailUsernameFormat
	}
	if !domainRegex.MatchString(domain) {
		return nil, ErrInvalidEmailDomainFormat
	}
	return New(generateEmail(username, domain))
}

// NewFromFullParts creates a new instance of BEmailParts from a username, domain name, and domain TLD.
//
// Parameters:
//
//	username: The username part of the email (before the '@').
//	domainName: The domain name (e.g., "example" in "example.com").
//	domainTLD: The top-level domain (TLD) of the email (e.g., "com" in "example.com").
//
// Returns:
//   - A BEmailParts instance representing the constructed email.
//   - An error if any of the parts are invalid.
//
// Example:
//
//	emailParts, err := NewFromFullParts("john.doe", "example", "com")
//	if err != nil {
//	    log.Fatalf("Invalid email parts: %v", err)
//	}
//
//	fmt.Println(emailParts.Email())               // Output: john.doe@example.com
//	fmt.Println(emailParts.Username())            // Output: john.doe
//	fmt.Println(emailParts.Domain())              // Output: example.com
//	fmt.Println(emailParts.DomainName())          // Output: example
//	fmt.Println(emailParts.DomainTLD())           // Output: .com
//	fmt.Println(emailParts.DomainTLDWithoutDot()) // Output: com
func NewFromFullParts(username, domainName, domainTLD string) (BEmailParts, error) {
	if !domainNameRegex.MatchString(domainName) {
		return nil, ErrInvalidEmailDomainNameFormat
	}
	if !domainTLDRegex.MatchString(domainTLD) {
		return nil, ErrInvalidEmailDomainTLDFormat
	}
	return NewFromUsernameAndDomain(username, generateDomain(domainName, domainTLD))
}

func (e *bEmailParts) Email() string {
	return generateEmail(e.username, e.domain)
}

func (e *bEmailParts) Username() string {
	return e.username
}

func (e *bEmailParts) Domain() string {
	return e.domain
}

func (e *bEmailParts) DomainName() string {
	return e.domain[:strings.Index(e.domain, domainSeparator)]
}

func (e *bEmailParts) DomainTLD() string {
	return e.domain[strings.Index(e.domain, domainSeparator):]
}

func (e *bEmailParts) DomainTLDWithoutDot() string {
	return strings.TrimPrefix(e.DomainTLD(), domainSeparator)
}

func (e *bEmailParts) SetUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return ErrInvalidEmailUsernameFormat
	}
	e.username = username
	return nil
}

func (e *bEmailParts) SetDomain(domain string) error {
	if !domainRegex.MatchString(domain) {
		return ErrInvalidEmailDomainFormat
	}
	e.domain = domain
	return nil
}

func (e *bEmailParts) SetDomainName(domainName string) error {
	if !domainNameRegex.MatchString(domainName) {
		return ErrInvalidEmailDomainNameFormat
	}
	e.domain = generateDomain(domainName, e.DomainTLD())
	return nil
}

func (e *bEmailParts) SetDomainTLD(domainTLD string) error {
	if !domainTLDRegex.MatchString(domainTLD) {
		return ErrInvalidEmailDomainTLDFormat
	}
	e.domain = generateDomain(e.DomainName(), domainTLD)
	return nil
}

func (e *bEmailParts) String() string {
	return e.Email()
}

func generateEmail(username, domain string) string {
	return fmt.Sprintf("%s%s%s", username, emailSeparator, domain)
}

func generateDomain(domainName, domainTLD string) string {
	if !strings.HasPrefix(domainTLD, domainSeparator) {
		domainTLD = domainSeparator + domainTLD
	}
	return fmt.Sprintf("%s%s", domainName, domainTLD)
}
