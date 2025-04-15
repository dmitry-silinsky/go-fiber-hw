package pages

import (
	"fmt"
	"io"
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

	assert.NoError(t, err, "Неожиданная ошибка выполнения запроса")
	assert.Equal(
		t,
		http.StatusOK,
		resp.StatusCode,
		fmt.Sprintf("Ожидаемый статус запроса: %d, фактический: %d", http.StatusOK, resp.StatusCode),
	)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Ошибка чтения тела ответа: %v", err)
	}

	assert.Equal(t, "Start", string(body))
}
