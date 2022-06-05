package svc

import(
  "github.com/golang-jwt/jwt"
)

type TokenService struct {

}

func NewTokenService() *TokenService {
  return &TokenService{}
}

func (tsvc *TokenService) ValidateToken(tknStr string) *jwt.MapClaims {
  return nil
}
