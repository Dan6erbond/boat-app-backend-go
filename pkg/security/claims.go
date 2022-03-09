package security

import "fmt"

type UserClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (c *UserClaims) GetID() string {
	return fmt.Sprint(c.ID)
}

func (c *UserClaims) GetUsername() string {
	return c.Username
}

type RefreshTokenClaims struct {
	UserClaims
	TokenID uint `json:"token_id"`
}
