package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"

	"github.com/revel/revel"
)

// PositionDBType тип сущности "событие" бд
type PositionDBType struct {
	Pk_id  int64  // идентификатор
	C_name string // название должности
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *PositionDBType) ToType() (p *entities.Position) {
	p = new(entities.Position)

	p.ID = dbt.Pk_id
	p.Name = dbt.C_name

	return
}

// MPosition маппер должностей
type MPosition struct {
	db *sql.DB
}

// SelectAll получение всех должностей
func (m *MPosition) SelectAll() (pdbts []*PositionDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			c_name
		FROM "library".ref_position;	
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MPosition.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		pdbt := new(PositionDBType)

		// считывание строки выборки
		err = rows.Scan(&pdbt.Pk_id, &pdbt.C_name)
		if err != nil {
			revel.AppLog.Errorf("MPosition.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		pdbts = append(pdbts, pdbt)
	}

	return
}
