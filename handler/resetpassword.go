package handler

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type EmailForm struct {
	Email    string
	Password string
	Errors   map[string]string
}

func (h *Handler) forgotPassword(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	form := EmailForm{}
	if err := h.templates.ExecuteTemplate(rw, "reset-password.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.decoder.Decode(&form, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if form.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		res := h.db.Debug().Exec("update users set password = ? where email = ?", password, form.Email)
		if res.RowsAffected > 0 {

			http.RedirectHandler("/login", http.StatusTemporaryRedirect)
		}

	}

}
