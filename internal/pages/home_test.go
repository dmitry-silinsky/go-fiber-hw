package pages

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	app := fiber.New()

	NewHandler(app)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	assert.Equal(t, "Start", string(body))
}
