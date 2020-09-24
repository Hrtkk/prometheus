package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Response ...
type Response struct {
	Message string `json:"message"`
}

// Jwks ...
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys ...
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// Product ...
type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	Product{ID: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security."},
	Product{ID: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience."},
	Product{ID: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go baack 65 millio years in the past and ride a T-rex."},
	Product{ID: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{ID: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery."},
	Product{ID: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR."},
}

// func main() {
// 	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
// 		// will fill in next
// 		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 			// Varify 'aud' claim
// 			aud := "https://dev-j3cu426s.us.auth0.com/"
// 			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
// 			print(checkAud)
// 			if !checkAud {
// 				return token, errors.New("invalid audience")
// 			}
// 			// Verify 'iss' Claim
// 			iss := "http://localhost:8080/"
// 			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
// 			if !checkIss {
// 				return token, errors.New("invalid issuer")
// 			}

// 			cert, err := getPemCert(token)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
// 			return result, nil
// 		},
// 		SigningMethod: jwt.SigningMethodRS256,
// 	})

// 	r := mux.NewRouter()

// 	r.Handle("/", handle).Methods("GET")
// 	// r.Handle("/", http.FileServer(http.Dir("./views/")))
// 	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
// 	r.Handle("/products", jwtMiddleware.Handler(ProductsHandler)).Methods("GET")
// 	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")

// 	// For dev only - Set up CORS so React client can consume our API
// 	corsWrapper := cors.New(cors.Options{
// 		AllowedMethods: []string{"GET", "POST"},
// 		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
// 	})

// 	http.ListenAndServe(":8080", corsWrapper.Handler(r))
// }

var handle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hii there, how are you..!!\n")
})

// ProductsHandler ...
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// AddFeedbackHandler ...
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("http://localhost:8080/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
