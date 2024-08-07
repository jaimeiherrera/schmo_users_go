package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaimeiherrera/schmo_users_go/api"
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
)

type Handlers struct {
	Components *api.Components
}

func NewHandlers(components api.Components) *Handlers {
	return &Handlers{
		Components: &components,
	}
}

func (h Handlers) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`{"message": "pong"}`,
	))
}

// TODO: Must implement context.Context
func (h Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.PathValue("id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	user, err := h.Components.UserUserCase.FindUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error getting user: `, err, `"}`),
		))
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error getting user: `, err, `"}`),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.Components.UserUserCase.FindUserAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error getting users: `, err, `"}`),
		))
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error getting users: `, err, `"}`),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	userCreated, err := h.Components.UserUserCase.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error creating user: `, err, `"}`),
		))
		return
	}

	resp, err := json.Marshal(userCreated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error creating user: `, err, `"}`),
		))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.PathValue("id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	if err := h.Components.UserUserCase.DeleteUser(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error deleting user: `, err, `"}`),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`{"message": "User deleted"}`,
	))
}

func (h Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.PathValue("id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	userUpdated, err := h.Components.UserUserCase.UpdateUser(userID, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error updating user: `, err, `"}`),
		))
		return
	}

	resp, err := json.Marshal(userUpdated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error updating user: `, err, `"}`),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
