package http

import (
	"github.com/tanopwan/rolele"
	"net/http"
	"strings"
)

func getActionFromHeader(header string) []string {
	return strings.FieldsFunc(header, func(c rune) bool {
		return c == ';' || c == ' '
	})
}

func GetRoleleMiddleware(allowActions []rolele.Action, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		xActions := req.Header.Get("X-Actions")
		actions := getActionFromHeader(xActions)

		allow := rolele.ApplyActions(allowActions, actions)
		if allow {
			h.ServeHTTP(w, req)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	})
}
