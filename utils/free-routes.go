package utils

var FreeRoutes = []Route{
	{
		Model:  "users",
		Action: []string{"create", "read"},
	},
}

type Route struct {
	Model  string
	Action []string
}
