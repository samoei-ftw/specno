package enums

type RoleType string

const (
	RoleUser  RoleType = "User"
	RoleAdmin RoleType = "Admin"
)

type PermissionType string

const (
	PermissionCreateProject PermissionType = "CreateProject"
	PermissionDeleteUser    PermissionType = "DeleteUser"
	PermissionAssignUser    PermissionType = "AssignUser"
)

func (r RoleType) IsValid() bool {
	return r == RoleUser || r == RoleAdmin
}

func (p PermissionType) IsValid() bool {
	switch p {
	case PermissionCreateProject, PermissionDeleteUser, PermissionAssignUser:
		return true
	}
	return false
}
