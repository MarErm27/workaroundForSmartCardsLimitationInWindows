package models

import "github.com/uadmin/uadmin"

// Category model ...
type Category struct {
	uadmin.Model
	Name string `uadmin:"required"`
	Icon string `uadmin:"image"`
}

// Validate function ...
func (c Category) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	category := Category{}
	if uadmin.Count(&category, "name = ? AND id <> ?", c.Name, c.ID) != 0 {
		errMsg["Name"] = "This Item name is already in the system"
	}
	return
}
