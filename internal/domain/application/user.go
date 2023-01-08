package application

import (
	"context"
	"github.com/ebar-go/ego/component"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"unknown/internal/domain/entity"
	"unknown/internal/domain/repo"
)

type UserApp struct {
	userRepo *repo.UserRepo
	sign     []byte
}

func NewUserApp(userRepo *repo.UserRepo) *UserApp {
	return &UserApp{userRepo: userRepo}
}

func (app *UserApp) Login(ctx context.Context, phone string) (user *entity.UserEntity, err error) {
	user = &entity.UserEntity{Phone: phone}
	err = app.userRepo.Create(ctx, user)
	return
}

type UserClaims struct {
	jwt.RegisteredClaims
	ID    int64  `json:"id"`
	Phone string `json:"phone"`
}

func (app *UserApp) GenToken(ctx context.Context, user *entity.UserEntity) (token string, err error) {
	token, err = component.JWT().GenerateToken(app.sign, &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		ID:    user.ID,
		Phone: user.Phone,
	})
	return
}

func (app *UserApp) Auth(ctx context.Context, token string) (user *entity.UserEntity, err error) {
	claims, err := component.JWT().ParseToken(app.sign, token)
	if err != nil {
		return
	}

	userClaims := claims.(*UserClaims)
	user = &entity.UserEntity{ID: userClaims.ID, Phone: userClaims.Phone}
	return

}
