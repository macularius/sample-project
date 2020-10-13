package position_provider

import (
	"database/sql"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"

	"github.com/revel/revel"
)

// PPosition провайдер контроллера должностей
type PPosition struct {
	positionMapper *mappers.MPosition
}

// Init
func (p *PPosition) Init(db *sql.DB) (err error) {
	// инициализация маппера должностей
	p.positionMapper = new(mappers.MPosition)
	p.positionMapper.Init(db)

	return
}

// GetPositions метод получения должностей
func (p *PPosition) GetPositions() (ps []*entities.Position, err error) {
	var (
		pdbts []*mappers.PositionDBType
		pos   *entities.Position
	)

	// получение данных должностей
	pdbts, err = p.positionMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PPosition.GetPositions : p.positionMapper.SelectAll, %s\n", err)
		return
	}

	for _, pdbt := range pdbts {
		// преобразование к типу сущности
		pos, err = pdbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PPosition.GetPositions : pdbt.ToType, %s\n", err)
			return
		}

		ps = append(ps, pos)
	}

	return
}
