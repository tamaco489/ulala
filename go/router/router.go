package router

import (
	"net/http"
	"os"
	"strings"

	"github.com/miyabiii1210/ulala/go/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DEFAULT_API_KEY string = "ABCDEFG123456789"
	DEFAULT_ORIGIN  string = "http://localhost:3000"
)

var (
	apiKey string
	origin string
)

func init() {
	if os.Getenv("API_KEY") == "" {
		apiKey = DEFAULT_API_KEY
	}
	if os.Getenv("ORIGIN") == "" {
		origin = DEFAULT_ORIGIN
	}
}

func headerValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		authHeader := req.Header.Get("Authorization")
		if (authHeader == "") || (!strings.HasPrefix(authHeader, "Bearer ")) || (strings.Split(authHeader, " ")[1] != apiKey) {
			return echo.NewHTTPError(http.StatusUnauthorized, "401: Unauthorized")
		}

		// fmt.Printf("[debug] %#v\n", req.Header)
		return next(c)
	}
}

func NewRouter(
	uc controller.IUserController,
	ac controller.IAuthController,
	mc controller.IMovieController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
				http.MethodOptions,
			},

			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderAccessControlAllowHeaders,
				"Authorization",
			},

			AllowOrigins: []string{
				origin,
			},

			AllowCredentials: true, // ブラウザにcookieを設定する際に必要
		},
	))

	e.Use(headerValidationMiddleware)

	// auth
	e.POST("/signup", ac.SignUp)
	e.POST("/signin", ac.SignIn)
	e.POST("/signout", ac.SignOut)

	// user
	e.GET("/users", uc.GetUsers)
	e.GET("/users/:uid", uc.GetUserByUID)
	e.PATCH("/users/:uid", uc.UpdateUser)
	e.DELETE("/users/:uid", uc.DeleteUser)

	// movie
	e.GET("/movies/categories", mc.GetMovieCategories) // typeテーブルの情報を全て取得
	e.GET("/movies/type", mc.GetMovieListByTypeID)     // 指定したtype_idの情報を取得, query parameter: type?id=1
	e.GET("/movies/:movie_id", mc.GetMovie)            // 指定したmovie_idの動画全ての情報を取得

	return e
}
