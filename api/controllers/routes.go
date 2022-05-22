package controllers

import "github.com/dy850078/gogoAPI/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
}
