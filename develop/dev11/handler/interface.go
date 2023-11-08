package handler

import "net/http"

type RESTHandler interface {
	MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc
	CreateEvent(w http.ResponseWriter, r *http.Request)
}
