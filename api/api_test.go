package api_test

import (
	"AdAlpha/api"
	"AdAlpha/model"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a api.Api

func TestApi_InitialiseSetsUpRoutes(t *testing.T) {
	a.Initialise()
	var routes []string
	a.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		r := fmt.Sprint(tpl, met)
		routes = append(routes, r)
		return nil
	})

	if len(routes) != 6 {
		t.Errorf("Expected number of routes %d. got %d\n", 6, len(routes))
	}

	routesContainsExpectedRoute(model.Contains(routes, "/instruction/buy[POST]"), "/instruction/buy[POST]", t)
	routesContainsExpectedRoute(model.Contains(routes, "/instruction/invest[POST]"), "/instruction/invest[POST]", t)
	routesContainsExpectedRoute(model.Contains(routes, "/instruction/sell[POST]"), "/instruction/sell[POST]", t)
	routesContainsExpectedRoute(model.Contains(routes, "/instruction/raise[POST]"), "/instruction/raise[POST]", t)
	routesContainsExpectedRoute(model.Contains(routes, "/history/investor/{id}[GET]"), "/history/investor/{id}[GET]", t)
	routesContainsExpectedRoute(model.Contains(routes, "/portfolio/investor/{id}[GET]"), "/portfolio/investor/{id}[GET]", t)

}

func routesContainsExpectedRoute(contains bool, route string, t *testing.T) {
	if !contains {
		t.Errorf("route is not initialised, route :%s", route)
	}
}
func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	a.Router.ServeHTTP(resp, r)

	return resp
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. got %d\n", expected, actual)
	}
}

func SetTestingEnvVariables() {

	err := os.Setenv("DB_NAME", "test")
	err = os.Setenv("DB_USERNAME", "test")
	err = os.Setenv("DB_PASSWORD", "test")
	err = os.Setenv("DB_PORT", "5001")
	err = os.Setenv("DB_HOST", "localhost")
	err = os.Setenv("BASE_CC", "GBP")

	if err != nil {
		fmt.Println("Issue setting test env variables")
	}
}
