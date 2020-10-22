package user_provider

import (
	"crypto/md5"
	"database/sql"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"

	"github.com/revel/revel"
)

// PUser провайдер контроллера пользователей
type PUser struct {
	userMapper     *mappers.MUser
	employeeMapper *mappers.MEmployee
}

// Init
func (p *PUser) Init() (err error) {
	var db *sql.DB // экземпляр подключения к бд

	// получение экземпляра подключения к бд
	db, err = helpers.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("PUser.Init : helpers.GetDBConnection, %s\n", err)
		return err
	}
	// инициализация маппера пользователей
	p.userMapper = new(mappers.MUser)
	p.userMapper.Init(db)

	// инициализация маппера сотрудников
	p.employeeMapper = new(mappers.MEmployee)
	p.employeeMapper.Init(db)

	return
}

// Validate метод
func (p *PUser) Validate(user *entities.User) (flag bool, err error) {
	var (
		udbt *mappers.UserDBType
	)

	// проверка существования пользователя
	udbt, err = p.userMapper.SelectUserByLogin(user.Login)
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.userMapper.SelectUserByLogin, %s\n", err)
		return
	}
	if udbt == nil {
		revel.AppLog.Debugf("PUser.Validate : p.userMapper.SelectUserByLogin, udbt not found\n")
		return false, nil
	}

	// шифрация пароля
	password := md5.Sum([]byte(user.Password))

	// проверка пароля пользователя
	flag, err = p.userMapper.CheckPassword(udbt, password[:])
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.userMapper.CheckPassword, %s\n", err)
		return
	}

	return
}

// AttachEmployee метод
func (p *PUser) AttachEmployee(user *entities.User) (err error) {
	var (
		udbt *mappers.UserDBType
		edbt *mappers.EmployeeDBType
	)

	// проверка существования пользователя
	udbt, err = p.userMapper.SelectUserByLogin(user.Login)
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.userMapper.SelectUserByLogin, %s\n", err)
		return
	}

	// получение сотрудника
	edbt, err = p.employeeMapper.SelectByID(udbt.Fk_employee)
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.employeeMapper.SelectByID, %s\n", err)
		return
	}

	// преобразования типа бд к типу сущности
	user.Employee, err = edbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PUser.Validate : p.employeeMapper.SelectByID, %s\n", err)
		return
	}

	return
}
