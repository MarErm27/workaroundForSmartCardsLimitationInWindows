package models

import "github.com/uadmin/uadmin"

// Nationality ...
type Nationality int

// Chinese ...
func (Nationality) Chinese() Nationality {
	return 1
}

// Filipino ...
func (Nationality) Filipino() Nationality {
	return 2
}

// Others ...
func (Nationality) Others() Nationality {
	return 3
}

// Friend model ...
type Friend struct {
	uadmin.Model
	Name        string      `uadmin:"required"`
	Email       string      `uadmin:"email"`
	Password    string      `uadmin:"password;list_exclude"`
	Nationality Nationality // <-- place it here
	Invite      string      `uadmin:"link"` // <-- place it here
}

// Save !
func (f *Friend) Save() {
	f.Invite = "https://www.google.com/"
	uadmin.Save(f)
}

func (f Friend) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	friend := Friend{}
	if uadmin.Count(&friend, "name = ? AND id <> ?", f.Name, f.ID) != 0 {
		errMsg["Name"] = "This Item name is already in the system"
	}
	return
}
