package user_provider

import (
	"database/sql"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"

	"github.com/revel/revel"
)

// PUser провайдер контроллера пользователей
type PUser struct {
	userMapper *mappers.MUser
}

// Init
func (p *PUser) Init(db *sql.DB) (err error) {
	// инициализация маппера книг
	p.userMapper = new(mappers.MUser)
	p.userMapper.Init(db)

	return
}

// Validate метод
func (p *PUser) Validate(user *entities.User, password string) (flag bool, err error) {
	var (
		udbt *mappers.UserDBType
	)

	// проверка существования пользователя
	udbt, err = p.userMapper.SelectUserByLogin(user.Login)
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.userMapper.SelectUserByLogin, %s\n", err)
		return
	}

	flag, err = p.userMapper.CheckPassword(udbt, password)
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.userMapper.CheckPassword, %s\n", err)
		return
	}

	return
}
