package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Response recorder
	w := httptest.NewRecorder()

	// Request processing
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	articleList = tmpArticleList
}
