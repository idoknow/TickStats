package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/routes"
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
)

var router *gin.Engine
var os_names = []string{"windows", "linux", "macos", "android", "ios"}

func init() {
	gin.SetMode(gin.TestMode)
	router = routes.SetupRouter()
}

func TestAddMetric(t *testing.T) {
	w := httptest.NewRecorder()

	var metric models.BasicMetricInput

	// random data
	metric.MetricsData = map[string]interface{}{
		"usage_cnt": rand.Intn(100),
		"os_name":   os_names[rand.Intn(len(os_names))],
	}

	// /api/metric/9c0a6ca3 with POST method, metric as JSON
	jsonData, _ := json.Marshal(metric)
	req, _ := http.NewRequest("POST", "/api/metric/9c0a6ca3", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
