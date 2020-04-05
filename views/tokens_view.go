package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MarErm27/readAndWriteKeepass"
	"github.com/MarErm27/tokenControl"

	"github.com/uadmin/uadmin"
	// Specify the username that you used inside github.com folder and
	// import this library
)

type smartCard struct {
	SecondName string
	NotBefore string
	TITLE       string
	KPP         string
	ID          string
	Fingerprint string
	State       bool
	Many        bool
	Number      int
}

//Cards represent smartcards in the system
// var Cards = []smartCard{
// 	// smartCard{"123", "abc", true},
// 	// smartCard{"456", "def", false},
// 	// smartCard{"789", "ghi", true},
// }

// Context !
type Context struct {
	Data    []smartCard
	URLPath string
}

var previousFingerPrint string
var previousDeviceID string

var buttonState bool
var previousState bool

// TokensHandler !
func TokensHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is not logged in
	if uadmin.IsAuthenticated(r) == nil {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	context := Context{}
	r.Host += "/http_handler/tokens"
	fmt.Printf("Req: %s\n", r.Host)

	value := r.FormValue("id")
	initializeDevices := r.FormValue("initialization")
	if initializeDevices == "true" {
		tokenControl.InitializationOfDevices()
	}
	fmt.Println("Got id:", value)
	rootGroup := readAndWriteKeepass.ReadFromKeepass()
	var s smartCard
	Cards := make([]smartCard, 0)
	subGroup := rootGroup.Groups[1]
	//for subGroupNumber, subGroup := range rootGroup.Groups {
	tn := 0
	for entryNumber, entry := range subGroup.Entries {
		tn += 1
		for entryFieldNumber, entryField := range entry.Values {
			fmt.Println(entryNumber, entryFieldNumber, entryField.Key, entryField.Value.Content)

			if entryField.Key == "Фамилия" {
				s.SecondName = entryField.Value.Content
			}
			if entryField.Key == "NotBefore" {
				s.NotBefore = entryField.Value.Content
			}
			
			if entryField.Key == "Fingerprint" {
				s.Fingerprint = entryField.Value.Content
			}
			if entryField.Key == "deviceID" {
				s.ID = entryField.Value.Content
				//fmt.Println("--------------------------")
				//continue
			}
			if entryField.Key == "Title" {
				s.TITLE = entryField.Value.Content
			}
			if entryField.Key == "UnstructuredName" {
				UnstructuredName := strings.Split(entryField.Value.Content, "-")
				if len(UnstructuredName) > 1 {
					s.KPP = strings.Split(entryField.Value.Content, "-")[1]
				}
			}
			if entryField.Key == "deviceState" {
				if entryField.Value.Content == "true" {
					s.State = true
				} else {
					s.State = false
				}
			}
			if entryField.Key == "numberOfCerts" {
				if entryField.Value.Content == "many" {
					s.Many = true
				} else {
					s.Many = false
				}
			}
			s.Number = tn
		}
		Cards = append(Cards, s)
	}

	//}
	n := 0
	for i, e := range Cards {
		Cards[i].State = tokenControl.CheckState(e.ID)

		if e.ID == value {
			if s.Many && n > 0 { // e.ID == previousDeviceID && e.State == buttonState && e.State == buttonState && e.ID == previousDeviceID e.Fingerprint != previousFingerPrint &&
				previousFingerPrint = e.Fingerprint
				previousDeviceID = e.ID
				//previousState = !previousState
				fmt.Printf("Switching state for id %s from %t to %t\n", e.ID, e.State, !e.State)

			} else {
				fmt.Printf("Switching state for id %s from %t to %t\n", e.ID, e.State, !e.State)
				buttonState = tokenControl.ChangeDeviceState(e.ID)
				n += 1

			}
			previousFingerPrint = e.Fingerprint
			previousDeviceID = e.ID
			Cards[i].State = buttonState
			//Cards[i].State = !Cards[i].State
		}
	}

	context.Data = Cards

	// Include this one to get the URL Path and return to your HTML
	context.URLPath = r.Host

	uadmin.RenderHTML(w, r, "templates/tokens.html", context)

	// //Check if incoming request have your header
	// //if r.Header.Get("your header") == "your value"
	// tmpl := template.Must(template.ParseFiles("templates/tokens.html"))
	// //w.Header().Add("Access-Control-Allow-Origin", "<your site here>")
	// if r.Method != "POST" {
	// 	tmpl.Execute(w, Cards)
	// 	return
	// }

	// value := r.FormValue("id")
	// fmt.Println("Got id:", value)

	// for i, e := range Cards {
	// 	if e.ID == value {
	// 		fmt.Printf("Switching state for id %s from %t to %t\n", e.ID, e.State, !e.State)
	// 		Cards[i].State = !Cards[i].State
	// 	}
	// }

	// //tmpl.Execute(w, Cards)
	// //http.Redirect(w, r, "/", 303)
	// http.Redirect(w, r, "/http_handler/tokens", http.StatusSeeOther)
}
