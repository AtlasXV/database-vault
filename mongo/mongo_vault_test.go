package mongo

import (
	"testing"
	"time"
)

func TestVault_GetCreds(t *testing.T) {
	type args struct {
		token      string
		vault_addr string
		role       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"refresh expired token", args{"s.Q1ufN8z2QbhLy0xGGnNKzxp7", "https://vault.etm.tech", "robinhood"}, true},
		{"refresh token", args{"s.EM8qEQ5GibI2pH5jTKVF86ti", "https://vault.etm.tech", "robinhood"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, _ := NewVault(tt.args.token, tt.args.vault_addr)
			err := v.GetCreds(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vault.GetCreds() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if v.Expire < time.Now().Unix() {
					t.Errorf("Vault.GetCreds() error = %v", "cred has expired")
				}
				if v.Username == "" || v.Password == "" {
					t.Errorf("Vault.GetCreds() error = %v", "get cred error")
				}
			}
		})
	}
}
