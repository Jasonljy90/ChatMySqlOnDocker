package main

import (
	"SessionCookies/token"
	"database/sql"
	"fmt"
	"html/template"

	//"io"
	"flag"
	"log"
	"net/http"
	"sync"

	//"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Declaration of global variable
var (
	tpl   *template.Template
	mutex sync.Mutex // Concurrency
	db    *sql.DB
	err1  error
)

// Declaration of type patient
type messageTest struct {
	MessageName    string `validate:"required"`
	MessageContent string `validate:"required"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func home(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "homePage.gohtml", nil)
}

func userLoginSuccess(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "userLoginSuccess.gohtml", nil)
}

//Load environment variables
func getEnvVars() {
	err := godotenv.Load("secret.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func userChat(res http.ResponseWriter, req *http.Request) {
	newUrl := "http://localhost:8000"
	http.Redirect(res, req, newUrl, http.StatusSeeOther)
}

/*
func userChat(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		msgName := req.FormValue("messagename")
		msgContent := req.FormValue("messagecontent")

		m := messageTest{msgName, msgContent}
		validate := validator.New()
		err := validate.Struct(m)
		if err != nil {
			io.WriteString(res, `
			<html>
			<meta http-equiv='refresh' content='5; url=/userchat '/>
			Please fill in all fields!<br>
			You will be redirected shortly in 5 seconds...<br>
			</html>
			`)
			return
		}
		SendMsg(msgName, msgContent) // publish message to producer
	}
	tpl.ExecuteTemplate(res, "chatAgentRider.gohtml", nil)
}
*/

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()

	//Opening Docker Database
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/userAccountDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Loading environment variables
	getEnvVars()

	// Generate key and convert to string for password reset feature
	key := generateSecretKey()
	keyStr = string(key)
	maker, err = token.NewJWTMaker(keyStr)
	if err != nil {
		fmt.Println("Error generating token maker!")
	}

	// http multiplexer with gorilla/mux
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/usersignup", userSignup)
	router.HandleFunc("/userlogin", userLogin)
	router.HandleFunc("/userloginsuccess", userLoginSuccess)
	router.HandleFunc("/userchat", userChat)
	router.HandleFunc("/deleteuser", deleteUser)
	router.HandleFunc("/userchangepassword", userChangePassword)
	router.HandleFunc("/userchangelanguage", userChangeLanguage)
	router.HandleFunc("/userresetpassword", userResetPassword)
	router.HandleFunc("/userresetchangepassword", userResetChangePassword)
	router.HandleFunc("/usertoken/{token}", resetUserPasswordLinkClicked)
	fmt.Println("Server is up")
	// http.ListenAndServeTLS(":5221", "cert.pem", "key.pem", router)
	http.ListenAndServe(":5221", router)

	http.HandleFunc("/userchat", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err1 := http.ListenAndServe(*addr, nil)
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
