// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FriendshipQuery) CollectFields(ctx context.Context, satisfies ...string) (*FriendshipQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FriendshipQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: f.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withUser = query
		case "friend":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: f.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withFriend = query
		}
	}
	return nil
}

type friendshipPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FriendshipPaginateOption
}

func newFriendshipPaginateArgs(rv map[string]interface{}) *friendshipPaginateArgs {
	args := &friendshipPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &FriendshipOrder{Field: &FriendshipOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithFriendshipOrder(order))
			}
		case *FriendshipOrder:
			if v != nil {
				args.opts = append(args.opts, WithFriendshipOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*FriendshipWhereInput); ok {
		args.opts = append(args.opts, WithFriendshipFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "friends":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFriends(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "friendships":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &FriendshipQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFriendships(alias, func(wq *FriendshipQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
