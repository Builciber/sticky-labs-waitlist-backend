package main

import (
	"database/sql"
	"internal/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	sessionSecret         string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	DB                    *database.Queries
	domain                string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error establishing database connection")
	}
	dbQueries := database.New(db)
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	waitlistApiMux := chi.NewRouter()
	waitlistApiMux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{""},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	cfg := apiConfig{
		TwitterConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		sessionSecret:         os.Getenv("SESSION_SECRET"),
		DB:                    dbQueries,
		domain:                os.Getenv("DOMAIN"),
	}
	waitlistMux := chi.NewRouter()
	waitlistMux.Get("/", cfg.handlerHome)
	waitlistMux.Get("/signup", cfg.handlerServeSignupPage)
	waitlistMux.Get("/profile", cfg.handlerServeProfilePage)
	mux.Handle("/", http.FileServer(http.Dir("./sticky")))
	mux.Handle("/Sticky-Labs/*", http.FileServer(http.Dir(".")))
	mux.Handle("/sticky/*", http.FileServer(http.Dir(".")))
	// 1. Register Twitter login and callback handlers
	oauth1Config := &oauth1.Config{
		ConsumerKey:    cfg.TwitterConsumerKey,
		ConsumerSecret: cfg.TwitterConsumerSecret,
		CallbackURL:    "https://walrus-app-bl8a2.ondigitalocean.app/twitter/callback",
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
	waitlistApiMux.Get("/logout", cfg.handlerLogout)
	waitlistApiMux.Post("/signup", cfg.handlerWaitlistSignup)
	waitlistApiMux.Get("/leaderboard", cfg.handlerWaitlistLeaderboard)
	mux.Handle("/twitter/login", twitter.LoginHandler(oauth1Config, nil))
	mux.Handle("/twitter/callback", twitter.CallbackHandler(oauth1Config, cfg.issueSessionToken(), nil))
	mux.Handle("/session", cfg.handlerIssueSessionToken())

	mux.Mount("/waitlist/", waitlistMux)
	mux.Mount("/waitlist/api/", waitlistApiMux)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("Started server on localhost at port 8080")
	err = server.ListenAndServe()
	log.Fatal(err)
}
