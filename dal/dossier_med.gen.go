// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"patient/model"
)

func newDossierMed(db *gorm.DB, opts ...gen.DOOption) dossierMed {
	_dossierMed := dossierMed{}

	_dossierMed.dossierMedDo.UseDB(db, opts...)
	_dossierMed.dossierMedDo.UseModel(&model.DossierMed{})

	tableName := _dossierMed.dossierMedDo.TableName()
	_dossierMed.ALL = field.NewAsterisk(tableName)
	_dossierMed.DateDiagnostic = field.NewTime(tableName, "date_diagnostic")
	_dossierMed.NNiss = field.NewString(tableName, "n_niss")
	_dossierMed.NomPathologie = field.NewString(tableName, "nom_pathologie")

	_dossierMed.fillFieldMap()

	return _dossierMed
}

type dossierMed struct {
	dossierMedDo dossierMedDo

	ALL            field.Asterisk
	DateDiagnostic field.Time
	NNiss          field.String
	NomPathologie  field.String

	fieldMap map[string]field.Expr
}

func (d dossierMed) Table(newTableName string) *dossierMed {
	d.dossierMedDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dossierMed) As(alias string) *dossierMed {
	d.dossierMedDo.DO = *(d.dossierMedDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dossierMed) updateTableName(table string) *dossierMed {
	d.ALL = field.NewAsterisk(table)
	d.DateDiagnostic = field.NewTime(table, "date_diagnostic")
	d.NNiss = field.NewString(table, "n_niss")
	d.NomPathologie = field.NewString(table, "nom_pathologie")

	d.fillFieldMap()

	return d
}

func (d *dossierMed) WithContext(ctx context.Context) *dossierMedDo {
	return d.dossierMedDo.WithContext(ctx)
}

func (d dossierMed) TableName() string { return d.dossierMedDo.TableName() }

func (d dossierMed) Alias() string { return d.dossierMedDo.Alias() }

func (d *dossierMed) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dossierMed) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["date_diagnostic"] = d.DateDiagnostic
	d.fieldMap["n_niss"] = d.NNiss
	d.fieldMap["nom_pathologie"] = d.NomPathologie
}

func (d dossierMed) clone(db *gorm.DB) dossierMed {
	d.dossierMedDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dossierMed) replaceDB(db *gorm.DB) dossierMed {
	d.dossierMedDo.ReplaceDB(db)
	return d
}

type dossierMedDo struct{ gen.DO }

func (d dossierMedDo) Debug() *dossierMedDo {
	return d.withDO(d.DO.Debug())
}

func (d dossierMedDo) WithContext(ctx context.Context) *dossierMedDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dossierMedDo) ReadDB() *dossierMedDo {
	return d.Clauses(dbresolver.Read)
}

func (d dossierMedDo) WriteDB() *dossierMedDo {
	return d.Clauses(dbresolver.Write)
}

func (d dossierMedDo) Session(config *gorm.Session) *dossierMedDo {
	return d.withDO(d.DO.Session(config))
}

func (d dossierMedDo) Clauses(conds ...clause.Expression) *dossierMedDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dossierMedDo) Returning(value interface{}, columns ...string) *dossierMedDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dossierMedDo) Not(conds ...gen.Condition) *dossierMedDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dossierMedDo) Or(conds ...gen.Condition) *dossierMedDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dossierMedDo) Select(conds ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dossierMedDo) Where(conds ...gen.Condition) *dossierMedDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dossierMedDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *dossierMedDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dossierMedDo) Order(conds ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dossierMedDo) Distinct(cols ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dossierMedDo) Omit(cols ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dossierMedDo) Join(table schema.Tabler, on ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dossierMedDo) LeftJoin(table schema.Tabler, on ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dossierMedDo) RightJoin(table schema.Tabler, on ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dossierMedDo) Group(cols ...field.Expr) *dossierMedDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dossierMedDo) Having(conds ...gen.Condition) *dossierMedDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dossierMedDo) Limit(limit int) *dossierMedDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dossierMedDo) Offset(offset int) *dossierMedDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dossierMedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *dossierMedDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dossierMedDo) Unscoped() *dossierMedDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dossierMedDo) Create(values ...*model.DossierMed) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dossierMedDo) CreateInBatches(values []*model.DossierMed, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dossierMedDo) Save(values ...*model.DossierMed) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dossierMedDo) First() (*model.DossierMed, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DossierMed), nil
	}
}

func (d dossierMedDo) Take() (*model.DossierMed, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DossierMed), nil
	}
}

func (d dossierMedDo) Last() (*model.DossierMed, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DossierMed), nil
	}
}

func (d dossierMedDo) Find() ([]*model.DossierMed, error) {
	result, err := d.DO.Find()
	return result.([]*model.DossierMed), err
}

func (d dossierMedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DossierMed, err error) {
	buf := make([]*model.DossierMed, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dossierMedDo) FindInBatches(result *[]*model.DossierMed, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dossierMedDo) Attrs(attrs ...field.AssignExpr) *dossierMedDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dossierMedDo) Assign(attrs ...field.AssignExpr) *dossierMedDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dossierMedDo) Joins(fields ...field.RelationField) *dossierMedDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dossierMedDo) Preload(fields ...field.RelationField) *dossierMedDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dossierMedDo) FirstOrInit() (*model.DossierMed, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DossierMed), nil
	}
}

func (d dossierMedDo) FirstOrCreate() (*model.DossierMed, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DossierMed), nil
	}
}

func (d dossierMedDo) FindByPage(offset int, limit int) (result []*model.DossierMed, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dossierMedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dossierMedDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dossierMedDo) Delete(models ...*model.DossierMed) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dossierMedDo) withDO(do gen.Dao) *dossierMedDo {
	d.DO = *do.(*gen.DO)
	return d
}
