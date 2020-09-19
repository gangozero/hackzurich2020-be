package server

import (
	"log"

	"github.com/gangozero/hackzurich2020-be/generated/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) UserLoginHandler() user.LoginHandler {
	return user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
		resp, err := s.Login(params.HTTPRequest.Context(), params.Data.Email.String())
		if err != nil {
			log.Printf("[Login] Error: %s", err.Error())
			return user.NewLoginDefault(500)
		}
		if resp == nil {
			return user.NewLoginDefault(404)
		}
		return user.NewLoginOK().WithPayload(resp)
	})
}
