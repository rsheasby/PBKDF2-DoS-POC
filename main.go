package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

const (
	salt       = "lesalt"
	iterations = 50000
)

func main() {
	r := chi.NewRouter()
	r.Post("/hashpls", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		password := r.FormValue("password")
		keyLen, e := strconv.Atoi(r.FormValue("keyLen"))
		if e != nil {
			w.WriteHeader(400)
			return
		}
		hash := Pbkdf2HmacSha256([]byte(password), []byte(salt), iterations, keyLen)
		w.Write([]byte(hex.EncodeToString(hash)))
	})
	fmt.Println("Server started.")
	http.ListenAndServe(":3000", r)
}
