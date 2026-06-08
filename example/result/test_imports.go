package sqlc

// These structs have the same fields as User in models.go
// After reduce, they become aliases and the imports become unused
type GetUserByIDRow = User

type GetUserByEmailRow = User

type ListUsersRow = User
