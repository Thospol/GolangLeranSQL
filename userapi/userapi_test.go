package userapi

import (
	"errors"
	"io/ioutil"
	"learnsql/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

type mockUserService struct {
	UserService
	err error
}

func (m *mockUserService) FindByID(id int) (*user.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &user.User{
		ID:        1,
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Email:     "singpor@gmail.com",
	}, nil
}

func (m *mockUserService) FindAll() ([]user.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return []user.User{
		{
			ID:        1,
			FirstName: "thospol",
			LastName:  "hemna",
			Email:     "isocare.thospol@gmail.com",
		},
		{
			ID:        2,
			FirstName: "thosporn",
			LastName:  "hemna",
			Email:     "isocare.thosporn@gmail.com",
		},
	}, nil
}

func TestGetAllUser(t *testing.T) {
	h := &Handler{
		userService: &mockUserService{},
	}
	m := mux.NewRouter()
	m.HandleFunc("/users", h.getAllUser)

	r := httptest.NewRequest("GET", "http://lvm.me:3000/users", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusOK), http.StatusText(resp.StatusCode))
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	expectJSON := `[{"id":1,"first_name":"thospol","last_name":"hemna","email":"isocare.thospol@gmail.com"},{"id":2,"first_name":"thosporn","last_name":"hemna","email":"isocare.thosporn@gmail.com"}]`
	assert.Equal(t, expectJSON, string(body))

}

func TestGetUserHandler(t *testing.T) {
	h := &Handler{
		userService: &mockUserService{},
	}
	m := mux.NewRouter()
	m.HandleFunc("/users/{id}", h.getUser)

	r := httptest.NewRequest("GET", "http://lvm.me:3000/users/1", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusOK), http.StatusText(resp.StatusCode))
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	expectJSON := `{"id":1,"first_name":"Weerasak","last_name":"Chongnguluam","email":"singpor@gmail.com"}`
	assert.Equal(t, expectJSON, string(body))
}

func TestGetUserHandlerError(t *testing.T) {
	h := &Handler{
		userService: &mockUserService{
			err: errors.New("Error"),
		},
	}
	m := mux.NewRouter()
	m.HandleFunc("/users/{id}", h.getUser)

	r := httptest.NewRequest("GET", "http://lvm.me:8000/users/1", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), http.StatusText(resp.StatusCode))
}
