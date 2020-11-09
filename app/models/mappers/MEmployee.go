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
	C_lastname     string // фамилия
	C_firstname    string // имя
	C_middlename   string // отчество
	C_phone_number string // телефонный номер
	C_email        string // почтовый адрес
	C_is_archive   int64  // признак архивности
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *EmployeeDBType) ToType() (e *entities.Employee, err error) {
	e = new(entities.Employee)

	e.ID = dbt.Pk_id
	e.Lastname = dbt.C_lastname
	e.Firstname = dbt.C_firstname
	e.Middlename = dbt.C_middlename
	e.PhoneNumber = dbt.C_phone_number
	e.Email = dbt.C_email

	return
}

// FromType функция преобразования типа сущности к типу бд
// допускается, что dbt is nil
func (_ *EmployeeDBType) FromType(e entities.Employee) (dbt *EmployeeDBType, err error) {
	dbt = &EmployeeDBType{
		Pk_id:          e.ID,
		C_lastname:     e.Lastname,
		C_firstname:    e.Firstname,
		C_middlename:   e.Middlename,
		C_phone_number: e.PhoneNumber,
		C_email:        e.Email,
	}

	return
}

// MEmployee маппер сотрудников
type MEmployee struct {
	db *sql.DB
}

// Init
func (m *MEmployee) Init(db *sql.DB) {
	m.db = db
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
			fk_position,
			c_firstname,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email
		FROM "library".t_employees
		WHERE c_is_archive = 0
		ORDER BY pk_id;
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
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
		err = rows.Scan(
			&e.Pk_id,          // pk_id
			&e.Fk_position,    // fk_position
			&e.C_firstname,    // c_firstname
			&e.C_lastname,     // c_lastname
			&e.C_middlename,   // c_middlename
			&e.C_phone_number, // c_phone_number
			&e.C_email,        // c_email
		)
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
			fk_position,
			c_firstname,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email
		FROM "library".t_employees
		WHERE pk_id = $1 and
			c_is_archive = 0
		ORDER BY pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(
		&e.Pk_id,          // pk_id
		&e.Fk_position,    // fk_position
		&e.C_firstname,    // c_firstname
		&e.C_lastname,     // c_lastname
		&e.C_middlename,   // c_middlename
		&e.C_phone_number, // c_phone_number
		&e.C_email,        // c_email
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEmployee.SelectByID : row.Scan, %s\n", err)
		return
	}

	return
}

// Insert добавление сотрудника
func (m *MEmployee) Insert(edbt *EmployeeDBType) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		INSERT INTO "library".t_employees(
			fk_position,
			c_firstname,
			c_lastname,
			c_middlename,
			c_phone_number,
			c_email,
			c_is_archive
		)
		VALUES(
			$1,	-- fk_position
			$2,	-- c_firstname
			$3,	-- c_lastname
			$4,	-- c_middlename
			$5,	-- c_phone_number
			$6,	-- c_email
			0	-- c_is_archive
		)
		returning pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query,
		edbt.Fk_position,    // fk_position
		edbt.C_firstname,    // c_firstname
		edbt.C_lastname,     // c_lastname
		edbt.C_middlename,   // c_middlename
		edbt.C_phone_number, // c_phone_number
		edbt.C_email,        // c_email
	)

	// считывание id
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEmployee.Insert : row.Scan, %s\n", err)
		return
	}

	return
}

// Update изменение сотрудника
func (m *MEmployee) Update(edbt *EmployeeDBType) (err error) {
	var (
		query string // строка запроса
	)

	revel.AppLog.Debugf("MEmployee.Update, edbt: %+v\n", edbt)

	// запрос
	query = `
		UPDATE "library".t_employees
		SET 
			fk_position = $2,
			c_firstname = $3,
			c_lastname = $4,
			c_middlename = $5,
			c_phone_number = $6,
			c_email = $7
		WHERE pk_id = $1;
	`

	// выполнение запроса
	_, err = m.db.Exec(query,
		edbt.Pk_id,          // pk_id
		edbt.Fk_position,    // fk_position
		edbt.C_firstname,    // c_firstname
		edbt.C_lastname,     // c_lastname
		edbt.C_middlename,   // c_middlename
		edbt.C_phone_number, // c_phone_number
		edbt.C_email,        // c_email
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEmployee.Update : m.db.Exec, %s\n", err)
		return
	}

	return
}

// Delete удаление сотрудника
func (m *MEmployee) Delete(edbt *EmployeeDBType) (err error) {
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
	_, err = m.db.Exec(query, edbt.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEmployee.Delete : m.db.Exec, %s\n", err)
		return
	}

	return
}
