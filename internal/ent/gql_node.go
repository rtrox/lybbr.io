// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lybbrio/internal/ent/author"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"lybbrio/internal/ent/shelf"
	"lybbrio/internal/ent/tag"
	"lybbrio/internal/ent/task"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/ent/userpermissions"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Author) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Book) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Identifier) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Language) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Publisher) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Series) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Shelf) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Tag) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Task) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *User) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *UserPermissions) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, ksuid.ID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, ksuid.ID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, ksuid.ID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id ksuid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id ksuid.ID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id ksuid.ID) (Noder, error) {
	switch table {
	case author.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Author.Query().
			Where(author.ID(uid))
		query, err := query.CollectFields(ctx, "Author")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case book.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Book.Query().
			Where(book.ID(uid))
		query, err := query.CollectFields(ctx, "Book")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case identifier.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Identifier.Query().
			Where(identifier.ID(uid))
		query, err := query.CollectFields(ctx, "Identifier")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case language.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Language.Query().
			Where(language.ID(uid))
		query, err := query.CollectFields(ctx, "Language")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case publisher.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Publisher.Query().
			Where(publisher.ID(uid))
		query, err := query.CollectFields(ctx, "Publisher")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case series.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Series.Query().
			Where(series.ID(uid))
		query, err := query.CollectFields(ctx, "Series")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case shelf.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Shelf.Query().
			Where(shelf.ID(uid))
		query, err := query.CollectFields(ctx, "Shelf")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tag.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Tag.Query().
			Where(tag.ID(uid))
		query, err := query.CollectFields(ctx, "Tag")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case task.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Task.Query().
			Where(task.ID(uid))
		query, err := query.CollectFields(ctx, "Task")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.User.Query().
			Where(user.ID(uid))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case userpermissions.Table:
		var uid ksuid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.UserPermissions.Query().
			Where(userpermissions.ID(uid))
		query, err := query.CollectFields(ctx, "UserPermissions")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []ksuid.ID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]ksuid.ID)
	id2idx := make(map[ksuid.ID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []ksuid.ID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[ksuid.ID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case author.Table:
		query := c.Author.Query().
			Where(author.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Author")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case book.Table:
		query := c.Book.Query().
			Where(book.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Book")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case identifier.Table:
		query := c.Identifier.Query().
			Where(identifier.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Identifier")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case language.Table:
		query := c.Language.Query().
			Where(language.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Language")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case publisher.Table:
		query := c.Publisher.Query().
			Where(publisher.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Publisher")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case series.Table:
		query := c.Series.Query().
			Where(series.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Series")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case shelf.Table:
		query := c.Shelf.Query().
			Where(shelf.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Shelf")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tag.Table:
		query := c.Tag.Query().
			Where(tag.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Tag")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case task.Table:
		query := c.Task.Query().
			Where(task.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Task")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userpermissions.Table:
		query := c.UserPermissions.Query().
			Where(userpermissions.IDIn(ids...))
		query, err := query.CollectFields(ctx, "UserPermissions")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
