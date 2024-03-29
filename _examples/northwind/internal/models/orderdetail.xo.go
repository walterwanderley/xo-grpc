package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// OrderDetail represents a row from 'OrderDetails'.
type OrderDetail struct {
	OrderDetailID sql.NullInt64 `json:"OrderDetailID"` // OrderDetailID
	OrderID       sql.NullInt64 `json:"OrderID"`       // OrderID
	ProductID     sql.NullInt64 `json:"ProductID"`     // ProductID
	Quantity      sql.NullInt64 `json:"Quantity"`      // Quantity
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [OrderDetail] exists in the database.
func (od *OrderDetail) Exists() bool {
	return od._exists
}

// Deleted returns true when the [OrderDetail] has been marked for deletion
// from the database.
func (od *OrderDetail) Deleted() bool {
	return od._deleted
}

// Insert inserts the [OrderDetail] to the database.
func (od *OrderDetail) Insert(ctx context.Context, db DB) error {
	switch {
	case od._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case od._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO OrderDetails (` +
		`OrderDetailID, OrderID, ProductID, Quantity` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	logf(sqlstr, od.OrderID, od.ProductID, od.Quantity)
	res, err := db.ExecContext(ctx, sqlstr, od.OrderDetailID, od.OrderID, od.ProductID, od.Quantity)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	od.OrderDetailID = sql.NullInt64{Valid: true, Int64: id}
	// set exists
	od._exists = true
	return nil
}

// Update updates a [OrderDetail] in the database.
func (od *OrderDetail) Update(ctx context.Context, db DB) error {
	switch {
	case !od._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case od._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE OrderDetails SET ` +
		`OrderID = $1, ProductID = $2, Quantity = $3 ` +
		`WHERE OrderDetailID = $4`
	// run
	logf(sqlstr, od.OrderID, od.ProductID, od.Quantity, od.OrderDetailID)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderID, od.ProductID, od.Quantity, od.OrderDetailID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [OrderDetail] to the database.
func (od *OrderDetail) Save(ctx context.Context, db DB) error {
	if od.Exists() {
		return od.Update(ctx, db)
	}
	return od.Insert(ctx, db)
}

// Upsert performs an upsert for [OrderDetail].
func (od *OrderDetail) Upsert(ctx context.Context, db DB) error {
	switch {
	case od._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO OrderDetails (` +
		`OrderDetailID, OrderID, ProductID, Quantity` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (OrderDetailID) DO ` +
		`UPDATE SET ` +
		`OrderID = EXCLUDED.OrderID, ProductID = EXCLUDED.ProductID, Quantity = EXCLUDED.Quantity `
	// run
	logf(sqlstr, od.OrderDetailID, od.OrderID, od.ProductID, od.Quantity)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderDetailID, od.OrderID, od.ProductID, od.Quantity); err != nil {
		return logerror(err)
	}
	// set exists
	od._exists = true
	return nil
}

// Delete deletes the [OrderDetail] from the database.
func (od *OrderDetail) Delete(ctx context.Context, db DB) error {
	switch {
	case !od._exists: // doesn't exist
		return nil
	case od._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM OrderDetails ` +
		`WHERE OrderDetailID = $1`
	// run
	logf(sqlstr, od.OrderDetailID)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderDetailID); err != nil {
		return logerror(err)
	}
	// set deleted
	od._deleted = true
	return nil
}

// OrderDetailByOrderDetailID retrieves a row from 'OrderDetails' as a [OrderDetail].
//
// Generated from index 'OrderDetails_OrderDetailID_pkey'.
func OrderDetailByOrderDetailID(ctx context.Context, db DB, orderDetailID sql.NullInt64) (*OrderDetail, error) {
	// query
	const sqlstr = `SELECT ` +
		`OrderDetailID, OrderID, ProductID, Quantity ` +
		`FROM OrderDetails ` +
		`WHERE OrderDetailID = $1`
	// run
	logf(sqlstr, orderDetailID)
	od := OrderDetail{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, orderDetailID).Scan(&od.OrderDetailID, &od.OrderID, &od.ProductID, &od.Quantity); err != nil {
		return nil, logerror(err)
	}
	return &od, nil
}

// Order returns the Order associated with the [OrderDetail]'s (OrderID).
//
// Generated from foreign key 'OrderDetails_OrderID_fkey'.
func (od *OrderDetail) Order(ctx context.Context, db DB) (*Order, error) {
	return OrderByOrderID(ctx, db, od.OrderID)
}

// Product returns the Product associated with the [OrderDetail]'s (ProductID).
//
// Generated from foreign key 'OrderDetails_ProductID_fkey'.
func (od *OrderDetail) Product(ctx context.Context, db DB) (*Product, error) {
	return ProductByProductID(ctx, db, od.ProductID)
}
