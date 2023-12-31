// Code generated by ent, DO NOT EDIT.

package task

import (
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task_enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ksuid.ID) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateTime, v))
}

// Progress applies equality check predicate on the "progress" field. It's identical to ProgressEQ.
func Progress(v float64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldProgress, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldMessage, v))
}

// Error applies equality check predicate on the "error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldError, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldEQ(FieldUserID, vc))
}

// IsSystemTask applies equality check predicate on the "is_system_task" field. It's identical to IsSystemTaskEQ.
func IsSystemTask(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldIsSystemTask, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdateTime, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v task_enums.TaskType) predicate.Task {
	vc := v
	return predicate.Task(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v task_enums.TaskType) predicate.Task {
	vc := v
	return predicate.Task(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...task_enums.TaskType) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...task_enums.TaskType) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(sql.FieldNotIn(FieldType, v...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v task_enums.Status) predicate.Task {
	vc := v
	return predicate.Task(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v task_enums.Status) predicate.Task {
	vc := v
	return predicate.Task(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...task_enums.Status) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...task_enums.Status) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(sql.FieldNotIn(FieldStatus, v...))
}

// ProgressEQ applies the EQ predicate on the "progress" field.
func ProgressEQ(v float64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldProgress, v))
}

// ProgressNEQ applies the NEQ predicate on the "progress" field.
func ProgressNEQ(v float64) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldProgress, v))
}

// ProgressIn applies the In predicate on the "progress" field.
func ProgressIn(vs ...float64) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldProgress, vs...))
}

// ProgressNotIn applies the NotIn predicate on the "progress" field.
func ProgressNotIn(vs ...float64) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldProgress, vs...))
}

// ProgressGT applies the GT predicate on the "progress" field.
func ProgressGT(v float64) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldProgress, v))
}

// ProgressGTE applies the GTE predicate on the "progress" field.
func ProgressGTE(v float64) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldProgress, v))
}

// ProgressLT applies the LT predicate on the "progress" field.
func ProgressLT(v float64) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldProgress, v))
}

// ProgressLTE applies the LTE predicate on the "progress" field.
func ProgressLTE(v float64) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldProgress, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldMessage))
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldMessage))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldMessage, v))
}

// ErrorEQ applies the EQ predicate on the "error" field.
func ErrorEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldError, v))
}

// ErrorNEQ applies the NEQ predicate on the "error" field.
func ErrorNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldError, v))
}

// ErrorIn applies the In predicate on the "error" field.
func ErrorIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldError, vs...))
}

// ErrorNotIn applies the NotIn predicate on the "error" field.
func ErrorNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldError, vs...))
}

// ErrorGT applies the GT predicate on the "error" field.
func ErrorGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldError, v))
}

// ErrorGTE applies the GTE predicate on the "error" field.
func ErrorGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldError, v))
}

// ErrorLT applies the LT predicate on the "error" field.
func ErrorLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldError, v))
}

// ErrorLTE applies the LTE predicate on the "error" field.
func ErrorLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldError, v))
}

// ErrorContains applies the Contains predicate on the "error" field.
func ErrorContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldError, v))
}

// ErrorHasPrefix applies the HasPrefix predicate on the "error" field.
func ErrorHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldError, v))
}

// ErrorHasSuffix applies the HasSuffix predicate on the "error" field.
func ErrorHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldError, v))
}

// ErrorIsNil applies the IsNil predicate on the "error" field.
func ErrorIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldError))
}

// ErrorNotNil applies the NotNil predicate on the "error" field.
func ErrorNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldError))
}

// ErrorEqualFold applies the EqualFold predicate on the "error" field.
func ErrorEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldError, v))
}

// ErrorContainsFold applies the ContainsFold predicate on the "error" field.
func ErrorContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldError, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...ksuid.ID) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Task(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...ksuid.ID) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Task(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldLTE(FieldUserID, vc))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v ksuid.ID) predicate.Task {
	vc := string(v)
	return predicate.Task(sql.FieldContainsFold(FieldUserID, vc))
}

// IsSystemTaskEQ applies the EQ predicate on the "is_system_task" field.
func IsSystemTaskEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldIsSystemTask, v))
}

// IsSystemTaskNEQ applies the NEQ predicate on the "is_system_task" field.
func IsSystemTaskNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldIsSystemTask, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(sql.NotPredicates(p))
}
