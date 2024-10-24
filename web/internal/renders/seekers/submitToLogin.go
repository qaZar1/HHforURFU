package renders

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func SubmitToLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		return
	}

	username := r.FormValue("username")
	if username == "undefined" {
		logrus.Errorf("Error: invalid username")
		return
	}

	password := r.FormValue("password")
	if username == "undefined" {
		logrus.Errorf("Error: invalid password")
		return
	}

	seeker, err := api.NewApiSeekers().CheckSeeker(username)
	if err != nil {
		logrus.Errorf("Error: %s", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(seeker.PasswordHash), []byte(password)); err != nil {
		logrus.Errorf("Error: %s", err)
		return
	}

	token, err := api.NewApiSeekers().Login()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	jsoniter.NewEncoder(w).Encode(token)

	w.WriteHeader(http.StatusOK)
}
