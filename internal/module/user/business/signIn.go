package userbusiness

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	usermodel "juliet/internal/module/user/model"
	"juliet/pkg/tokenFactory"
)

type SignInStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.UserModel, error)
}

type signInBiz struct {
	store        SignInStore
	tokenService tokenFactory.TokenFactory
}

func NewSignInStorage(
	store SignInStore,
	tokenService tokenFactory.TokenFactory,
) *signInBiz {
	return &signInBiz{
		store:        store,
		tokenService: tokenService,
	}
}

func (biz *signInBiz) HandleSignIn(
	ctx context.Context,
	payload *usermodel.SignInModel,
) (string, error) {
	var (
		err   error
		token string
		user  *usermodel.UserModel
	)

	user, err = biz.store.FindUser(ctx, map[string]interface{}{
		"email": payload.Email,
	})
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", err
	}

	token, err = biz.tokenService.CreateTokenWithClaims(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
