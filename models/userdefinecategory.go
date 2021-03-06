// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Userdefinecategory is an object representing the database table.
type Userdefinecategory struct {
	Index   int    `boil:"index" json:"index" toml:"index" yaml:"index"`
	Blogger int    `boil:"blogger" json:"blogger" toml:"blogger" yaml:"blogger"`
	Cate    string `boil:"cate" json:"cate" toml:"cate" yaml:"cate"`
	Files   int    `boil:"files" json:"files" toml:"files" yaml:"files"`

	R *userdefinecategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userdefinecategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserdefinecategoryColumns = struct {
	Index   string
	Blogger string
	Cate    string
	Files   string
}{
	Index:   "index",
	Blogger: "blogger",
	Cate:    "cate",
	Files:   "files",
}

// Generated where

var UserdefinecategoryWhere = struct {
	Index   whereHelperint
	Blogger whereHelperint
	Cate    whereHelperstring
	Files   whereHelperint
}{
	Index:   whereHelperint{field: "`userdefinecategory`.`index`"},
	Blogger: whereHelperint{field: "`userdefinecategory`.`blogger`"},
	Cate:    whereHelperstring{field: "`userdefinecategory`.`cate`"},
	Files:   whereHelperint{field: "`userdefinecategory`.`files`"},
}

// UserdefinecategoryRels is where relationship names are stored.
var UserdefinecategoryRels = struct {
}{}

// userdefinecategoryR is where relationships are stored.
type userdefinecategoryR struct {
}

// NewStruct creates a new relationship struct
func (*userdefinecategoryR) NewStruct() *userdefinecategoryR {
	return &userdefinecategoryR{}
}

// userdefinecategoryL is where Load methods for each relationship are stored.
type userdefinecategoryL struct{}

var (
	userdefinecategoryAllColumns            = []string{"index", "blogger", "cate", "files"}
	userdefinecategoryColumnsWithoutDefault = []string{"blogger", "cate"}
	userdefinecategoryColumnsWithDefault    = []string{"index", "files"}
	userdefinecategoryPrimaryKeyColumns     = []string{"index"}
)

type (
	// UserdefinecategorySlice is an alias for a slice of pointers to Userdefinecategory.
	// This should generally be used opposed to []Userdefinecategory.
	UserdefinecategorySlice []*Userdefinecategory
	// UserdefinecategoryHook is the signature for custom Userdefinecategory hook methods
	UserdefinecategoryHook func(boil.Executor, *Userdefinecategory) error

	userdefinecategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userdefinecategoryType                 = reflect.TypeOf(&Userdefinecategory{})
	userdefinecategoryMapping              = queries.MakeStructMapping(userdefinecategoryType)
	userdefinecategoryPrimaryKeyMapping, _ = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, userdefinecategoryPrimaryKeyColumns)
	userdefinecategoryInsertCacheMut       sync.RWMutex
	userdefinecategoryInsertCache          = make(map[string]insertCache)
	userdefinecategoryUpdateCacheMut       sync.RWMutex
	userdefinecategoryUpdateCache          = make(map[string]updateCache)
	userdefinecategoryUpsertCacheMut       sync.RWMutex
	userdefinecategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userdefinecategoryBeforeInsertHooks []UserdefinecategoryHook
var userdefinecategoryBeforeUpdateHooks []UserdefinecategoryHook
var userdefinecategoryBeforeDeleteHooks []UserdefinecategoryHook
var userdefinecategoryBeforeUpsertHooks []UserdefinecategoryHook

var userdefinecategoryAfterInsertHooks []UserdefinecategoryHook
var userdefinecategoryAfterSelectHooks []UserdefinecategoryHook
var userdefinecategoryAfterUpdateHooks []UserdefinecategoryHook
var userdefinecategoryAfterDeleteHooks []UserdefinecategoryHook
var userdefinecategoryAfterUpsertHooks []UserdefinecategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Userdefinecategory) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Userdefinecategory) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Userdefinecategory) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Userdefinecategory) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Userdefinecategory) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Userdefinecategory) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Userdefinecategory) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Userdefinecategory) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Userdefinecategory) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userdefinecategoryAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserdefinecategoryHook registers your hook function for all future operations.
func AddUserdefinecategoryHook(hookPoint boil.HookPoint, userdefinecategoryHook UserdefinecategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userdefinecategoryBeforeInsertHooks = append(userdefinecategoryBeforeInsertHooks, userdefinecategoryHook)
	case boil.BeforeUpdateHook:
		userdefinecategoryBeforeUpdateHooks = append(userdefinecategoryBeforeUpdateHooks, userdefinecategoryHook)
	case boil.BeforeDeleteHook:
		userdefinecategoryBeforeDeleteHooks = append(userdefinecategoryBeforeDeleteHooks, userdefinecategoryHook)
	case boil.BeforeUpsertHook:
		userdefinecategoryBeforeUpsertHooks = append(userdefinecategoryBeforeUpsertHooks, userdefinecategoryHook)
	case boil.AfterInsertHook:
		userdefinecategoryAfterInsertHooks = append(userdefinecategoryAfterInsertHooks, userdefinecategoryHook)
	case boil.AfterSelectHook:
		userdefinecategoryAfterSelectHooks = append(userdefinecategoryAfterSelectHooks, userdefinecategoryHook)
	case boil.AfterUpdateHook:
		userdefinecategoryAfterUpdateHooks = append(userdefinecategoryAfterUpdateHooks, userdefinecategoryHook)
	case boil.AfterDeleteHook:
		userdefinecategoryAfterDeleteHooks = append(userdefinecategoryAfterDeleteHooks, userdefinecategoryHook)
	case boil.AfterUpsertHook:
		userdefinecategoryAfterUpsertHooks = append(userdefinecategoryAfterUpsertHooks, userdefinecategoryHook)
	}
}

// OneG returns a single userdefinecategory record from the query using the global executor.
func (q userdefinecategoryQuery) OneG() (*Userdefinecategory, error) {
	return q.One(boil.GetDB())
}

// One returns a single userdefinecategory record from the query.
func (q userdefinecategoryQuery) One(exec boil.Executor) (*Userdefinecategory, error) {
	o := &Userdefinecategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for userdefinecategory")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Userdefinecategory records from the query using the global executor.
func (q userdefinecategoryQuery) AllG() (UserdefinecategorySlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Userdefinecategory records from the query.
func (q userdefinecategoryQuery) All(exec boil.Executor) (UserdefinecategorySlice, error) {
	var o []*Userdefinecategory

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Userdefinecategory slice")
	}

	if len(userdefinecategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Userdefinecategory records in the query, and panics on error.
func (q userdefinecategoryQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Userdefinecategory records in the query.
func (q userdefinecategoryQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count userdefinecategory rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userdefinecategoryQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q userdefinecategoryQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if userdefinecategory exists")
	}

	return count > 0, nil
}

// Userdefinecategories retrieves all the records using an executor.
func Userdefinecategories(mods ...qm.QueryMod) userdefinecategoryQuery {
	mods = append(mods, qm.From("`userdefinecategory`"))
	return userdefinecategoryQuery{NewQuery(mods...)}
}

// FindUserdefinecategoryG retrieves a single record by ID.
func FindUserdefinecategoryG(index int, selectCols ...string) (*Userdefinecategory, error) {
	return FindUserdefinecategory(boil.GetDB(), index, selectCols...)
}

// FindUserdefinecategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserdefinecategory(exec boil.Executor, index int, selectCols ...string) (*Userdefinecategory, error) {
	userdefinecategoryObj := &Userdefinecategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `userdefinecategory` where `index`=?", sel,
	)

	q := queries.Raw(query, index)

	err := q.Bind(nil, exec, userdefinecategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from userdefinecategory")
	}

	return userdefinecategoryObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Userdefinecategory) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Userdefinecategory) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no userdefinecategory provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userdefinecategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userdefinecategoryInsertCacheMut.RLock()
	cache, cached := userdefinecategoryInsertCache[key]
	userdefinecategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userdefinecategoryAllColumns,
			userdefinecategoryColumnsWithDefault,
			userdefinecategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `userdefinecategory` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `userdefinecategory` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `userdefinecategory` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userdefinecategoryPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into userdefinecategory")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userdefinecategoryMapping["index"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Index,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for userdefinecategory")
	}

CacheNoHooks:
	if !cached {
		userdefinecategoryInsertCacheMut.Lock()
		userdefinecategoryInsertCache[key] = cache
		userdefinecategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Userdefinecategory record using the global executor.
// See Update for more documentation.
func (o *Userdefinecategory) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Userdefinecategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Userdefinecategory) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userdefinecategoryUpdateCacheMut.RLock()
	cache, cached := userdefinecategoryUpdateCache[key]
	userdefinecategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userdefinecategoryAllColumns,
			userdefinecategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update userdefinecategory, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `userdefinecategory` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userdefinecategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, append(wl, userdefinecategoryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update userdefinecategory row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for userdefinecategory")
	}

	if !cached {
		userdefinecategoryUpdateCacheMut.Lock()
		userdefinecategoryUpdateCache[key] = cache
		userdefinecategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userdefinecategoryQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userdefinecategoryQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for userdefinecategory")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for userdefinecategory")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserdefinecategorySlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserdefinecategorySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userdefinecategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `userdefinecategory` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userdefinecategoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userdefinecategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userdefinecategory")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Userdefinecategory) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

var mySQLUserdefinecategoryUniqueColumns = []string{
	"index",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Userdefinecategory) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no userdefinecategory provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userdefinecategoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserdefinecategoryUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userdefinecategoryUpsertCacheMut.RLock()
	cache, cached := userdefinecategoryUpsertCache[key]
	userdefinecategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userdefinecategoryAllColumns,
			userdefinecategoryColumnsWithDefault,
			userdefinecategoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userdefinecategoryAllColumns,
			userdefinecategoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert userdefinecategory, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "userdefinecategory", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `userdefinecategory` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for userdefinecategory")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userdefinecategoryMapping["index"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userdefinecategoryType, userdefinecategoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for userdefinecategory")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for userdefinecategory")
	}

CacheNoHooks:
	if !cached {
		userdefinecategoryUpsertCacheMut.Lock()
		userdefinecategoryUpsertCache[key] = cache
		userdefinecategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Userdefinecategory record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Userdefinecategory) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Userdefinecategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Userdefinecategory) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Userdefinecategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userdefinecategoryPrimaryKeyMapping)
	sql := "DELETE FROM `userdefinecategory` WHERE `index`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from userdefinecategory")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for userdefinecategory")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userdefinecategoryQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userdefinecategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userdefinecategory")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for userdefinecategory")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserdefinecategorySlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserdefinecategorySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userdefinecategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userdefinecategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `userdefinecategory` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userdefinecategoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userdefinecategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for userdefinecategory")
	}

	if len(userdefinecategoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Userdefinecategory) ReloadG() error {
	if o == nil {
		return errors.New("models: no Userdefinecategory provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Userdefinecategory) Reload(exec boil.Executor) error {
	ret, err := FindUserdefinecategory(exec, o.Index)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserdefinecategorySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UserdefinecategorySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserdefinecategorySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserdefinecategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userdefinecategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `userdefinecategory`.* FROM `userdefinecategory` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userdefinecategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserdefinecategorySlice")
	}

	*o = slice

	return nil
}

// UserdefinecategoryExistsG checks if the Userdefinecategory row exists.
func UserdefinecategoryExistsG(index int) (bool, error) {
	return UserdefinecategoryExists(boil.GetDB(), index)
}

// UserdefinecategoryExists checks if the Userdefinecategory row exists.
func UserdefinecategoryExists(exec boil.Executor, index int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `userdefinecategory` where `index`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, index)
	}
	row := exec.QueryRow(sql, index)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if userdefinecategory exists")
	}

	return exists, nil
}
