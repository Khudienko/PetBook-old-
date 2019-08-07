package handler

import (
	"PetBook/pkg/tokens"
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func (c *Controller) MyPageGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(context.Get(r, "email"))
		// if user has pet with filled fields
		tokens.GenerateHTML(w, nil, "mypage")
		// else generate "cabinetPet"
	}
}
