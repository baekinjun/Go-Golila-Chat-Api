package route

import (
	"net/http"
	chat "chat_api/chat"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// type TemplateRenderer struct {
// 	templates *template.Template
// }

// // Render renders a template document
// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check success")
}

func Init() *echo.Echo {
	e := echo.New()
	e.Debug = true
	renderer := &TemplateRenderer{
	 	templates: template.Must(template.ParseGlob("*.html")),
	 }
	 e.Renderer = renderer
	 Recover
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/chat/:room_id/", chat.Start)
	e.GET("/health/", healthCheck)
	return e
}
