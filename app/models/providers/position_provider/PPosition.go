package position_provider

import (
	"database/sql"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PPosition провайдер контроллера должностей
type PPosition struct {
	positionMapper *mappers.MPosition
}

// Init
func (p *PPosition) Init(db *sql.DB) (err error) {
	// инициализация маппера книг
	p.positionMapper = new(mappers.MPosition)
	p.positionMapper.Init(db)

	return
}

// GetPositions метод получения событий
func (p *PPosition) GetPositions() (ps []*entities.Position, err error) {
	return
}
