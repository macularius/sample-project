package helpers

import (
	"database/sql"
	"errors"
	"fmt"

	// неиспользуемый импорт пакета драйвера для работы с PostgreSQL
	// при подключении пакета срабатывает функция github.com/lib/pq/conn.go:48,
	// которая регистрирует драйвер "postgres".
	//
	// func init() {
	// 	  sql.Register("postgres", &Driver{})
	// }
	//
	// Далее этот драйвер используется при создании подключения
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

// объявление переменных пакета
var (
	ErrFailedConnection = errors.New("не удалось установить соединение с базой") // Ошибка создания соединения с базой
)

// Паттерн singleton на примере типа dbConnector:
//
// снаружи пакета данный тип представлен интерфейсом IDBConnector
// получение экземпляра dbConnector происходит с помощью функции GetConnector()

// IDBConnector интерфейс типов для взаимодействия с базой
type IDBConnector interface {
	GetDBConnection() (*sql.DB, error)
}

// экземпляр коннектора
var connector *dbConnector

// GetConnector функция получения экземпляра коннектора
func GetConnector() (IDBConnector, error) {
	// создание экземпляра коннектора, если он не создан
	if connector == nil {
		connector = new(dbConnector)
	}

	return connector, nil
}

// тип для работы с бд
type dbConnector struct{}

// GetDBConnection получение экземпляра подключения к базе
func (c *dbConnector) GetDBConnection() (db *sql.DB, err error) {
	var (
		dbspec   string // строка подключения к бд
		user     string // имя пользователя субд
		password string // пароль пользователя субд
		ok       bool   // флаг успешности получения настройки
	)

	// получение строки подключения
	if dbspec, ok = revel.Config.String("db.spec"); !ok {
		err = ErrFailedConnection
		revel.AppLog.Errorf("Не удалось получить строку подключения к бд из файла конфигурации: %v\n", err)
		return
	}
	// получение пользователя бд
	if user, ok = revel.Config.String("db.user"); !ok {
		err = ErrFailedConnection
		revel.AppLog.Errorf("Не удалось получить пользователя из файла конфигурации: %v\n", err)
		return
	}
	// получение пароля пользователя бд
	if password, ok = revel.Config.String("db.password"); !ok {
		err = ErrFailedConnection
		revel.AppLog.Errorf("Не удалось получить пароль из файла конфигурации: %v\n", err)
		return
	}

	// формирование строки подключения к базе
	connStr := fmt.Sprintf("postgres://%s:%s@%s", user, password, dbspec)
	revel.AppLog.Debugf("postgres://%s:%s@%s", user, password, dbspec)
	// создание соединения бд
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		revel.AppLog.Errorf("Не удалось установить соединение с базой данных: %v\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		revel.AppLog.Errorf("Не удалось выполнить тестовое подключение: %v\n", err)
		return
	}

	return
}
