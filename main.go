package main

import (
	//"time"
	"net/http"

	"github.com/MarErm27/todo/api"
	//"github.com/MarErm27/todo/models"
	"github.com/MarErm27/todo/views"
	"github.com/uadmin/uadmin"
)

func redirectToAdmin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func main() {

	uadmin.RootURL = "/admin/" // <-- place it here

	// uadmin.Register(
	// 	models.Todo{},
	// 	models.Category{}, // <-- place it here
	// 	models.Friend{},
	// 	models.Справочники{},
	// 	models.Item{}, //  <-- place it here
	// )

	// uadmin.RegisterInlines(models.Category{}, map[string]string{
	// 	"Todo": "CategoryID",
	// })
	// uadmin.RegisterInlines(models.Friend{}, map[string]string{
	// 	"Todo": "FriendID",
	// })
	// uadmin.RegisterInlines(models.Item{}, map[string]string{
	// 	"Todo": "ItemID",
	// })
	//http.HandleFunc("/admin/api/", api.APIHandler) // <-- place it here
	// httpMux := http.NewServeMux()
	// httpMux.Handle("/", http.HandlerFunc(redirectToAdmin))
	http.HandleFunc("/", http.HandlerFunc(redirectToAdmin))
	http.HandleFunc("/api/", api.APIHandler) // <-- place it here
	http.HandleFunc("/http_handler/", views.HTTPHandler)
	uadmin.StartServer()
}
