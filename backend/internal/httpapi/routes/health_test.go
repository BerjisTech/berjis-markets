package routes

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestRegisterHealth(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.New()
    group := router.Group("/api")
    RegisterHealth(group)

    req := httptest.NewRequest(http.MethodGet, "/api/healthz", nil)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    if resp.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", resp.Code)
    }
}
