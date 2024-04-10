package repoimpl

import (
	"app02/db"
	"app02/log"
	"app02/models"
	"app02/models/req"
	"app02/reponsitory"
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) reponsitory.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) CreateUser(context context.Context, user models.User) (models.User, error) {
	statement := `
		INSERT INTO users(id, name, email, password, created_at, updated_at)
		Values(:id, :name, :email, :password, :created_at, :updated_at)
	`
	user.Created_At = time.Now()
	user.Updated_At = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqIn) (models.User, error) {
	var user = models.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
