package user_provider

import (
	"database/sql"
	"sample-project/app/models/mappers"
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

// Login
func (p *PUser) Login() (flag bool, err error) {
	return
}

// Logout
func (p *PUser) Logout() (err error) {
	return
}
