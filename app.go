package admin

import (
	"crypto/rsa"
	"golang.org/x/net/context"

	jwtgo "github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// App holds information about application configuration
type App struct {
	projectID      string
	serviceAccount string
	jwtConfig      *jwt.Config
	privateKey     *rsa.PrivateKey
	databaseURL    string
}

// AppOptions is the firebase app options for initialize app
type AppOptions struct {
	ProjectID      string
	ServiceAccount []byte
	DatabaseURL    string
}

// InitializeApp initializes firebase application with options
func InitializeApp(options AppOptions) (*App, error) {
	var err error

	app := App{
		projectID:   options.ProjectID,
		databaseURL: options.DatabaseURL,
	}

	if options.ServiceAccount != nil {
		app.jwtConfig, err = google.JWTConfigFromJSON(options.ServiceAccount, scopes...)
		if err != nil {
			return nil, err
		}
		app.privateKey, err = jwtgo.ParseRSAPrivateKeyFromPEM(app.jwtConfig.PrivateKey)
		if err != nil {
			return nil, err
		}
	}

	return &app, nil
}

// Auth creates new Auth instance
// each instance has the save firebase app instance
// but difference public keys instance
// better create only one instance
func (app *App) Auth(context context.Context) (*Auth, error) {
	return newAuth(app, context)
}

// Database creates new Database instance
func (app *App) Database() *Database {
	return newDatabase(app)
}
