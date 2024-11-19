package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var (
	// The signing key for the token.
	// signingKey = []byte("ce62Pso3DXdbvKeiY2_7CUUhXg6tRlk4zt8M958ZX1aDxJsYvV08vVNX4AR3z4iw")
	// signingKey = []byte("54yP_4is7prda-qCwzka1")

	// The issuer of our token.
	issuer = "https://insouslide.us.auth0.com/"

	// The audience of our token.
	audience = []string{"https://api.insou.ai"}

	// Our token must be signed using this data.
	// keyFunc = func(ctx context.Context) (interface{}, error) {
	// 	return signingKey, nil
	// }

	// We want this struct to be filled in with
	// our custom claims from the token.
	customClaims = func() validator.CustomClaims {
		return &CustomClaimsExample{}
	}
)

// CustomClaimsExample contains custom data we want from the token.
type CustomClaimsExample struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	ShouldReject bool   `json:"shouldReject,omitempty"`
}

// Validate errors out if `ShouldReject` is true.
func (c *CustomClaimsExample) Validate(ctx context.Context) error {
	if c.ShouldReject {
		return errors.New("should reject was set to true")
	}
	return nil
}

//

type AuthMiddleware struct {
	jwtValidator *validator.Validator
	middleware   *jwtmiddleware.JWTMiddleware
}

func NewAuthMiddleware() *AuthMiddleware {
	issuerURL, _ := url.Parse(issuer)

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuer,
		audience,
		validator.WithCustomClaims(customClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)

	if err != nil {
		panic("jwt init error")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		if err != nil {

			fmt.Println("========== jwt err:", err)

			res, _ := json.Marshal(map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  fmt.Sprintf("Unauthorized:%s", err.Error()),
			})
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(res)
		}
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return &AuthMiddleware{
		jwtValidator: jwtValidator,
		middleware:   middleware,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handler := m.middleware.CheckJWT(next)
		handler.ServeHTTP(w, r)
		// next(w, r)
	}
}
