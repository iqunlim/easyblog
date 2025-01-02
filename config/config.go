package config

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

// Why is this not in the os package by default
func GetEnvWithDefault(key, fallback string) string {
	v, e := os.LookupEnv(key)
	if !e {
		return fallback
	}
	return v
}

func GetEnvWithWarning(key string, fallback string) string {
	v, e := os.LookupEnv(key)
	if !e {
		log.Printf("Warning: You should set %s as the default key is \"%s\"", key, fallback)
		v = fallback
	}
	return v
}

func GetProjectRoot() string {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := path.Dir(e)
	log.Println("Project Root Path:", path)
	return path
}



var (
	
	ProjectRoot = GetProjectRoot()

	Domain = GetEnvWithWarning("DOMAIN", "localhost")

	// CORS Config.
	// AllowOirigins will be set to the domain specified in 
	// the DOMAIN Environment variable
	CORSConfig = cors.Config{
		AllowOrigins: []string{Domain},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Type", "x-requested-with"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}

	// The location of the sqlite database. test.db by default.
	// In the docker container, the location will probably be 
	// /app/db/sqlite.db or ./db/sqlite.db
	DatabaseFileLocation = GetEnvWithDefault("DATABASE_LOC", "test.db")
	

	// TODO: Configure this
	// You can find it in controller/route.go r.SetTrustedProxies(nil) 
	//TrustedProxies = []string{"localhost"}

	// GORM Configuration options
	GORMConfig = &gorm.Config{} 

	// The JWT Key. 
	// Required to be set! .... or it would be if it didnt break tests
	// TODO: Fix this
	JWTKey = GetEnvWithWarning("JWT_KEY", "default") //RequiredEnv("JWT_KEY") 

	// The JWT Realm for the JWT token module.
	// Defaults to "EasyBlogJWT"
	JWTRealm = GetEnvWithDefault("JWT_REALM", "EasyBlogJWT")

	//Defines if allowing the register page to be accessed
	// REGISTER is the variable. Defaults to "false".
	// Set to "true" to allow registrations.
	// Currently, it's best to run this once and then never again
	RegisterAllowed = GetEnvWithDefault("REGISTER", "true")

)
