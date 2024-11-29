package bemailparts_test

import (
	"github.com/bearaujus/bemailparts"
	"testing"
)

func TestBEmailParts(t *testing.T) {
	validator := func(t *testing.T, e bemailparts.BEmailParts, wantEmail, wantUsername,
		wantDomain, wantDomainName, wantDomainTLD, wantDomainTLDWithoutDot string,
	) {
		t.Helper()
		if e.Email() != wantEmail {
			t.Errorf("Email() got = %v, want %v", e.Email(), wantEmail)
			return
		}

		if e.Username() != wantUsername {
			t.Errorf("Username() got = %v, want %v", e.Username(), wantUsername)
			return
		}

		if e.Domain() != wantDomain {
			t.Errorf("Domain() got = %v, want %v", e.Domain(), wantDomain)
			return
		}

		if e.DomainName() != wantDomainName {
			t.Errorf("DomainName() got = %v, want %v", e.DomainName(), wantDomainName)
			return
		}

		if e.DomainTLD() != wantDomainTLD {
			t.Errorf("DomainTLD() got = %v, want %v", e.DomainTLD(), wantDomainTLD)
			return
		}

		if e.DomainTLDWithoutDot() != wantDomainTLDWithoutDot {
			t.Errorf("DomainTLDWithoutDot() got = %v, want %v", e.DomainTLDWithoutDot(), wantDomainTLDWithoutDot)
			return
		}

		if e.String() != wantEmail {
			t.Errorf("String() got = %v, want %v", e.String(), wantEmail)
		}
	}
	t.Run("test bemailparts", func(t *testing.T) {
		e, err := bemailparts.New("test.username@test-domain.com")
		if err != nil {
			t.Fatal(err)
		}
		validator(t, e,
			"test.username@test-domain.com",
			"test.username",
			"test-domain.com",
			"test-domain",
			".com",
			"com",
		)

		err = e.SetUsername("test.update.username")
		if err != nil {
			t.Fatal(err)
		}
		validator(t, e,
			"test.update.username@test-domain.com",
			"test.update.username",
			"test-domain.com",
			"test-domain",
			".com",
			"com",
		)

		err = e.SetDomain("test-update-domain.com")
		if err != nil {
			t.Fatal(err)
		}
		validator(t, e,
			"test.update.username@test-update-domain.com",
			"test.update.username",
			"test-update-domain.com",
			"test-update-domain",
			".com",
			"com",
		)

		err = e.SetDomainName("test-update-domain-name")
		if err != nil {
			t.Fatal(err)
		}
		validator(t, e,
			"test.update.username@test-update-domain-name.com",
			"test.update.username",
			"test-update-domain-name.com",
			"test-update-domain-name",
			".com",
			"com",
		)

		err = e.SetDomainTLD(".co.id")
		if err != nil {
			t.Fatal(err)
		}
		validator(t, e,
			"test.update.username@test-update-domain-name.co.id",
			"test.update.username",
			"test-update-domain-name.co.id",
			"test-update-domain-name",
			".co.id",
			"co.id",
		)

		err = e.SetUsername("!@#@%$!@%")
		if err == nil {
			t.Error("expecting and error on SetUsername() but got nil")
		}

		err = e.SetDomain("!@#@%$!@%")
		if err == nil {
			t.Error("expecting and error on SetDomain() but got nil")
		}

		err = e.SetDomainName("!@#@%$!@%")
		if err == nil {
			t.Error("expecting and error on SetDomainName() but got nil")
		}

		err = e.SetDomainTLD("!@#@%$!@%")
		if err == nil {
			t.Error("expecting and error on SetDomainTLD() but got nil")
		}
	})
}

func TestNew(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{email: "test.username@test-domain.com"},
			want:    "test.username@test-domain.com",
			wantErr: false,
		},
		{
			name:    "error invalid email format",
			args:    args{email: "test.username@test-domain.c123123@#!@#"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bemailparts.New(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestNewFromUsernameAndDomain(t *testing.T) {
	type args struct {
		username string
		domain   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username: "test.username",
				domain:   "test-domain.com",
			},
			want:    "test.username@test-domain.com",
			wantErr: false,
		},
		{
			name: "error invalid email username format",
			args: args{
				username: "!@#!@#!%!@%asd",
				domain:   "test-domain.com",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error invalid email domain format",
			args: args{
				username: "test.username",
				domain:   "test!@%@!-domain.!@#@!$@!",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bemailparts.NewFromUsernameAndDomain(tt.args.username, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromUsernameAndDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("NewFromUsernameAndDomain() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestNewFromFullParts(t *testing.T) {
	type args struct {
		username   string
		domainName string
		domainTLD  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username:   "test.username",
				domainName: "test-domain",
				domainTLD:  "com",
			},
			want:    "test.username@test-domain.com",
			wantErr: false,
		},
		{
			name: "success domain tld with dot",
			args: args{
				username:   "test.username",
				domainName: "test-domain",
				domainTLD:  ".com",
			},
			want:    "test.username@test-domain.com",
			wantErr: false,
		},
		{
			name: "error invalid email domain name format",
			args: args{
				username:   "test.username",
				domainName: "!@#!@%$!@#!@",
				domainTLD:  "com",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error invalid email domain tld format",
			args: args{
				username:   "test.username",
				domainName: "test-domain",
				domainTLD:  "@!#!@%@!%",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bemailparts.NewFromFullParts(tt.args.username, tt.args.domainName, tt.args.domainTLD)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromFullParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("NewFromFullParts() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
