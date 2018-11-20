package userapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"learnsql/user"

	"github.com/gorilla/mux"
)

type UserService interface {
	FindByID(id int) (*user.User, error)
	FindAll() ([]user.User, error)
	Update(u *user.User) error
	Delete(u *user.User) error
	Insert(u *user.User) error
}

type Handler struct {
	userService UserService
}

//เติม utilitie func WriteError
func WriteError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, "users:"+err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

//เติม utilitie func interface{} ไม่มี method โยนtypeอะไรไปก็ได้
func WriteJson(w http.ResponseWriter, value interface{}) bool {
	b, err := json.Marshal(value)

	if WriteError(w, err) {
		return true
	}
	fmt.Fprintf(w, "%s", b)
	return false
}

func (h *Handler) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAll()
	if WriteError(w, err) {
		return
	}
	WriteJson(w, users)

}

//ชื่อ func กลายเป็น method get users
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //map key [string]string
	log.Println("ID", vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if WriteError(w, err) {
		return
	}

	user, err := h.userService.FindByID(id)
	if WriteError(w, err) {
		return
	}
	WriteJson(w, user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAll()
	if WriteError(w, err) {
		return
	}
	WriteJson(w, users)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAll()
	if WriteError(w, err) {
		return
	}
	WriteJson(w, users)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAll()
	if WriteError(w, err) {
		return
	}
	WriteJson(w, users)
}

func StartServer(addr string, db *sql.DB) error {
	r := mux.NewRouter() //
	h := &Handler{       //สร้าง typeไหม่ที่ implement method FindByID
		userService: &user.Service{
			DB: db,
		},
	} //สร้าง handler
	r.HandleFunc("/users/{id}", h.getUser).Methods("GET")
	r.HandleFunc("/users", h.getAllUser).Methods("GET")
	r.HandleFunc("users/{id}", h.updateUser).Methods("POST", "PUT")
	r.HandleFunc("users/{id}", h.deleteUser).Methods("DELETE")
	r.HandleFunc("users/", h.deleteUser).Methods("POST")
	return http.ListenAndServe(addr, r)
}
