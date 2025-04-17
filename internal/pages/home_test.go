package pages

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	engine := html.New("../../views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
        t.Fatalf("Ошибка при парсинге HTML: %v", err)
    }

	ul := doc.Find("ul.hor-list")
    if ul.Length() == 0 {
        t.Fatalf("ul с классом hor-list не найден")
    }

	var foundCategories []string
	ul.Find("li").Each(func(i int, s *goquery.Selection) {
		foundCategories = append(foundCategories, s.Text())
	})

	for _, c := range categories {
		found := false
		for _, actual := range foundCategories {
			if actual == c {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Элемент <li> с текстом %q не найден", c)
		}
	}
}
