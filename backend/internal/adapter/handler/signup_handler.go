package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KeihakuOh/career-connect/internal/usecase"
)

type SignupHandler struct {
	signupUseCase usecase.SignupUseCase
}

func NewSignupHandler(signupUseCase usecase.SignupUseCase) SignupHandler {
	return SignupHandler{
		signupUseCase: signupUseCase,
	}
}

func (h *SignupHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var input usecase.SignupInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	output, err := h.signupUseCase.Signup(r.Context(), &input)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			responseWithError(w, http.StatusConflict, err.Error())
		case "invalid user type":
			responseWithError(w, http.StatusBadRequest, err.Error())
		default:
			responseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responseWithJSON(w, http.StatusCreated, output)
}

// Helper functions for HTTP responses
func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
