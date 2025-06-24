package rest

import (
	"fmt"
	"github.com/nickapopolus/waystone-api-gateway/internal/utility"
	"net/http"
)

func (router *Router) PostCreateURL(w http.ResponseWriter, r *http.Request) {
	url := router.Models.URLCreateRequest

	err := utility.ReadJSON(w, r, &url)
	if err != nil {
		fmt.Println(err)
		err = utility.ErrJSON(w, err)
		if err != nil {
			panic(err)
		}
		return
	}
	err = utility.WriteJSON(w, http.StatusCreated, url)
	if err != nil {
		err = utility.ErrJSON(w, err)
		if err != nil {
			panic(err)
		}
		return
	}
	return
}
