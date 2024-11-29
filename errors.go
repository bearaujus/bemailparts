package bemailparts

import "errors"

var (
	ErrInvalidEmailFormat           = errors.New("invalid email format")
	ErrInvalidEmailUsernameFormat   = errors.New("invalid email username format")
	ErrInvalidEmailDomainFormat     = errors.New("invalid email domain format")
	ErrInvalidEmailDomainNameFormat = errors.New("invalid email domain name format")
	ErrInvalidEmailDomainTLDFormat  = errors.New("invalid email domain tld format")
)
