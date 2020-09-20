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

func (s *Server) UserGetProductListHandler() user.GetProductListHandler {
	return user.GetProductListHandlerFunc(func(params user.GetProductListParams, principal interface{}) middleware.Responder {
		userID, err := getUserFromBearer(principal)
		if err != nil {
			log.Printf("[GetProductList] Error: %s", err.Error())
			return user.NewGetProductListDefault(401)
		}

		resp, err := s.GetProductList(params.HTTPRequest.Context(), userID)
		if err != nil {
			log.Printf("[GetProductList] Error: %s", err.Error())
			return user.NewGetProductListDefault(500)
		}
		if resp == nil {
			return user.NewGetProductListDefault(404)
		}
		return user.NewGetProductListOK().WithPayload(resp)
	})
}

func (s *Server) UserDeleteProductHandler() user.DeleteProductHandler {
	return user.DeleteProductHandlerFunc(func(params user.DeleteProductParams, principal interface{}) middleware.Responder {
		userID, err := getUserFromBearer(principal)
		if err != nil {
			log.Printf("[DeleteProduct] Error: %s", err.Error())
			return user.NewDeleteProductDefault(401)
		}

		err = s.DeleteProduct(params.HTTPRequest.Context(), userID, params.ID)
		if err != nil {
			log.Printf("[DeleteProduct] Error: %s", err.Error())
			return user.NewDeleteProductDefault(500)
		}

		return user.NewDeleteProductNoContent()
	})
}

func (s *Server) UserGetRecipeMatchListHandler() user.GetRecipeMatchListHandler {
	return user.GetRecipeMatchListHandlerFunc(func(params user.GetRecipeMatchListParams, principal interface{}) middleware.Responder {
		userID, err := getUserFromBearer(principal)
		if err != nil {
			log.Printf("[GetRecipeMatchList] Error: %s", err.Error())
			return user.NewGetRecipeMatchListDefault(401)
		}

		resp, err := s.GetRecipeMatchList(params.HTTPRequest.Context(), userID)
		if err != nil {
			log.Printf("[GetRecipeMatchList] Error: %s", err.Error())
			return user.NewGetRecipeMatchListDefault(500)
		}
		return user.NewGetRecipeMatchListOK().WithPayload(resp)
	})
}

func (s *Server) UserGetRecipeDetailsHandler() user.GetRecipeDetailsHandler {
	return user.GetRecipeDetailsHandlerFunc(func(params user.GetRecipeDetailsParams, principal interface{}) middleware.Responder {
		userID, err := getUserFromBearer(principal)
		if err != nil {
			log.Printf("[GetRecipeDetails] Error: %s", err.Error())
			return user.NewGetRecipeDetailsDefault(401)
		}

		resp, err := s.GetRecipeDetails(params.HTTPRequest.Context(), userID, params.ID)
		if err != nil {
			log.Printf("[GetRecipeDetails] Error: %s", err.Error())
			return user.NewGetRecipeDetailsDefault(500)
		}
		if resp == nil {
			return user.NewGetRecipeDetailsDefault(404)
		}
		return user.NewGetRecipeDetailsOK().WithPayload(resp)
	})
}
