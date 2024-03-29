package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Product represents a row from 'Products'.
type Product struct {
	ProductID   sql.NullInt64   `json:"ProductID"`   // ProductID
	ProductName sql.NullString  `json:"ProductName"` // ProductName
	SupplierID  sql.NullInt64   `json:"SupplierID"`  // SupplierID
	CategoryID  sql.NullInt64   `json:"CategoryID"`  // CategoryID
	Unit        sql.NullString  `json:"Unit"`        // Unit
	Price       sql.NullFloat64 `json:"Price"`       // Price
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Product] exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the [Product] has been marked for deletion
// from the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the [Product] to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO Products (` +
		`ProductID, ProductName, SupplierID, CategoryID, Unit, Price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price)
	res, err := db.ExecContext(ctx, sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	p.ProductID = sql.NullInt64{Valid: true, Int64: id}
	// set exists
	p._exists = true
	return nil
}

// Update updates a [Product] in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE Products SET ` +
		`ProductName = $1, SupplierID = $2, CategoryID = $3, Unit = $4, Price = $5 ` +
		`WHERE ProductID = $6`
	// run
	logf(sqlstr, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price, p.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price, p.ProductID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Product] to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for [Product].
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO Products (` +
		`ProductID, ProductName, SupplierID, CategoryID, Unit, Price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (ProductID) DO ` +
		`UPDATE SET ` +
		`ProductName = EXCLUDED.ProductName, SupplierID = EXCLUDED.SupplierID, CategoryID = EXCLUDED.CategoryID, Unit = EXCLUDED.Unit, Price = EXCLUDED.Price `
	// run
	logf(sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.Unit, p.Price); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the [Product] from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM Products ` +
		`WHERE ProductID = $1`
	// run
	logf(sqlstr, p.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProductByProductID retrieves a row from 'Products' as a [Product].
//
// Generated from index 'Products_ProductID_pkey'.
func ProductByProductID(ctx context.Context, db DB, productID sql.NullInt64) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`ProductID, ProductName, SupplierID, CategoryID, Unit, Price ` +
		`FROM Products ` +
		`WHERE ProductID = $1`
	// run
	logf(sqlstr, productID)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, productID).Scan(&p.ProductID, &p.ProductName, &p.SupplierID, &p.CategoryID, &p.Unit, &p.Price); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}

// Category returns the Category associated with the [Product]'s (CategoryID).
//
// Generated from foreign key 'Products_CategoryID_fkey'.
func (p *Product) Category(ctx context.Context, db DB) (*Category, error) {
	return CategoryByCategoryID(ctx, db, p.CategoryID)
}

// Supplier returns the Supplier associated with the [Product]'s (SupplierID).
//
// Generated from foreign key 'Products_SupplierID_fkey'.
func (p *Product) Supplier(ctx context.Context, db DB) (*Supplier, error) {
	return SupplierBySupplierID(ctx, db, p.SupplierID)
}
