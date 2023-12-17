package authService

import (
	"time"

	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/pkg/token"
)

type expiresToken struct {
	Access  time.Time
	Refresh time.Time
}

func (s *authService) generateToken(claims authDomain.Passport) (*authDomain.ResToken, *expiresToken, error) {
	newToken := token.New(s.cfg, claims)
	accessClaims, accessToken, err := newToken.Signed(token.ACCESS)
	if err != nil {
		return nil, nil, err
	}

	refreshClaims, refreshToken, err := newToken.Signed(token.REFRESH)
	if err != nil {
		return nil, nil, err
	}

	res := &authDomain.ResToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	accessExp, err := time.Parse(time.RFC3339, accessClaims.ExpiresAt.Format(time.RFC3339))
	if err != nil {
		return nil, nil, err
	}

	refreshExp, err := time.Parse(time.RFC3339, refreshClaims.ExpiresAt.Format(time.RFC3339))
	if err != nil {
		return nil, nil, err
	}

	exp := &expiresToken{
		Access:  accessExp,
		Refresh: refreshExp,
	}

	return res, exp, nil
}
