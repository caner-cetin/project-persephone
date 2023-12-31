// Package core contains all the routers and the helper functions for routers.
package core

import (
	"github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

var HandlerFunc = func() http.Handler {
	router := chi.NewRouter()
	//
	// PRE-SET MIDDLEWARES
	//
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// please do not fiddle with middleware order.
	router.Use(AssignServer)
	router.Use(AssignWriter)
	router.Use(AssignRequest)
	router.Use(AssignLogger)
	router.Use(AssignValidator)
	router.Use(AssignQueryBuilder)

	//
	// POSTGRES CONNECTION POOL INITIALIZATION
	//
	db, err := GetPgPool()
	if err != nil {
		panic(err)
	}
	router.Use(AssignDB(db))
	// MOUNT YOUR ROUTERS HERE.
	router.Route("/api", func(r chi.Router) {
		r.Mount("/user", NewUserHandler())
		r.Mount("/world", NewCityHandler())
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:3000/api/swagger/doc.json"),
		))
	})
	return otelhttp.NewHandler(router, "server", otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
		return operation + " " + r.URL.Path
	}))
}

// types.go contains the most common types used in the routers. if a struct is used only in one handler, or
// strictly related to the handler, it will be defined in the handler file, usually at the top of the handler.
//
// please follow the same convention.

// Server , please dont assign Server fields without middlewares, it will cause a panic if you dont know what you are doing.
//
// Middleware initialization order:
//
//	AssignLogger("/signup")
//	AssignDB("users", client)
//	AssignTracer("/signup", "USER_SIGNUP")
type Server struct {
	// DB
	DB *pgxpool.Pool
	// Logger contains a SugaredLogger that is prefixed to the specific route specified in middleware initialization.
	//
	// Example:
	//
	//		router.Route("/core", func(r chi.Router) {
	//			r.Use(AssignLogger("/signup"))
	//			r.Use(AssignDB("users", client))
	//		})
	//
	// Will create a terminal-logger with the "/signup" as the prefix.
	//
	// So for any logging, you can access them like so:
	//
	//		func (s Server) CreateUser() {
	//			s.Logger.Info("User created")
	//		}
	Logger *slog.Logger
	// Tracer contains a Jaeger tracer that is used to trace the specific route specified in middleware initialization.
	//
	// Example:
	//
	//		router.Route("/core", func(r chi.Router) {
	//			r.Use(AssignTracer("/signup", "USER_SIGNUP"))
	//		})
	//
	// Will create a tracer with the "/signup" as an attribute that is sticked to the span, and USER_SIGNUP as the operation name.
	Tracer trace.Span
	// Validator contains a validator that is used to validate any struct with an example format of:
	// 		type UserSignupRequest struct {
	//			Email    string `json:"email" validate:"email" binding:"required"`
	//			Username string `json:"username" binding:"required"`
	//			Password string `json:"password" validate:"password" binding:"required"`
	//			Test     bool   `json:"test" validate:"boolean"`
	//		}
	//
	// Example:
	//
	//		var signUpForm UserSignupRequest
	//		if err := DecodeJSONBody(r, &signUpForm); err != nil {
	//			s.LogError(err, w, http.StatusBadRequest)
	//			return
	//		}
	//		err := s.Validator.Struct(signUpForm)
	//
	//

	Validator   *validator.Validate
	Writer      http.ResponseWriter
	Request     *http.Request
	StmtBuilder squirrel.StatementBuilderType
}

type serverKey string

const ServerKeyString serverKey = "server"

type JWTFields struct {
	// UUID is the id representing the user, and the key of the user in the database bucket "users"
	UUID string `json:"uuid"`
	// Expires shows the expiration datetime in unix.
	Expires float64 `json:"exp"`
	Role    string  `json:"role"`
	Token   string  `json:"token"`
	Status  string  `json:"status"`
}

const (
	JWTUUIDKey    = "uuid"
	JWTExpiresKey = "exp"
	JWTRoleKey    = "role"
	JWTStatusKey  = "status"
)

const (
	tokenDurationLogin   = 30 * time.Minute
	tokenDurationSession = 24 * time.Hour
	tokenDurationRefresh = 7 * 24 * time.Hour
)

const (
	tokenStatusActive       = "ACTIVE"
	tokenStatusWaitingLogin = "WAITING_LOGIN"
	tokenStatusRefresh      = "REFRESH"
)

const (
	JWT_ENCRYPT_ALGORITHM = "HS256"
	JWT_ENCRYPT_KEY       = "f18e90e9b71cf9af78022ef12350602d"
)

// for db
const (
	DBConnSSLDisabled = "disable"
	DBConnSSLEnabled  = "require"
)
const (
	DBPassword = "caner"
	DBUser     = "caner"
	DBPort     = 5432
	DBHost     = "localhost"
	DefaultDB  = "persephone"
)

const (
	AllowedUserEmailUpdateInterval = time.Hour * 24 * 7
	AllowedUsernameUpdateInterval  = time.Hour * 24 * 90
)

const (
	CountryTable                 = "countries"
	CountryIDDBField             = "id"
	CountryNameDBField           = "name"
	CountryISO3DBField           = "iso3"
	CountryISO2DBField           = "iso2"
	CountryPhoneCodeDBField      = "phonecode"
	CountryCapitalDBField        = "capital"
	CountryCurrencyDBField       = "currency"
	CountryCurrencyNameDBField   = "currency_name"
	CountryCurrencySymbolDBField = "currency_symbol"
	CountryTLDDField             = "tld"
	CountryNativeDBField         = "native"
	CountryRegionDBField         = "region"
	CountrySubregionDBField      = "subregion"
	CountryLatitudeDBField       = "latitude"
	CountryLongitudeDBField      = "longitude"
	CountryEmojiDBField          = "emoji"
	CountryEmojiUDBField         = "emojiU"
	CountryCreatedAtDBField      = "created_at"
	CountryUpdatedAtDBField      = "updated_at"
	CountryTranslationsDBField   = "translations"
	CountryTimezoneIDDBField     = "timezone_id"
	CountryNumericCodeDBField    = "numeric_code"
)

const (
	StateTable              = "states"
	StateIDDBField          = "id"
	StateNameDBField        = "name"
	StateCountryIDDBField   = "country_id"
	StateCountryCodeDBField = "country_code"
	StateTypeDBField        = "type"
	StateLatitudeDBField    = "latitude"
	StateLongitudeDBField   = "longitude"
	StateCreatedAtDBField   = "created_at"
	StateUpdatedAtDBField   = "updated_at"
	StateFlagDBField        = "flag"
	StateWikiDataIDDBField  = "wiki_data_id"
)

const (
	TimezoneTableName            = "timezones"
	TimezoneIDDBField            = "id"
	TimezoneZoneNameDBField      = "zone_name"
	TimezoneGmtOffsetDBField     = "gmt_offset"
	TimezoneGmtOffsetNameDBField = "gmt_offset_name"
	TimezoneAbbreviationDBField  = "abbreviation"
	TimezoneTzNameDBField        = "tz_name"
)

const (
	CityTable              = "cities"
	CityIDDBField          = "id"
	CityNameDBField        = "name"
	CityStateIDDBField     = "state_id"
	CityStateCodeDBField   = "state_code"
	CityCountryIDDBField   = "country_id"
	CityCountryCodeDBField = "country_code"
	CityLatitudeDBField    = "latitude"
	CityLongitudeDBField   = "longitude"
	CityCreatedAtDBField   = "created_at"
	CityUpdatedAtDBField   = "updated_at"
	CityWikiDataIDDBField  = "wiki_data_id"
)
const (
	RestaurantsTable              = "places"
	RestaurantIDDBField           = "id"
	RestaurantNameDBField         = "name"
	RestaurantCuisineDBField      = "cuisine"
	RestaurantOpeningHoursDBField = "opening_hours"
	RestaurantPhoneDBField        = "phone"
	RestaurantWebsiteDBField      = "website"
	RestaurantFullAddressDBField  = "full_address"
	RestaurantHouseNumberDBField  = "house_number"
	RestaurantStreetDBField       = "street"
	RestaurantPostcodeDBField     = "postcode"
	RestaurantCityDBField         = "city"
	RestaurantCountryDBField      = "country"
	RestaurantStateDBField        = "state"
	RestaurantLatitudeDBField     = "latitude"
	RestaurantLongitudeDBField    = "longitude"
)

type StmtBuilders interface {
	ToSql() (string, []interface{}, error)
}
