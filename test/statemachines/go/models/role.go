package models

type Role struct {
	Name string

	// "U" for Unit, "S" for supervisor
	Acronym string

	// Role that inherit the access right of this Role
	RolesWithSamePermissions []*Role
}
