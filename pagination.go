package pkg

import (
	"gorm.io/gorm"
)

type Paginator struct {
	DB        *gorm.DB
	Page      int
	Limit     int
	Sort      string
	Total     int64
	Resources interface{}
}

type PaginationParams struct {
	Page  int
	Limit int
	Sort  string
}

func NewPaginator(db *gorm.DB, resources interface{}, params PaginationParams) *Paginator {
	if params.Page == 0 {
		params.Page = 1
	}

	return &Paginator{
		DB:        db,
		Page:      params.Page,
		Limit:     params.Limit,
		Sort:      params.Sort,
		Resources: resources,
	}
}

func (p *Paginator) Paginate() error {
	offset := (p.Page - 1) * p.Limit

	query := p.DB.Model(p.Resources)
	if p.Sort != "" {
		query = query.Order(p.Sort)
	}

	if err := query.Count(&p.Total).Error; err != nil {
		return err
	}

	if err := query.Limit(p.Limit).Offset(offset).Find(p.Resources).Error; err != nil {
		return err
	}

	return nil
}

func (p *Paginator) GetResult() interface{} {
	return p.Resources
}

func (p *Paginator) GetTotalCount() int64 {
	return p.Total
}
