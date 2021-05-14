package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/koetsuserizawa/wire_example/blogservice"
	"github.com/koetsuserizawa/wire_example/userservice"
)

type UserServiceHandler struct {
	s userservice.UserServiceInterface
}

func NewUserServiceHandler(us userservice.UserServiceInterface) *UserServiceHandler {
	return &UserServiceHandler{
		s: us,
	}
}

func (h *UserServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		//Register user.
		b, err := parseRequestBody(r)
		id := b["id"].(string)
		name := b["name"].(string)
		if err != nil || name == "" || id == "" {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		errIns := h.s.Register(id, name)
		if errIns != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

type BlogServiceHandler struct {
	bs blogservice.BlogServiceInterface
	us userservice.UserServiceInterface
}

func NewBlogServiceHandler(bs blogservice.BlogServiceInterface, us userservice.UserServiceInterface) *BlogServiceHandler {
	return &BlogServiceHandler{
		bs: bs,
		us: us,
	}
}

func (h *BlogServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := parseRequestBody(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.us.Get(b["id"].(string))
	if u == nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		//Get all articles of an user.
		arts, err := h.bs.Read(u.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println(arts)
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		//Post an article.
		err := h.bs.Write(u.ID, b["text"].(string))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func parseRequestBody(r *http.Request) (map[string]interface{}, error) {
	b, err := ioutil.ReadAll(r.Body)
	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	return j, err
}
