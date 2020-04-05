package api

import (
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder
	"github.com/MarErm27/todo/models"
	"github.com/uadmin/uadmin"
)

// AddFriendHandler !
func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/add_friend")
	res := map[string]interface{}{}

	// Fetch data from Friend DB
	friend := models.Friend{}

	// Set the parameters of Name, Email, and Password
	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")

	// Validate if the friendName variable is empty.
	if friendName == "" {
		res["status"] = "ERROR"
		res["err_msg"] = "Name is required."
		uadmin.ReturnJSON(w, r, res)
		return
	}

	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	//friend := Friend{}
	if uadmin.Count(&friend, "name = ? OR email = ?", friendName, friendEmail) != 0 {
		// Initialize the error messages
		//errMsg = map[string]string{}
		//errMsg["Name"] = "This Friend is already in the system"
		res["status"] = "This Friend is already in the system"
		uadmin.ReturnJSON(w, r, res)
	} else {
		// Store input into the Name, Email, and Password fields
		friend.Name = friendName
		friend.Email = friendEmail
		friend.Password = friendPassword
		// Store input in the Friend model
		uadmin.Save(&friend)

		res["status"] = "ok"
		uadmin.ReturnJSON(w, r, res)

	}
}
