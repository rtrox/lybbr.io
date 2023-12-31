// Code generated by ent, DO NOT EDIT.

package userpermissions

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the userpermissions type in the database.
	Label = "user_permissions"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCanEdit holds the string denoting the canedit field in the database.
	FieldCanEdit = "can_edit"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldCanCreatePublic holds the string denoting the cancreatepublic field in the database.
	FieldCanCreatePublic = "can_create_public"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userpermissions in the database.
	Table = "user_permissions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_permissions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userpermissions fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldCanEdit,
	FieldAdmin,
	FieldCanCreatePublic,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "lybbrio/internal/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCanEdit holds the default value on creation for the "CanEdit" field.
	DefaultCanEdit bool
	// DefaultAdmin holds the default value on creation for the "Admin" field.
	DefaultAdmin bool
	// DefaultCanCreatePublic holds the default value on creation for the "CanCreatePublic" field.
	DefaultCanCreatePublic bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ksuid.ID
)

// OrderOption defines the ordering options for the UserPermissions queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCanEdit orders the results by the CanEdit field.
func ByCanEdit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCanEdit, opts...).ToFunc()
}

// ByAdmin orders the results by the Admin field.
func ByAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdmin, opts...).ToFunc()
}

// ByCanCreatePublic orders the results by the CanCreatePublic field.
func ByCanCreatePublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCanCreatePublic, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
