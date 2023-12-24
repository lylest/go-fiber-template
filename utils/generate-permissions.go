package utils

import "drywave/models"

func GeneratePermissions(roleId int) []models.Permission {
	switch roleId {
	case 2:
		return createAdminPermissions()

	case 3:
		return createStaffPermissions()

	case 4:
		return createCustomerPermissions()

	default:
		return createStaffPermissions()
	}
}

func createAdminPermissions() []models.Permission {
	var newPermissions = make([]models.Permission, 0)

	for index, value := range ModelList {
		permission := models.Permission{
			ID:   index,
			Name: value,
			List: []string{"create", "read", "update", "remove"},
		}
		newPermissions = append(newPermissions, permission)
	}
	return newPermissions
}

func createStaffPermissions() []models.Permission {
	var newPermissions = make([]models.Permission, 0)

	for index, value := range ModelList {
		permission := models.Permission{
			ID:   index,
			Name: value,
			List: []string{},
		}
		newPermissions = append(newPermissions, permission)
	}
	return newPermissions
}

func createCustomerPermissions() []models.Permission {
	var newPermissions = make([]models.Permission, 0)

	for index, value := range ModelList {
		permission := models.Permission{
			ID:   index,
			Name: value,
			List: []string{},
		}

		if value == "customers" {
			permission.List = append(permission.List, "create", "read", "update")
		}

		if value == "shops" {
			permission.List = append(permission.List, "read")
		}

		if value == "services" {
			permission.List = append(permission.List, "read")
		}

		if value == "categories" {
			permission.List = append(permission.List, "read")
		}

		if value == "orders" {
			permission.List = append(permission.List, "create", "read", "update")
		}

		if value == "notifications" {
			permission.List = append(permission.List, "read", "update")
		}

		newPermissions = append(newPermissions, permission)
	}
	return newPermissions
}
