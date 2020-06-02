# database-vault
```go
package main

import (
	"fmt"

	"github.com/AtlasXV/database-vault/mongo"
)

func main() {
	vault, err := mongo.NewVault("token", "vault_addr")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = vault.GetCreds("db_role")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s,%s,%d", vault.Username, vault.Password, vault.Expire)
}
```
