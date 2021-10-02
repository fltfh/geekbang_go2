package repo

import (
	"context"
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"template/internal/domain"
	"template/internal/biz"
)

type UserRepo struct {
	db *gorm.DB
	logger *logrus.Logger
}

func NewUserRepo (db *gorm.DB, logger *logrus.Logger) biz.UserRepo{
	return &UserRepo{
		db: db,
		logger: logger,
	}
}

func (u *UserRepo) GetById(ctx context.Context, id int64) (*domain.User, error){
	tx := u.db.BeginTx(ctx, &sql.TxOptions{})
	user := &domain.User{}
	tx.First(user, id)
	return user, nil
}
