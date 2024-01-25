package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestULIDAPI(t *testing.T) {
	app := fiber.New()
	ulidAPI(app)
	t.Run("Test ULID API", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/ulid", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var ulidResponse Response
		err = json.NewDecoder(resp.Body).Decode(&ulidResponse)
		assert.NoError(t, err)
		assert.True(t, ulidResponse.Success)
		assert.NotNil(t, ulidResponse.Data.(map[string]interface{})["ulid"])
	})
}
