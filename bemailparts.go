package bemailpart

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
	domainTLDPattern  = `[a-zA-Z]+`
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

type bEmailParts struct {
	username string
	domain   string
}

func New(email string) (*bEmailParts, error) {
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

func NewFromUsernameAndDomain(username, domain string) (*bEmailParts, error) {
	if !usernameRegex.MatchString(username) {
		return nil, ErrInvalidEmailUsernameFormat
	}
	if !domainRegex.MatchString(domain) {
		return nil, ErrInvalidEmailDomainFormat
	}
	return New(generateEmail(username, domain))
}

func NewFromFullParts(username, domainName, domainTLD string) (*bEmailParts, error) {
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
	return e.domain[:strings.LastIndex(e.domain, domainSeparator)]
}

func (e *bEmailParts) DomainTLD() string {
	return e.domain[strings.LastIndex(e.domain, domainSeparator):]
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

// String reconstructs the email address from the EmailPart components.
func (e *bEmailParts) String() string {
	return e.Email()
}

func generateEmail(username, domain string) string {
	return fmt.Sprintf("%s@%s", username, domain)
}

func generateDomain(domainName, domainTLD string) string {
	return fmt.Sprintf("%s.%s", domainName, domainTLD)
}
