package user_provider

import "sample-project/app/models/mappers"

// PUser провайдер контроллера пользователей
type PUser struct {
	userMapper *mappers.MUser
}

// Login
func (p *PUser) Login() (flag bool, err error) {
	return
}

// Logout
func (p *PUser) Logout() (err error) {
	return
}
