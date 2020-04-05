package models

import "github.com/uadmin/uadmin"

// Item model ...
type Item struct {
	uadmin.Model
	Name         string     `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`
	Description  string     `uadmin:"multilingual"`
	Category     []Category `uadmin:"list_exclude"`
	CategoryList string     `uadmin:"read_only;hidden"`
	Cost         int        `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`
	Rating       int        `uadmin:"min:1;max:5"`
}

// CategorySave ...
func (i *Item) CategorySave() {
	// Initializes the catList as empty string
	catList := ""

	// This process will get the name of the category, store into the
	// catList and if the index value is not equal to the number of
	// category, it will insert the comma symbol at the end of the word.
	for x, key := range i.Category {
		catList += key.Name
		if x != len(i.Category)-1 {
			catList += ", "
		}
	}

	// Store the catList variable to the CategoryList field in the Item model
	i.CategoryList = catList

	// Override save
	uadmin.Save(i)
}

// Save ...
func (i *Item) Save() {
	if i.ID == 0 {
		i.CategorySave()
	}

	i.CategorySave()
}

// Validate function ...
func (i Item) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	item := Item{}
	if uadmin.Count(&item, "name = ? AND id <> ?", i.Name, i.ID) != 0 {
		errMsg["Name"] = "This Item name is already in the system"
	}
	return
}
