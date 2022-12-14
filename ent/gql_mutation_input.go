// Code generated by ent, DO NOT EDIT.

package ent

// CreateFriendshipInput represents a mutation input for creating friendships.
type CreateFriendshipInput struct {
	UserID   int
	FriendID int
}

// Mutate applies the CreateFriendshipInput on the FriendshipMutation builder.
func (i *CreateFriendshipInput) Mutate(m *FriendshipMutation) {
	m.SetUserID(i.UserID)
	m.SetFriendID(i.FriendID)
}

// SetInput applies the change-set in the CreateFriendshipInput on the FriendshipCreate builder.
func (c *FriendshipCreate) SetInput(i CreateFriendshipInput) *FriendshipCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateFriendshipInput represents a mutation input for updating friendships.
type UpdateFriendshipInput struct {
	ClearUser   bool
	UserID      *int
	ClearFriend bool
	FriendID    *int
}

// Mutate applies the UpdateFriendshipInput on the FriendshipMutation builder.
func (i *UpdateFriendshipInput) Mutate(m *FriendshipMutation) {
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if i.ClearFriend {
		m.ClearFriend()
	}
	if v := i.FriendID; v != nil {
		m.SetFriendID(*v)
	}
}

// SetInput applies the change-set in the UpdateFriendshipInput on the FriendshipUpdate builder.
func (c *FriendshipUpdate) SetInput(i UpdateFriendshipInput) *FriendshipUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateFriendshipInput on the FriendshipUpdateOne builder.
func (c *FriendshipUpdateOne) SetInput(i UpdateFriendshipInput) *FriendshipUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name      *string
	FriendIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.FriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name            *string
	AddFriendIDs    []int
	RemoveFriendIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddFriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.RemoveFriendIDs; len(v) > 0 {
		m.RemoveFriendIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
