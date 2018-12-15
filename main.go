package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	mux.HandleFunc("/webchat", webchat)
	mux.HandleFunc("/htmlturn", htmlTurn)
	mux.HandleFunc("/chgaccount", chgAccount)
	mux.HandleFunc("/xls2cat", xls2cat)
	mux.HandleFunc("/discount", discount)
	mux.HandleFunc("/tax", tax)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_jxc.go
	mux.HandleFunc("/insertitem", insertItem)
	mux.HandleFunc("/selectitem", selectItem)
	mux.HandleFunc("/getcustomername", getCustomerName)
	mux.HandleFunc("/getallccid", getAllCcId)
	mux.HandleFunc("/getproductns", getProductNS)
	mux.HandleFunc("/deleteitem", deleteItem)
	mux.HandleFunc("/updateitem", updateItem)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/updateaccount", updateAccount)
	mux.HandleFunc("/updatepassword", updatePassword)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
