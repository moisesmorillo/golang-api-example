package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/moisesmorillo/golang-api-example/enums"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	c := SetupControllerCase(http.MethodGet, enums.RouterHealthCheckPath, nil)

	if assert.NoError(t, HealthCheck(c.context)) {
		assert.Equal(t, http.StatusOK, c.Res.Code)
		assert.Equal(t, `{"message":"Available!"}`, strings.TrimSpace(c.Res.Body.String()))
	}
}

// common code for handler testing
type ControllerCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupControllerCase(method string, url string, body io.Reader) ControllerCase {
	path := fmt.Sprintf(enums.RouterGlobalPath+"%s", url)

	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return ControllerCase{req, res, c}
}
