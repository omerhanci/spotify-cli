package auth

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

const redirectURI = "http://localhost:8080/callback"

var (
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

type AuthClient struct {
	token  *oauth2.Token
	auth   *spotifyauth.Authenticator
	Client *spotify.Client
}

func NewAuthClient() *AuthClient {
	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState))
	tokenRecord := viper.Get("token")

	if tokenRecord != nil {
		tokenRecordMap := tokenRecord.(map[string]interface{})
		token := &oauth2.Token{
			AccessToken:  tokenRecordMap["access_token"].(string),
			TokenType:    tokenRecordMap["token_type"].(string),
			RefreshToken: tokenRecordMap["refresh_token"].(string),
		}
		client := GetAuthClient(auth, token, nil)

		// test for current token validation
		user, err := client.CurrentUser(context.Background())
		if err != nil && err.(spotify.Error).Status == 401 {
			fmt.Println("Session expired..")
		} else {
			fmt.Println("Welcome!", user.DisplayName)
			return &AuthClient{
				token:  token,
				auth:   auth,
				Client: client,
			}
		}
	}

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	authClient := &AuthClient{
		auth: auth,
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Token(r.Context(), state, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, state)
		}
		viper.Set("token", token)
		viper.WriteConfig()
		authClient.token = token

		// use the token to get an authenticated client
		client := GetAuthClient(auth, token, r)
		ch <- client
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	authClient.Client = <-ch
	authClient.Client.CurrentUser(context.Background())
	return authClient
}

func GetAuthClient(auth *spotifyauth.Authenticator, token *oauth2.Token, r *http.Request) *spotify.Client {
	if r == nil {
		return spotify.New(auth.Client(context.Background(), token))
	}
	return spotify.New(auth.Client(r.Context(), token))
}
