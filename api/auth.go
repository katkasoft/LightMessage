package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

var jwtKey = []byte("MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCCqKz18AQ3i9vxW4uaO3UlrhMQMPS6Lnk2QjiPuODnwEA==")

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(4320 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	password, err := HashPassword(req.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}

	_, err = DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", req.Username, req.Email, password)
	if err != nil {
		log.Printf("Error while writing in DB: %v", err)
		http.Error(w, "User already exists or DB error", http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(req.Username)
	if err != nil {
		log.Printf("Error while generating JWT token: %v", err)
		http.Error(w, "Error whule generating JWT token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(4320 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})

	log.Printf("Login: Username:%s Password:%s", req.Username, req.Password)
}

func init() {
	log.Println("Auth API initialised")
}
