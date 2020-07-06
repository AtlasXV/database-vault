package mongo

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

const REFRESH_GAP time.Duration = 36000 * time.Second

type Vault struct {
	Username string
	Password string
	Expire   int64
	client   *api.Client
}

func NewVault(token, vault_addr string) (*Vault, error) {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client.SetToken(token)
	v := &Vault{client: client}
	return v, nil
}

func (v *Vault) GetCreds(role string) error {
	token_secret, token_err := v.client.Auth().Token().LookupSelf()
	if token_err != nil {
		fmt.Println(token_err)
		return token_err
	}
	token_ttl, ttl_err := token_secret.TokenTTL()
	if ttl_err != nil {
		fmt.Println(ttl_err)
		return ttl_err
	}
	if token_ttl <= REFRESH_GAP {
		_, renew_err := v.client.Auth().Token().RenewSelf(2764800)
		if renew_err != nil {
			fmt.Println(renew_err)
			return renew_err
		}
	}
	secret, err := v.client.Logical().Read("database/creds/" + role)
	if err != nil {
		fmt.Println(err)
		return err
	}

	v.Username = secret.Data["username"].(string)
	v.Password = secret.Data["password"].(string)
	v.Expire = time.Now().Unix() + int64(secret.LeaseDuration)
	return nil
}
