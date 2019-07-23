// Package db ...
package db

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/context"
)

// BaseRepository ...
type BaseRepository struct {
	ctx context.Context
	cn  *gorm.DB
}

// NewBaseRepository ...
func NewBaseRepository(ctx context.Context) BaseRepository {
	return BaseRepository{ctx: ctx}
}

// DB ...
func (b *BaseRepository) DB() (*gorm.DB, error) {
	if b.cn != nil {
		return b.cn, nil
	}

	old := b.ctx.Get("DB")
	if old != nil {
		b.cn = old.(*gorm.DB)
		return b.cn, nil
	}

	cn, err := Setup(b.ctx)
	if err != nil {
		return nil, err
	}
	b.cn = cn
	b.ctx.Set("DB", b.cn)
	return b.cn, err

}

// SetDB ...
func (b *BaseRepository) SetDB(cn *gorm.DB) {
	b.cn = cn
}

// Close ...
func (b *BaseRepository) Close() {
	//TODO do something need to close here please
}
