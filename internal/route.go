package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	eSwagger "github.com/swaggo/echo-swagger"

	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/auth"
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/router"

	_ "github.com/dimaskiddo/go-whatsapp-multidevice-rest/docs"

	ctlAuth "github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/auth"
	typAuth "github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/auth/types"

	ctlIndex "github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/index"
	ctlWhatsApp "github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/whatsapp"
)

func Routes(e *echo.Echo) {
	// Route for Index
	// ---------------------------------------------
	e.GET(router.BaseURL, ctlIndex.Index)
	e.GET(router.BaseURL+"/docs/*", eSwagger.WrapHandler)

	// Route for Auth
	// ---------------------------------------------
	e.GET(router.BaseURL+"/auth", ctlAuth.Auth, auth.BasicAuth())

	// Route for WhatsApp
	// ---------------------------------------------
	authJWTConfig := middleware.JWTConfig{
		Claims:     &typAuth.AuthJWTClaims{},
		SigningKey: []byte(auth.AuthJWTSecret),
	}

	e.POST(router.BaseURL+"/login", ctlWhatsApp.Login, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/text", ctlWhatsApp.SendText, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/location", ctlWhatsApp.SendLocation, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/contact", ctlWhatsApp.SendContact, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/link", ctlWhatsApp.SendLink, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/document", ctlWhatsApp.SendDocument, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/image", ctlWhatsApp.SendImage, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/audio", ctlWhatsApp.SendAudio, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/video", ctlWhatsApp.SendVideo, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/send/sticker", ctlWhatsApp.SendSticker, middleware.JWTWithConfig(authJWTConfig))
	e.POST(router.BaseURL+"/logout", ctlWhatsApp.Logout, middleware.JWTWithConfig(authJWTConfig))
}
