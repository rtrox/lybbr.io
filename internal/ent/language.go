// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/schema/ksuid"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Language is the model entity for the Language schema.
type Language struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LanguageQuery when eager-loading is set.
	Edges        LanguageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LanguageEdges holds the relations/edges for other nodes in the graph.
type LanguageEdges struct {
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedBooks map[string][]*Book
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e LanguageEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Language) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case language.FieldID, language.FieldName, language.FieldCode:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Language fields.
func (l *Language) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case language.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				l.ID = ksuid.ID(value.String)
			}
		case language.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case language.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				l.Code = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Language.
// This includes values selected through modifiers, order, etc.
func (l *Language) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the Language entity.
func (l *Language) QueryBooks() *BookQuery {
	return NewLanguageClient(l.config).QueryBooks(l)
}

// Update returns a builder for updating this Language.
// Note that you need to call Language.Unwrap() before calling this method if this Language
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Language) Update() *LanguageUpdateOne {
	return NewLanguageClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Language entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Language) Unwrap() *Language {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Language is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Language) String() string {
	var builder strings.Builder
	builder.WriteString("Language(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(l.Code)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Language) NamedBooks(name string) ([]*Book, error) {
	if l.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Language) appendNamedBooks(name string, edges ...*Book) {
	if l.Edges.namedBooks == nil {
		l.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		l.Edges.namedBooks[name] = []*Book{}
	} else {
		l.Edges.namedBooks[name] = append(l.Edges.namedBooks[name], edges...)
	}
}

// Languages is a parsable slice of Language.
type Languages []*Language