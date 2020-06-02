# database-vault
```go
package main

import (
	"fmt"

	"github.com/AtlasXV/database-vault/mongo"
)

func main() {
	// 创建一个新的vault
	vault, err := mongo.NewVault("token", "vault_addr")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 根据需求获取对应db的角色的认证信息
	err = vault.GetCreds("db_role")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s,%s,%d", vault.Username, vault.Password, vault.Expire)
	// vault.Username db用户名
	// vault.Password db密码
	// vault.Expire   认证过期时间，请在此时间之前重新请求新的认证信息
}
```
