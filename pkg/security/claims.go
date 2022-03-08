package security

type UserClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type RefreshTokenClaims struct {
	UserClaims
	TokenID uint `json:"token_id"`
}
