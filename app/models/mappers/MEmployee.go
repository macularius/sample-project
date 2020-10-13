package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"

	"github.com/revel/revel"
)

// EmployeeDBType тип сущности "сотрудник" бд
type EmployeeDBType struct {
	Pk_id          int64  // идентификатор
	Fk_position    int64  // FK на должность
	Fk_user        int64  // FK на пользователя
	C_lastname     string // фамилия
	C_firstname    string // имя
	C_middlename   string // отчество
	C_phone_number string // телефонный номер
	C_email        string // почтовый адрес
	C_is_archive   int64  // признак архивности
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *EmployeeDBType) ToType() (e *entities.Employee) {
	e = new(entities.Employee)

	e.ID = dbt.Pk_id
	e.Lastname = dbt.C_lastname
	e.Firstname = dbt.C_firstname
	e.Middlename = dbt.C_middlename
	e.PhoneNumber = dbt.C_phone_number
	e.Email = dbt.C_email

	return
}

// MEmployee маппер сотрудников
type MEmployee struct {
	db *sql.DB
}

// SelectAll получение всех сотрудников
func (m *MEmployee) SelectAll() (es []*EmployeeDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			fk_user,
			fk_position,
			c_name,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email,
			c_is_archive
		FROM "library".t_employees;	
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MEmployee.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		e := new(EmployeeDBType)

		// считывание строки выборки
		err = rows.Scan(&e.Pk_id, &e.Fk_user, &e.Fk_position, &e.C_firstname, &e.C_lastname, &e.C_middlename, &e.C_phone_number, &e.C_email, &e.C_is_archive)
		if err != nil {
			revel.AppLog.Errorf("MEmployee.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		es = append(es, e)
	}

	return
}

// SelectByID получение сотрудника по ID
func (m *MEmployee) SelectByID(id int64) (e *EmployeeDBType, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	e = new(EmployeeDBType)

	// запрос
	query = `
		SELECT
			pk_id,
			fk_user,
			fk_position,
			c_name,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email,
			c_is_archive
		FROM "library".t_employees
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&e.Pk_id, &e.Fk_user, &e.Fk_position, &e.C_firstname, &e.C_lastname, &e.C_middlename, &e.C_phone_number, &e.C_email, &e.C_is_archive)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MEmployee.SelectByID : row.Scan, %s\n", err)
		return
	}

	return
}

// Insert добавление сотрудника
func (m *MEmployee) Insert(employee EmployeeDBType) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		INSERT INTO "library".t_employees(
			fk_user,
			fk_position,
			c_firstname,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email,
			c_is_archive
		)
		VALUES(
			0,	-- fk_user
			0,	-- fk_position
			'',	-- c_firstname
			'',	-- c_lastname
			'',	-- c_middlename
			'',	-- c_phone_number
			'',	-- c_email
			0	-- c_is_archive
		);	
		returning pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query,
		employee.Pk_id,          // pk_id
		employee.Fk_user,        // fk_user
		employee.Fk_position,    // fk_position
		employee.C_firstname,    // c_firstname
		employee.C_lastname,     // c_lastname
		employee.C_middlename,   // c_middlename
		employee.C_phone_number, // c_phone_number
		employee.C_email,        // c_email
		employee.C_is_archive,   // c_is_archive
	)

	// считывание id
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MEmployee.Insert : row.Scan, %s\n", err)
		return
	}

	return
}

// Update изменение сотрудника
func (m *MEmployee) Update(employee EmployeeDBType) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		UPDATE "library".t_employees
			SET 
				fk_user = $2,
				fk_position = $3,
				c_firstname = $4,
				c_lastname = $5,
				c_middlename = $6,
				c_phone_number = $7,
				c_email = $8,
				c_is_archive = $9
		WHERE pk_id = $1;
	`

	// выполнение запроса
	_, err = m.db.Exec(query,
		employee.Pk_id,          // pk_id
		employee.Fk_user,        // fk_user
		employee.Fk_position,    // fk_position
		employee.C_firstname,    // c_firstname
		employee.C_lastname,     // c_lastname
		employee.C_middlename,   // c_middlename
		employee.C_phone_number, // c_phone_number
		employee.C_email,        // c_email
		employee.C_is_archive,   // c_is_archive
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MEmployee.Update : m.db.Exec, %s\n", err)
		return
	}

	return
}

// Delete удаление сотрудника
func (m *MEmployee) Delete(employee EmployeeDBType) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		UPDATE "library".t_employees
		SET c_is_archive = 1
		WHERE pk_id = $1;
	`

	// выполнение запроса
	_, err = m.db.Exec(query, employee.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MEmployee.Delete : m.db.Exec, %s\n", err)
		return
	}

	return
}
