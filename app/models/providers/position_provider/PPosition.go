package position_provider

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PPosition провайдер контроллера должностей
type PPosition struct {
	positionMapper *mappers.MPosition
}

// GetPositions метод получения событий
func (p *PPosition) GetPositions() (ps []*entities.Position, err error) {
	return
}
