// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/author"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAuthor = "Author"
	TypeBook   = "Book"
)

// AuthorMutation represents an operation that mutates the Author nodes in the graph.
type AuthorMutation struct {
	config
	op            Op
	typ           string
	id            *ksuid.ID
	name          *string
	sort          *string
	link          *string
	clearedFields map[string]struct{}
	books         map[ksuid.ID]struct{}
	removedbooks  map[ksuid.ID]struct{}
	clearedbooks  bool
	done          bool
	oldValue      func(context.Context) (*Author, error)
	predicates    []predicate.Author
}

var _ ent.Mutation = (*AuthorMutation)(nil)

// authorOption allows management of the mutation configuration using functional options.
type authorOption func(*AuthorMutation)

// newAuthorMutation creates new mutation for the Author entity.
func newAuthorMutation(c config, op Op, opts ...authorOption) *AuthorMutation {
	m := &AuthorMutation{
		config:        c,
		op:            op,
		typ:           TypeAuthor,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAuthorID sets the ID field of the mutation.
func withAuthorID(id ksuid.ID) authorOption {
	return func(m *AuthorMutation) {
		var (
			err   error
			once  sync.Once
			value *Author
		)
		m.oldValue = func(ctx context.Context) (*Author, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Author.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAuthor sets the old Author of the mutation.
func withAuthor(node *Author) authorOption {
	return func(m *AuthorMutation) {
		m.oldValue = func(context.Context) (*Author, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AuthorMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AuthorMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Author entities.
func (m *AuthorMutation) SetID(id ksuid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AuthorMutation) ID() (id ksuid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AuthorMutation) IDs(ctx context.Context) ([]ksuid.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []ksuid.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Author.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *AuthorMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AuthorMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AuthorMutation) ResetName() {
	m.name = nil
}

// SetSort sets the "sort" field.
func (m *AuthorMutation) SetSort(s string) {
	m.sort = &s
}

// Sort returns the value of the "sort" field in the mutation.
func (m *AuthorMutation) Sort() (r string, exists bool) {
	v := m.sort
	if v == nil {
		return
	}
	return *v, true
}

// OldSort returns the old "sort" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldSort(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSort: %w", err)
	}
	return oldValue.Sort, nil
}

// ResetSort resets all changes to the "sort" field.
func (m *AuthorMutation) ResetSort() {
	m.sort = nil
}

// SetLink sets the "link" field.
func (m *AuthorMutation) SetLink(s string) {
	m.link = &s
}

// Link returns the value of the "link" field in the mutation.
func (m *AuthorMutation) Link() (r string, exists bool) {
	v := m.link
	if v == nil {
		return
	}
	return *v, true
}

// OldLink returns the old "link" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldLink(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLink is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLink requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLink: %w", err)
	}
	return oldValue.Link, nil
}

// ClearLink clears the value of the "link" field.
func (m *AuthorMutation) ClearLink() {
	m.link = nil
	m.clearedFields[author.FieldLink] = struct{}{}
}

// LinkCleared returns if the "link" field was cleared in this mutation.
func (m *AuthorMutation) LinkCleared() bool {
	_, ok := m.clearedFields[author.FieldLink]
	return ok
}

// ResetLink resets all changes to the "link" field.
func (m *AuthorMutation) ResetLink() {
	m.link = nil
	delete(m.clearedFields, author.FieldLink)
}

// AddBookIDs adds the "books" edge to the Book entity by ids.
func (m *AuthorMutation) AddBookIDs(ids ...ksuid.ID) {
	if m.books == nil {
		m.books = make(map[ksuid.ID]struct{})
	}
	for i := range ids {
		m.books[ids[i]] = struct{}{}
	}
}

// ClearBooks clears the "books" edge to the Book entity.
func (m *AuthorMutation) ClearBooks() {
	m.clearedbooks = true
}

// BooksCleared reports if the "books" edge to the Book entity was cleared.
func (m *AuthorMutation) BooksCleared() bool {
	return m.clearedbooks
}

// RemoveBookIDs removes the "books" edge to the Book entity by IDs.
func (m *AuthorMutation) RemoveBookIDs(ids ...ksuid.ID) {
	if m.removedbooks == nil {
		m.removedbooks = make(map[ksuid.ID]struct{})
	}
	for i := range ids {
		delete(m.books, ids[i])
		m.removedbooks[ids[i]] = struct{}{}
	}
}

// RemovedBooks returns the removed IDs of the "books" edge to the Book entity.
func (m *AuthorMutation) RemovedBooksIDs() (ids []ksuid.ID) {
	for id := range m.removedbooks {
		ids = append(ids, id)
	}
	return
}

// BooksIDs returns the "books" edge IDs in the mutation.
func (m *AuthorMutation) BooksIDs() (ids []ksuid.ID) {
	for id := range m.books {
		ids = append(ids, id)
	}
	return
}

// ResetBooks resets all changes to the "books" edge.
func (m *AuthorMutation) ResetBooks() {
	m.books = nil
	m.clearedbooks = false
	m.removedbooks = nil
}

// Where appends a list predicates to the AuthorMutation builder.
func (m *AuthorMutation) Where(ps ...predicate.Author) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AuthorMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AuthorMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Author, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AuthorMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AuthorMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Author).
func (m *AuthorMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AuthorMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, author.FieldName)
	}
	if m.sort != nil {
		fields = append(fields, author.FieldSort)
	}
	if m.link != nil {
		fields = append(fields, author.FieldLink)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AuthorMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case author.FieldName:
		return m.Name()
	case author.FieldSort:
		return m.Sort()
	case author.FieldLink:
		return m.Link()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AuthorMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case author.FieldName:
		return m.OldName(ctx)
	case author.FieldSort:
		return m.OldSort(ctx)
	case author.FieldLink:
		return m.OldLink(ctx)
	}
	return nil, fmt.Errorf("unknown Author field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AuthorMutation) SetField(name string, value ent.Value) error {
	switch name {
	case author.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case author.FieldSort:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSort(v)
		return nil
	case author.FieldLink:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLink(v)
		return nil
	}
	return fmt.Errorf("unknown Author field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AuthorMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AuthorMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AuthorMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Author numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AuthorMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(author.FieldLink) {
		fields = append(fields, author.FieldLink)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AuthorMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AuthorMutation) ClearField(name string) error {
	switch name {
	case author.FieldLink:
		m.ClearLink()
		return nil
	}
	return fmt.Errorf("unknown Author nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AuthorMutation) ResetField(name string) error {
	switch name {
	case author.FieldName:
		m.ResetName()
		return nil
	case author.FieldSort:
		m.ResetSort()
		return nil
	case author.FieldLink:
		m.ResetLink()
		return nil
	}
	return fmt.Errorf("unknown Author field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AuthorMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.books != nil {
		edges = append(edges, author.EdgeBooks)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AuthorMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case author.EdgeBooks:
		ids := make([]ent.Value, 0, len(m.books))
		for id := range m.books {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AuthorMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedbooks != nil {
		edges = append(edges, author.EdgeBooks)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AuthorMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case author.EdgeBooks:
		ids := make([]ent.Value, 0, len(m.removedbooks))
		for id := range m.removedbooks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AuthorMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedbooks {
		edges = append(edges, author.EdgeBooks)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AuthorMutation) EdgeCleared(name string) bool {
	switch name {
	case author.EdgeBooks:
		return m.clearedbooks
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AuthorMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Author unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AuthorMutation) ResetEdge(name string) error {
	switch name {
	case author.EdgeBooks:
		m.ResetBooks()
		return nil
	}
	return fmt.Errorf("unknown Author edge %s", name)
}

// BookMutation represents an operation that mutates the Book nodes in the graph.
type BookMutation struct {
	config
	op             Op
	typ            string
	id             *ksuid.ID
	title          *string
	sort           *string
	addedAt        *time.Time
	pubDate        *time.Time
	_path          *string
	isbn           *string
	description    *string
	clearedFields  map[string]struct{}
	authors        map[ksuid.ID]struct{}
	removedauthors map[ksuid.ID]struct{}
	clearedauthors bool
	done           bool
	oldValue       func(context.Context) (*Book, error)
	predicates     []predicate.Book
}

var _ ent.Mutation = (*BookMutation)(nil)

// bookOption allows management of the mutation configuration using functional options.
type bookOption func(*BookMutation)

// newBookMutation creates new mutation for the Book entity.
func newBookMutation(c config, op Op, opts ...bookOption) *BookMutation {
	m := &BookMutation{
		config:        c,
		op:            op,
		typ:           TypeBook,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBookID sets the ID field of the mutation.
func withBookID(id ksuid.ID) bookOption {
	return func(m *BookMutation) {
		var (
			err   error
			once  sync.Once
			value *Book
		)
		m.oldValue = func(ctx context.Context) (*Book, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Book.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBook sets the old Book of the mutation.
func withBook(node *Book) bookOption {
	return func(m *BookMutation) {
		m.oldValue = func(context.Context) (*Book, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BookMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BookMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Book entities.
func (m *BookMutation) SetID(id ksuid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BookMutation) ID() (id ksuid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BookMutation) IDs(ctx context.Context) ([]ksuid.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []ksuid.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Book.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *BookMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *BookMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *BookMutation) ResetTitle() {
	m.title = nil
}

// SetSort sets the "sort" field.
func (m *BookMutation) SetSort(s string) {
	m.sort = &s
}

// Sort returns the value of the "sort" field in the mutation.
func (m *BookMutation) Sort() (r string, exists bool) {
	v := m.sort
	if v == nil {
		return
	}
	return *v, true
}

// OldSort returns the old "sort" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldSort(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSort: %w", err)
	}
	return oldValue.Sort, nil
}

// ResetSort resets all changes to the "sort" field.
func (m *BookMutation) ResetSort() {
	m.sort = nil
}

// SetAddedAt sets the "addedAt" field.
func (m *BookMutation) SetAddedAt(t time.Time) {
	m.addedAt = &t
}

// AddedAt returns the value of the "addedAt" field in the mutation.
func (m *BookMutation) AddedAt() (r time.Time, exists bool) {
	v := m.addedAt
	if v == nil {
		return
	}
	return *v, true
}

// OldAddedAt returns the old "addedAt" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldAddedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddedAt: %w", err)
	}
	return oldValue.AddedAt, nil
}

// ResetAddedAt resets all changes to the "addedAt" field.
func (m *BookMutation) ResetAddedAt() {
	m.addedAt = nil
}

// SetPubDate sets the "pubDate" field.
func (m *BookMutation) SetPubDate(t time.Time) {
	m.pubDate = &t
}

// PubDate returns the value of the "pubDate" field in the mutation.
func (m *BookMutation) PubDate() (r time.Time, exists bool) {
	v := m.pubDate
	if v == nil {
		return
	}
	return *v, true
}

// OldPubDate returns the old "pubDate" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldPubDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPubDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPubDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPubDate: %w", err)
	}
	return oldValue.PubDate, nil
}

// ClearPubDate clears the value of the "pubDate" field.
func (m *BookMutation) ClearPubDate() {
	m.pubDate = nil
	m.clearedFields[book.FieldPubDate] = struct{}{}
}

// PubDateCleared returns if the "pubDate" field was cleared in this mutation.
func (m *BookMutation) PubDateCleared() bool {
	_, ok := m.clearedFields[book.FieldPubDate]
	return ok
}

// ResetPubDate resets all changes to the "pubDate" field.
func (m *BookMutation) ResetPubDate() {
	m.pubDate = nil
	delete(m.clearedFields, book.FieldPubDate)
}

// SetPath sets the "path" field.
func (m *BookMutation) SetPath(s string) {
	m._path = &s
}

// Path returns the value of the "path" field in the mutation.
func (m *BookMutation) Path() (r string, exists bool) {
	v := m._path
	if v == nil {
		return
	}
	return *v, true
}

// OldPath returns the old "path" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPath: %w", err)
	}
	return oldValue.Path, nil
}

// ResetPath resets all changes to the "path" field.
func (m *BookMutation) ResetPath() {
	m._path = nil
}

// SetIsbn sets the "isbn" field.
func (m *BookMutation) SetIsbn(s string) {
	m.isbn = &s
}

// Isbn returns the value of the "isbn" field in the mutation.
func (m *BookMutation) Isbn() (r string, exists bool) {
	v := m.isbn
	if v == nil {
		return
	}
	return *v, true
}

// OldIsbn returns the old "isbn" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldIsbn(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsbn is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsbn requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsbn: %w", err)
	}
	return oldValue.Isbn, nil
}

// ClearIsbn clears the value of the "isbn" field.
func (m *BookMutation) ClearIsbn() {
	m.isbn = nil
	m.clearedFields[book.FieldIsbn] = struct{}{}
}

// IsbnCleared returns if the "isbn" field was cleared in this mutation.
func (m *BookMutation) IsbnCleared() bool {
	_, ok := m.clearedFields[book.FieldIsbn]
	return ok
}

// ResetIsbn resets all changes to the "isbn" field.
func (m *BookMutation) ResetIsbn() {
	m.isbn = nil
	delete(m.clearedFields, book.FieldIsbn)
}

// SetDescription sets the "description" field.
func (m *BookMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *BookMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *BookMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[book.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *BookMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[book.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *BookMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, book.FieldDescription)
}

// AddAuthorIDs adds the "authors" edge to the Author entity by ids.
func (m *BookMutation) AddAuthorIDs(ids ...ksuid.ID) {
	if m.authors == nil {
		m.authors = make(map[ksuid.ID]struct{})
	}
	for i := range ids {
		m.authors[ids[i]] = struct{}{}
	}
}

// ClearAuthors clears the "authors" edge to the Author entity.
func (m *BookMutation) ClearAuthors() {
	m.clearedauthors = true
}

// AuthorsCleared reports if the "authors" edge to the Author entity was cleared.
func (m *BookMutation) AuthorsCleared() bool {
	return m.clearedauthors
}

// RemoveAuthorIDs removes the "authors" edge to the Author entity by IDs.
func (m *BookMutation) RemoveAuthorIDs(ids ...ksuid.ID) {
	if m.removedauthors == nil {
		m.removedauthors = make(map[ksuid.ID]struct{})
	}
	for i := range ids {
		delete(m.authors, ids[i])
		m.removedauthors[ids[i]] = struct{}{}
	}
}

// RemovedAuthors returns the removed IDs of the "authors" edge to the Author entity.
func (m *BookMutation) RemovedAuthorsIDs() (ids []ksuid.ID) {
	for id := range m.removedauthors {
		ids = append(ids, id)
	}
	return
}

// AuthorsIDs returns the "authors" edge IDs in the mutation.
func (m *BookMutation) AuthorsIDs() (ids []ksuid.ID) {
	for id := range m.authors {
		ids = append(ids, id)
	}
	return
}

// ResetAuthors resets all changes to the "authors" edge.
func (m *BookMutation) ResetAuthors() {
	m.authors = nil
	m.clearedauthors = false
	m.removedauthors = nil
}

// Where appends a list predicates to the BookMutation builder.
func (m *BookMutation) Where(ps ...predicate.Book) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BookMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BookMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Book, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BookMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BookMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Book).
func (m *BookMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BookMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.title != nil {
		fields = append(fields, book.FieldTitle)
	}
	if m.sort != nil {
		fields = append(fields, book.FieldSort)
	}
	if m.addedAt != nil {
		fields = append(fields, book.FieldAddedAt)
	}
	if m.pubDate != nil {
		fields = append(fields, book.FieldPubDate)
	}
	if m._path != nil {
		fields = append(fields, book.FieldPath)
	}
	if m.isbn != nil {
		fields = append(fields, book.FieldIsbn)
	}
	if m.description != nil {
		fields = append(fields, book.FieldDescription)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BookMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case book.FieldTitle:
		return m.Title()
	case book.FieldSort:
		return m.Sort()
	case book.FieldAddedAt:
		return m.AddedAt()
	case book.FieldPubDate:
		return m.PubDate()
	case book.FieldPath:
		return m.Path()
	case book.FieldIsbn:
		return m.Isbn()
	case book.FieldDescription:
		return m.Description()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BookMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case book.FieldTitle:
		return m.OldTitle(ctx)
	case book.FieldSort:
		return m.OldSort(ctx)
	case book.FieldAddedAt:
		return m.OldAddedAt(ctx)
	case book.FieldPubDate:
		return m.OldPubDate(ctx)
	case book.FieldPath:
		return m.OldPath(ctx)
	case book.FieldIsbn:
		return m.OldIsbn(ctx)
	case book.FieldDescription:
		return m.OldDescription(ctx)
	}
	return nil, fmt.Errorf("unknown Book field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookMutation) SetField(name string, value ent.Value) error {
	switch name {
	case book.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case book.FieldSort:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSort(v)
		return nil
	case book.FieldAddedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddedAt(v)
		return nil
	case book.FieldPubDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPubDate(v)
		return nil
	case book.FieldPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPath(v)
		return nil
	case book.FieldIsbn:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsbn(v)
		return nil
	case book.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BookMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BookMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Book numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BookMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(book.FieldPubDate) {
		fields = append(fields, book.FieldPubDate)
	}
	if m.FieldCleared(book.FieldIsbn) {
		fields = append(fields, book.FieldIsbn)
	}
	if m.FieldCleared(book.FieldDescription) {
		fields = append(fields, book.FieldDescription)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BookMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BookMutation) ClearField(name string) error {
	switch name {
	case book.FieldPubDate:
		m.ClearPubDate()
		return nil
	case book.FieldIsbn:
		m.ClearIsbn()
		return nil
	case book.FieldDescription:
		m.ClearDescription()
		return nil
	}
	return fmt.Errorf("unknown Book nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BookMutation) ResetField(name string) error {
	switch name {
	case book.FieldTitle:
		m.ResetTitle()
		return nil
	case book.FieldSort:
		m.ResetSort()
		return nil
	case book.FieldAddedAt:
		m.ResetAddedAt()
		return nil
	case book.FieldPubDate:
		m.ResetPubDate()
		return nil
	case book.FieldPath:
		m.ResetPath()
		return nil
	case book.FieldIsbn:
		m.ResetIsbn()
		return nil
	case book.FieldDescription:
		m.ResetDescription()
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BookMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.authors != nil {
		edges = append(edges, book.EdgeAuthors)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BookMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case book.EdgeAuthors:
		ids := make([]ent.Value, 0, len(m.authors))
		for id := range m.authors {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BookMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedauthors != nil {
		edges = append(edges, book.EdgeAuthors)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BookMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case book.EdgeAuthors:
		ids := make([]ent.Value, 0, len(m.removedauthors))
		for id := range m.removedauthors {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BookMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedauthors {
		edges = append(edges, book.EdgeAuthors)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BookMutation) EdgeCleared(name string) bool {
	switch name {
	case book.EdgeAuthors:
		return m.clearedauthors
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BookMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Book unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BookMutation) ResetEdge(name string) error {
	switch name {
	case book.EdgeAuthors:
		m.ResetAuthors()
		return nil
	}
	return fmt.Errorf("unknown Book edge %s", name)
}
