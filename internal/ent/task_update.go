// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/ent/task"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TaskUpdate) SetUpdateTime(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(te task_enums.Status) *TaskUpdate {
	tu.mutation.SetStatus(te)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(te *task_enums.Status) *TaskUpdate {
	if te != nil {
		tu.SetStatus(*te)
	}
	return tu
}

// SetProgress sets the "progress" field.
func (tu *TaskUpdate) SetProgress(f float64) *TaskUpdate {
	tu.mutation.ResetProgress()
	tu.mutation.SetProgress(f)
	return tu
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableProgress(f *float64) *TaskUpdate {
	if f != nil {
		tu.SetProgress(*f)
	}
	return tu
}

// AddProgress adds f to the "progress" field.
func (tu *TaskUpdate) AddProgress(f float64) *TaskUpdate {
	tu.mutation.AddProgress(f)
	return tu
}

// SetMessage sets the "message" field.
func (tu *TaskUpdate) SetMessage(s string) *TaskUpdate {
	tu.mutation.SetMessage(s)
	return tu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableMessage(s *string) *TaskUpdate {
	if s != nil {
		tu.SetMessage(*s)
	}
	return tu
}

// ClearMessage clears the value of the "message" field.
func (tu *TaskUpdate) ClearMessage() *TaskUpdate {
	tu.mutation.ClearMessage()
	return tu
}

// SetError sets the "error" field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.mutation.SetError(s)
	return tu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of the "error" field.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.mutation.ClearError()
	return tu
}

// SetIsSystemTask sets the "is_system_task" field.
func (tu *TaskUpdate) SetIsSystemTask(b bool) *TaskUpdate {
	tu.mutation.SetIsSystemTask(b)
	return tu
}

// SetNillableIsSystemTask sets the "is_system_task" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableIsSystemTask(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetIsSystemTask(*b)
	}
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() error {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		if task.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized task.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := task.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.SetField(task.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Progress(); ok {
		_spec.SetField(task.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedProgress(); ok {
		_spec.AddField(task.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Message(); ok {
		_spec.SetField(task.FieldMessage, field.TypeString, value)
	}
	if tu.mutation.MessageCleared() {
		_spec.ClearField(task.FieldMessage, field.TypeString)
	}
	if value, ok := tu.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tu.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	if value, ok := tu.mutation.IsSystemTask(); ok {
		_spec.SetField(task.FieldIsSystemTask, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TaskUpdateOne) SetUpdateTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(te task_enums.Status) *TaskUpdateOne {
	tuo.mutation.SetStatus(te)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(te *task_enums.Status) *TaskUpdateOne {
	if te != nil {
		tuo.SetStatus(*te)
	}
	return tuo
}

// SetProgress sets the "progress" field.
func (tuo *TaskUpdateOne) SetProgress(f float64) *TaskUpdateOne {
	tuo.mutation.ResetProgress()
	tuo.mutation.SetProgress(f)
	return tuo
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableProgress(f *float64) *TaskUpdateOne {
	if f != nil {
		tuo.SetProgress(*f)
	}
	return tuo
}

// AddProgress adds f to the "progress" field.
func (tuo *TaskUpdateOne) AddProgress(f float64) *TaskUpdateOne {
	tuo.mutation.AddProgress(f)
	return tuo
}

// SetMessage sets the "message" field.
func (tuo *TaskUpdateOne) SetMessage(s string) *TaskUpdateOne {
	tuo.mutation.SetMessage(s)
	return tuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableMessage(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetMessage(*s)
	}
	return tuo
}

// ClearMessage clears the value of the "message" field.
func (tuo *TaskUpdateOne) ClearMessage() *TaskUpdateOne {
	tuo.mutation.ClearMessage()
	return tuo
}

// SetError sets the "error" field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.mutation.SetError(s)
	return tuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of the "error" field.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.mutation.ClearError()
	return tuo
}

// SetIsSystemTask sets the "is_system_task" field.
func (tuo *TaskUpdateOne) SetIsSystemTask(b bool) *TaskUpdateOne {
	tuo.mutation.SetIsSystemTask(b)
	return tuo
}

// SetNillableIsSystemTask sets the "is_system_task" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableIsSystemTask(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetIsSystemTask(*b)
	}
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		if task.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized task.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := task.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.SetField(task.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Progress(); ok {
		_spec.SetField(task.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedProgress(); ok {
		_spec.AddField(task.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Message(); ok {
		_spec.SetField(task.FieldMessage, field.TypeString, value)
	}
	if tuo.mutation.MessageCleared() {
		_spec.ClearField(task.FieldMessage, field.TypeString)
	}
	if value, ok := tuo.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tuo.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	if value, ok := tuo.mutation.IsSystemTask(); ok {
		_spec.SetField(task.FieldIsSystemTask, field.TypeBool, value)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
