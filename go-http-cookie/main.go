package main

import (
	"fmt"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

// M will be the map later on use
type M map[string]interface{}

var cookieName = "CookieData"

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

// ActionIndex is a function that will handle the process when user hit the server
func ActionIndex(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

// ActionDelete will delete current cookie that are stored by changing the expiration time
// and setting the MaxAge of the cookie to "-1"
func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
