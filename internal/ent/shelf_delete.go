// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/shelf"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShelfDelete is the builder for deleting a Shelf entity.
type ShelfDelete struct {
	config
	hooks    []Hook
	mutation *ShelfMutation
}

// Where appends a list predicates to the ShelfDelete builder.
func (sd *ShelfDelete) Where(ps ...predicate.Shelf) *ShelfDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ShelfDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ShelfDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ShelfDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shelf.Table, sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ShelfDeleteOne is the builder for deleting a single Shelf entity.
type ShelfDeleteOne struct {
	sd *ShelfDelete
}

// Where appends a list predicates to the ShelfDelete builder.
func (sdo *ShelfDeleteOne) Where(ps ...predicate.Shelf) *ShelfDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *ShelfDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shelf.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ShelfDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
