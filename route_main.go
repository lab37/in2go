package main

import (
	"001/data"
	"html/template"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		templates := template.Must(template.ParseFiles("templates/index.html"))
		templates.Execute(writer, "")

	}
}

func htmlTurn(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		keys := request.URL.Query()

		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		dest := keys["dest"][0]
		templates := template.Must(template.ParseFiles("templates/jxc/" + dest))
		templates.Execute(writer, "")

	}

}

func xls2cat(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		templates := template.Must(template.ParseFiles("templates/xls2cat.html"))
		templates.Execute(writer, "")
	}

}

func discount(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		templates := template.Must(template.ParseFiles("templates/discount.html"))
		templates.Execute(writer, "")
	}

}

func tax(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		templates := template.Must(template.ParseFiles("templates/tax.html"))
		templates.Execute(writer, "")
	}

}

func webchat(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "webchat")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "webchat")
		}
	}
}

func chgAccount(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, _ := sess.GetUser()
		templates := template.Must(template.ParseFiles("templates/chgacount.html"))
		templates.Execute(writer, user)
	}
}
