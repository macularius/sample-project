package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
)

// MUser маппер пользователей
type MUser struct {
	db *sql.DB
}

// SelectUser получение пользователя по логину и паролю
func (m *MUser) SelectUser(login, password string) (u *entities.User, err error) {
	return
}
