package models

// SecurityRole is represent one role in the RBAC system
type SecurityRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
