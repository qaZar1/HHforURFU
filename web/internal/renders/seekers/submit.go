package renders

import (
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		return
	}

	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	resume := r.FormValue("resume")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logrus.Errorf("Error: invalid hash")
		return
	}

	seeker := models.SeekerWithoutID{
		FirstName:    name,
		Username:     username,
		PasswordHash: string(hash),
		Resume:       resume,
	}

	ok, err := api.NewApiSeekers().RegisterSeeker(seeker)
	if err != nil || !ok {
		http.Error(w, "Пользователь уже существует! Попробуйте другой username!", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}
