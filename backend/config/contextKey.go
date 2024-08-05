package config

// UserIDKey is the context key for accessing the user ID set by the authentication middleware.
type contextKey string

const UserIDKey contextKey = "user_id"
