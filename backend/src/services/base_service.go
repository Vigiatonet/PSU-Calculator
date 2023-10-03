package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/common"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"

	"gorm.io/gorm"
)

// FIXME: if these generic functions break there is no way i can debug this. (But is saves time)
// There is nothing we can do -Napoleon

type Preloads struct {
	name string
}

type BaseService[T, Tu, Tc, Tr any] struct {
	DB      *gorm.DB
	Preload []Preloads
	Logger  logging.Logger
}

func NewBaseService[T, Tu, Tc, Tr any]() *BaseService[T, Tu, Tc, Tr] {
	db := db.GetDB()
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	return &BaseService[T, Tu, Tc, Tr]{
		DB:      db,
		Logger:  logger,
		Preload: []Preloads{},
	}
}

func LoadPreloads(db *gorm.DB, preloads []Preloads) *gorm.DB {
	for _, p := range preloads {
		err := db.Preload(p.name).Error
		if err != nil {
			panic(err)
		} else {
			db = db.Preload(p.name)
		}
	}
	return db
}

func (b *BaseService[T, Tu, Tc, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {

	model := new(T)
	db := LoadPreloads(b.DB, b.Preload)
	err := db.Model(&model).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	res, err := common.TypeConverter[Tr](model)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *BaseService[T, Tu, Tc, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	db := b.DB.WithContext(ctx).Begin() // atomic transaction so we dont have corrupted models in db
	model, err := common.TypeConverter[T](req)
	if err != nil {
		return nil, err
	}
	err = db.Create(&model).Error
	if err != nil {
		db.Rollback()
		b.Logger.Error(err, logging.Postgres, logging.Insert, "cant create model", nil)
		return nil, err
	}

	baseModel, err := common.TypeConverter[models.BaseModel](model) // convert to base so we get access to Id field
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return b.GetById(ctx, baseModel.ID)
}

func (b *BaseService[T, Tu, Tc, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ConvertToSnakeCase(k)] = v // gorm only accept snake case args like postgres so i have to convert then to snake_case
	}
	snakeMap["updated_at"] = &sql.NullTime{Valid: true, Time: time.Now()}
	snakeMap["updated_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}
	model := new(T)
	db := b.DB.WithContext(ctx).Begin()
	err = db.Model(&model).Where("id = ?", id).Updates(snakeMap).Error
	if err != nil {
		db.Rollback()
		b.Logger.Error(err, logging.Postgres, logging.Update, "cant Update model", nil)
		return nil, err
	}
	db.Commit()
	return b.GetById(ctx, id)
}

func (b *BaseService[T, Tu, Tc, Tr]) Delete(ctx context.Context, id int) error {
	model := new(T)
	db := b.DB.WithContext(ctx).Begin()
	err := db.First(&model, id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&model).Error
	if err != nil {
		db.Rollback()
		b.Logger.Error(err, logging.Postgres, logging.Delete, "cant delete model", nil)
		return err
	}
	db.Commit()
	return nil
}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(strings.Title(name))
			if ok { // Adding db filters and connect them with an AND in the end
				fld.Name = common.ConvertToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEquals":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}

				}
			}
		}
	}
	return strings.Join(query, " AND ")
}

// getSort
func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(strings.Title(tp.ColId))
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ConvertToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

func NewPageList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PageList[T] {
	pl := &dto.PageList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPervious = pl.PageNumber > 1
	return pl

}

func Paginate[T, Tr any](pagination *dto.PaginationInputWithFilter, preloads []Preloads, db *gorm.DB) (*dto.PageList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr
	db = LoadPreloads(db, preloads)
	q := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)
	var total_rows int64

	err := db.Model(&model).Where(q).Count(&total_rows).Error
	if err != nil {
		return nil, err
	}
	err = db.Where(q).Offset(pagination.GetOffSet()).Limit(pagination.GetPageSize()).Order(sort).Find(&items).Error
	if err != nil {
		return nil, err
	}
	rItems, err = common.TypeConverter[[]Tr](items)
	if err != nil {
		return nil, err
	}
	return NewPageList[Tr](rItems, total_rows, pagination.PageNumber, int64(pagination.GetPageSize())), nil

}

func (bs *BaseService[T, Tu, Tc, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[Tr], error) {
	return Paginate[T, Tr](req, bs.Preload, bs.DB)
}
