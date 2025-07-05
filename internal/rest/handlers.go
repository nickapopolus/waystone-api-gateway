package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nickapopolus/waystone-api-gateway/internal/data"
	"github.com/nickapopolus/waystone-api-gateway/internal/utility"
	v1 "github.com/nickapopolus/waystone-api-gateway/proto/urlservice/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

func (router *Router) PostCreateURL(w http.ResponseWriter, r *http.Request) {
	url := data.URLCreateRequest{}

	err := utility.ReadJSON(w, r, &url)
	if err != nil {
		fmt.Println(err)
		err = utility.ErrJSON(w, err)
		if err != nil {
			panic(err)
		}
		return
	}
	grpccontext, _ := context.WithTimeout(r.Context(), time.Second*1000)
	urlRequest := &v1.CreateURLRequest{
		OriginalUrl:   url.OriginalURL,
		MaxClicks:     &url.MaxClicks,
		CustomSlug:    &url.CustomSlug,
		EpirationDate: timestamppb.New(time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC)),
		IsActive:      true,
	}
	completeURL, err := router.URLClient.CreateURL(grpccontext, urlRequest)
	fmt.Println(json.Marshal(completeURL))

	err = utility.WriteJSON(w, http.StatusCreated, completeURL)
	if err != nil {
		err = utility.ErrJSON(w, err)
		if err != nil {
			panic(err)
		}
		return
	}
	return
}
