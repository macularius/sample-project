package mappers

import (
	"database/sql"

	"github.com/revel/revel"
)

// UserDBType тип сущности "пользователь" бд
type UserDBType struct {
	Pk_id  int64  // идентификатор
	C_name string // логин
}

// MUser маппер пользователей
type MUser struct {
	db *sql.DB
}

// SelectUser получение пользователя по логину и паролю
func (m *MUser) SelectUser(login string) (u *UserDBType, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	u = new(UserDBType)

	// запрос
	query = `
		SELECT
			pk_id,
			c_name
		FROM "library".t_users
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, login)

	// считывание строки выборки
	err = row.Scan(&u.Pk_id, &u.C_name)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MUser.SelectUser : row.Scan, %s\n", err)
		return
	}

	return
}

// CheckPassword проверка пароля пользователя
func (m *MUser) CheckPassword(user UserDBType, password string) (f bool, err error) {
	var (
		query string                        // строка запроса
		row   *sql.Row                      // выборка данных
		u     *UserDBType = new(UserDBType) //
	)

	// запрос
	query = `
		SELECT
			pk_id
		FROM "library".t_users
		WHERE
			pk_id = $1 and
			c_login = $2 and
			c_password = $3;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, user.Pk_id, user.C_name, password)

	// считывание строки выборки
	err = row.Scan(&u.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MUser.CheckPassword : row.Scan, %s\n", err)
		return
	}

	return true, nil
}
