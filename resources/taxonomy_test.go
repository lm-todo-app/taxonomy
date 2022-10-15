package resources

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/lm-todo-app/taxonomy/db"
	"github.com/stretchr/testify/assert"
)

func TestTaxonomyRoute(t *testing.T) {
	app := createTestServer()

	tests := []struct {
		description  string
		route        string
		action       string
		expectedCode int
	}{
		{
			description:  "Incorrect url path",
			route:        "/thisisnotvalid",
			action:       "GET",
			expectedCode: 404,
		},
		{
			description:  "Test GET single taxonomy",
			route:        "/taxonomy/1",
			action:       "GET",
			expectedCode: 200,
		},
		{
			description:  "Test GET single taxonomy that does not exist",
			route:        "/taxonomy/100",
			action:       "GET",
			expectedCode: 404,
		},
		{
			description:  "Test GET all taxonomy",
			route:        "/taxonomy",
			action:       "GET",
			expectedCode: 200,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.action, test.route, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}

func createTestServer() *fiber.App {
	db.InitDatabase()
	appConfig := InitConfig()
	app := fiber.New(appConfig)
	SetupRoutes(app)
	return app
}
