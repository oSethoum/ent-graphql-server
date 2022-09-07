// Code generated by ent, DO NOT EDIT.

package friendship

import (
	"time"
)

const (
	// Label holds the string label denoting the friendship type in the database.
	Label = "friendship"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldFriendID holds the string denoting the friend_id field in the database.
	FieldFriendID = "friend_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeFriend holds the string denoting the friend edge name in mutations.
	EdgeFriend = "friend"
	// Table holds the table name of the friendship in the database.
	Table = "friendships"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "friendships"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// FriendTable is the table that holds the friend relation/edge.
	FriendTable = "friendships"
	// FriendInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	FriendInverseTable = "users"
	// FriendColumn is the table column denoting the friend relation/edge.
	FriendColumn = "friend_id"
)

// Columns holds all SQL columns for friendship fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUserID,
	FieldFriendID,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
