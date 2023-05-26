package database

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	TransactionExecutor interface {
		Session(config *gorm.Session) *SimpleTransactionExecutor
		WithContext(ctx context.Context) *SimpleTransactionExecutor
		Debug() (tx *SimpleTransactionExecutor)
		Set(key string, value interface{}) *SimpleTransactionExecutor
		Get(key string) (interface{}, bool)
		InstanceSet(key string, value interface{}) *SimpleTransactionExecutor
		InstanceGet(key string) (interface{}, bool)
		AddError(err error) error
		DB() (*sql.DB, error)
		SetupJoinTable(model interface{}, field string, joinTable interface{}) error
		Use(plugin gorm.Plugin) error
		ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string
		Association(column string) *gorm.Association
		Create(value interface{}) (tx *SimpleTransactionExecutor)
		CreateInBatches(value interface{}, batchSize int) (tx *SimpleTransactionExecutor)
		Save(value interface{}) (tx *SimpleTransactionExecutor)
		First(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		Take(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		Last(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		Find(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *SimpleTransactionExecutor
		FirstOrInit(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		FirstOrCreate(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		Update(column string, value interface{}) (tx *SimpleTransactionExecutor)
		Updates(values interface{}) (tx *SimpleTransactionExecutor)
		UpdateColumn(column string, value interface{}) (tx *SimpleTransactionExecutor)
		UpdateColumns(values interface{}) (tx *SimpleTransactionExecutor)
		Delete(value interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor)
		Count(count *int64) (tx *SimpleTransactionExecutor)
		Row() *sql.Row
		Rows() (*sql.Rows, error)
		Scan(dest interface{}) (tx *SimpleTransactionExecutor)
		Pluck(column string, dest interface{}) (tx *SimpleTransactionExecutor)
		ScanRows(rows *sql.Rows, dest interface{}) error
		Connection(fc func(tx *gorm.DB) error) (err error)
		Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
		Begin(opts ...*sql.TxOptions) *SimpleTransactionExecutor
		Commit() *SimpleTransactionExecutor
		Rollback() *SimpleTransactionExecutor
		SavePoint(name string) *SimpleTransactionExecutor
		RollbackTo(name string) *SimpleTransactionExecutor
		Exec(sql string, values ...interface{}) (tx *SimpleTransactionExecutor)
		Migrator() gorm.Migrator
		AutoMigrate(dst ...interface{}) error
		Model(value interface{}) (tx *SimpleTransactionExecutor)
		Clauses(conds ...clause.Expression) (tx *SimpleTransactionExecutor)
		Table(name string, args ...interface{}) (tx *SimpleTransactionExecutor)
		Distinct(args ...interface{}) (tx *SimpleTransactionExecutor)
		Select(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor)
		Omit(columns ...string) (tx *SimpleTransactionExecutor)
		Where(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor)
		Not(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor)
		Or(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor)
		Joins(query string, args ...interface{}) (tx *SimpleTransactionExecutor)
		InnerJoins(query string, args ...interface{}) (tx *SimpleTransactionExecutor)
		Group(name string) (tx *SimpleTransactionExecutor)
		Having(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor)
		Order(value interface{}) (tx *SimpleTransactionExecutor)
		Limit(limit int) (tx *SimpleTransactionExecutor)
		Offset(offset int) (tx *SimpleTransactionExecutor)
		Scopes(funcs ...func(*gorm.DB) *gorm.DB) (tx *SimpleTransactionExecutor)
		Preload(query string, args ...interface{}) (tx *SimpleTransactionExecutor)
		Attrs(attrs ...interface{}) (tx *SimpleTransactionExecutor)
		Assign(attrs ...interface{}) (tx *SimpleTransactionExecutor)
		Unscoped() (tx *SimpleTransactionExecutor)
		Raw(sql string, values ...interface{}) (tx *SimpleTransactionExecutor)
		Error() error
	}

	SimpleTransactionExecutor struct {
		GDB *gorm.DB
	}
)

func Wrap(db *gorm.DB) *SimpleTransactionExecutor {
	return &SimpleTransactionExecutor{db}
}

func (s *SimpleTransactionExecutor) Session(config *gorm.Session) *SimpleTransactionExecutor {
	return Wrap(s.GDB.Session(config))
}

func (s *SimpleTransactionExecutor) WithContext(ctx context.Context) *SimpleTransactionExecutor {
	return Wrap(s.GDB.WithContext(ctx))
}

func (s *SimpleTransactionExecutor) Debug() (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Debug())
}

func (s *SimpleTransactionExecutor) Set(key string, value interface{}) *SimpleTransactionExecutor {
	return Wrap(s.GDB.Set(key, value))
}

func (s *SimpleTransactionExecutor) Get(key string) (interface{}, bool) {
	return s.GDB.Get(key)
}

func (s *SimpleTransactionExecutor) InstanceSet(key string, value interface{}) *SimpleTransactionExecutor {
	return Wrap(s.GDB.InstanceSet(key, value))
}

func (s *SimpleTransactionExecutor) InstanceGet(key string) (interface{}, bool) {
	return s.GDB.InstanceGet(key)
}

func (s *SimpleTransactionExecutor) AddError(err error) error {
	return s.GDB.AddError(err)
}

func (s *SimpleTransactionExecutor) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return s.GDB.SetupJoinTable(model, field, joinTable)
}

func (s *SimpleTransactionExecutor) Use(plugin gorm.Plugin) error {
	return s.GDB.Use(plugin)
}

func (s *SimpleTransactionExecutor) ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string {
	return s.GDB.ToSQL(queryFn)
}

func (s *SimpleTransactionExecutor) Association(column string) *gorm.Association {
	return s.GDB.Association(column)
}

func (s *SimpleTransactionExecutor) Create(value interface{}) *SimpleTransactionExecutor {
	return Wrap(s.GDB.Create(value))
}

func (s *SimpleTransactionExecutor) CreateInBatches(value interface{}, batchSize int) *SimpleTransactionExecutor {
	return Wrap(s.GDB.CreateInBatches(value, batchSize))
}

func (s *SimpleTransactionExecutor) Save(value interface{}) *SimpleTransactionExecutor {
	return Wrap(s.GDB.Save(value))
}

func (s *SimpleTransactionExecutor) First(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.First(dest, conds...))
}

func (s *SimpleTransactionExecutor) Take(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Take(dest, conds...))
}

func (s *SimpleTransactionExecutor) Last(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Last(dest, conds...))
}

func (s *SimpleTransactionExecutor) Find(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Find(dest, conds...))
}

func (s *SimpleTransactionExecutor) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *SimpleTransactionExecutor {
	return Wrap(s.GDB.FindInBatches(dest, batchSize, fc))
}

func (s *SimpleTransactionExecutor) FirstOrInit(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.FirstOrInit(dest, conds))
}

func (s *SimpleTransactionExecutor) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.FirstOrCreate(dest, conds))
}

func (s *SimpleTransactionExecutor) Update(column string, value interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Update(column, value))
}

func (s *SimpleTransactionExecutor) Updates(values interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.First(values))
}

func (s *SimpleTransactionExecutor) UpdateColumn(column string, value interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.UpdateColumn(column, value))
}

func (s *SimpleTransactionExecutor) UpdateColumns(values interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.UpdateColumns(values))
}

func (s *SimpleTransactionExecutor) Delete(value interface{}, conds ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Delete(value, conds))
}

func (s *SimpleTransactionExecutor) Count(count *int64) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Count(count))
}

func (s *SimpleTransactionExecutor) Row() *sql.Row {
	return s.GDB.Row()
}

func (s *SimpleTransactionExecutor) Rows() (*sql.Rows, error) {
	return s.GDB.Rows()
}

func (s *SimpleTransactionExecutor) Scan(dest interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Scan(dest))
}

func (s *SimpleTransactionExecutor) Pluck(column string, dest interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Pluck(column, dest))
}

func (s *SimpleTransactionExecutor) ScanRows(rows *sql.Rows, dest interface{}) error {
	return s.GDB.ScanRows(rows, dest)
}

func (s *SimpleTransactionExecutor) Connection(fc func(tx *gorm.DB) error) (err error) {
	return s.GDB.Connection(fc)
}

func (s *SimpleTransactionExecutor) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return s.GDB.Connection(fc)
}

func (s *SimpleTransactionExecutor) Begin(opts ...*sql.TxOptions) *SimpleTransactionExecutor {
	return Wrap(s.GDB.Begin(opts...))
}
func (s *SimpleTransactionExecutor) Commit() *SimpleTransactionExecutor {
	return Wrap(s.GDB.Commit())
}
func (s *SimpleTransactionExecutor) Rollback() *SimpleTransactionExecutor {
	return Wrap(s.GDB.Rollback())
}
func (s *SimpleTransactionExecutor) SavePoint(name string) *SimpleTransactionExecutor {
	return Wrap(s.GDB.SavePoint(name))
}
func (s *SimpleTransactionExecutor) RollbackTo(name string) *SimpleTransactionExecutor {
	return Wrap(s.GDB.RollbackTo(name))
}
func (s *SimpleTransactionExecutor) Exec(sql string, values ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Exec(sql, values...))
}
func (s *SimpleTransactionExecutor) Migrator() gorm.Migrator {
	return s.GDB.Migrator()
}
func (s *SimpleTransactionExecutor) AutoMigrate(dst ...interface{}) error {
	return s.GDB.AutoMigrate(dst...)
}
func (s *SimpleTransactionExecutor) Model(value interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Model(value))
}
func (s *SimpleTransactionExecutor) Clauses(conds ...clause.Expression) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Clauses(conds...))
}
func (s *SimpleTransactionExecutor) Table(name string, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Table(name, args...))
}
func (s *SimpleTransactionExecutor) Distinct(args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Distinct(args...))
}
func (s *SimpleTransactionExecutor) Select(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Select(query, args...))
}
func (s *SimpleTransactionExecutor) Omit(columns ...string) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Omit(columns...))
}
func (s *SimpleTransactionExecutor) Where(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Where(query, args...))
}
func (s *SimpleTransactionExecutor) Not(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Not(query, args...))
}
func (s *SimpleTransactionExecutor) Or(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Or(query, args...))
}
func (s *SimpleTransactionExecutor) Joins(query string, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Joins(query, args...))
}
func (s *SimpleTransactionExecutor) InnerJoins(query string, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.InnerJoins(query, args...))
}
func (s *SimpleTransactionExecutor) Group(name string) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Group(name))
}
func (s *SimpleTransactionExecutor) Having(query interface{}, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Having(query, args...))
}
func (s *SimpleTransactionExecutor) Order(value interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Order(value))
}
func (s *SimpleTransactionExecutor) Limit(limit int) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Limit(limit))
}
func (s *SimpleTransactionExecutor) Offset(offset int) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Offset(offset))
}
func (s *SimpleTransactionExecutor) Scopes(funcs ...func(db *gorm.DB) *gorm.DB) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Scopes(funcs...))
}
func (s *SimpleTransactionExecutor) Preload(query string, args ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Preload(query, args))
}
func (s *SimpleTransactionExecutor) Attrs(attrs ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Attrs(attrs...))
}
func (s *SimpleTransactionExecutor) Assign(attrs ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Assign(attrs...))
}
func (s *SimpleTransactionExecutor) Unscoped() (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Unscoped())
}
func (s *SimpleTransactionExecutor) Raw(sql string, values ...interface{}) (tx *SimpleTransactionExecutor) {
	return Wrap(s.GDB.Raw(sql, values...))
}

func (s *SimpleTransactionExecutor) DB() (*sql.DB, error) {
	return s.GDB.DB()
}

func (s *SimpleTransactionExecutor) ExecuteTransaction(command Command) error {
	var err error
	panicHappened := true
	transaction := s.Begin()
	if err = transaction.Error(); err != nil {
		return transaction.Error()
	}

	defer func() {
		if err != nil || panicHappened {
			transaction.Rollback()
		}
	}()

	if err = command(s); err != nil {
		transaction.Rollback()
	}
	panicHappened = false
	return transaction.Commit().Error()
}

func (s *SimpleTransactionExecutor) Error() error {
	return s.GDB.Error
}
