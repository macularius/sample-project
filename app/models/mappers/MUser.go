package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"

	"github.com/revel/revel"
)

// UserDBType тип сущности "пользователь" бд
type UserDBType struct {
	Pk_id   int64  // идентификатор
	C_login string // логин
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *UserDBType) ToType() (u *entities.User, err error) {
	u = new(entities.User)

	u.ID = dbt.Pk_id
	u.Login = dbt.C_login

	return
}

// FromType функция преобразования типа сущности к типу бд
// допускается, что dbt is nil
func (dbt *UserDBType) FromType(u entities.User) (err error) {
	dbt = &UserDBType{
		Pk_id:   u.ID,
		C_login: u.Login,
	}

	return
}

// MUser маппер пользователей
type MUser struct {
	db *sql.DB
}

// Init
func (m *MUser) Init(db *sql.DB) {
	m.db = db
}

// Inser добавление пользователя
func (m *MUser) Insert(u *UserDBType, password string) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		INSERT INTO "library".t_users(
			c_login,
			c_password
		)
		VALUES(
			$1,
			$2
		);	
	`

	// выполнение запроса
	row = m.db.QueryRow(query, u.Pk_id, password)

	// считывание строки выборки
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MUser.Insert : row.Scan, %s\n", err)
		return
	}

	return
}

// SelectUserByID получение пользователя по id
func (m *MUser) SelectUserByID(id int64) (u *UserDBType, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	u = new(UserDBType)

	// запрос
	query = `
		SELECT
			pk_id,
			c_login
		FROM "library".t_users
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&u.Pk_id, &u.C_login)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MUser.SelectUser : row.Scan, %s\n", err)
		return
	}

	return
}

// SelectUserByLogin получение пользователя по логину
func (m *MUser) SelectUserByLogin(login string) (u *UserDBType, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	u = new(UserDBType)

	// запрос
	query = `
		SELECT
			pk_id,
			c_login
		FROM "library".t_users
		WHERE c_login = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, login)

	// считывание строки выборки
	err = row.Scan(&u.Pk_id, &u.C_login)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MUser.SelectUser : row.Scan, %s\n", err)
		return
	}

	return
}

// CheckPassword проверка пароля пользователя
func (m *MUser) CheckPassword(user *UserDBType, password string) (f bool, err error) {
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
	row = m.db.QueryRow(query, user.Pk_id, user.C_login, password)

	// считывание строки выборки
	err = row.Scan(&u.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MUser.CheckPassword : row.Scan, %s\n", err)
		return
	}

	return true, nil
}
