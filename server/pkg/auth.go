package server

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/idtoken"
)

// This should be taken from https://console.cloud.google.com/apis/credentials
var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")

var rootHtmlTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sign in with Google</title>
    <script src="https://accounts.google.com/gsi/client" async></script>
    <style>
        body {
            background-color: #EAEEF3;
        }

        .container {
            padding: 0px 20px;
        }

        .container h3 {
            font-family: sans-serif;
        }
    </style>
</head>
<body>
    <div class="container">
        <h3>Sign in with Google to Chat</h3>
        <div
            id="g_id_onload"
            data-client_id="{{.clientID}}"
            data-login_uri="{{.callbackURL}}">
        </div>
        <div
            class="g_id_signin"
            data-type="standard"
            data-theme="filled_blue"
            data-text="sign_in_with"
            data-shape="rectangular"
            data-width="200"
            data-logo_alignment="left">
        </div>
    </div>
</body>
</html>

`))

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	if len(googleClientID) == 0 {
		http.Error(w, "Set GOOGLE_CLIENT_ID env var", http.StatusBadRequest)
		return
	}

	err := rootHtmlTemplate.Execute(w, map[string]string{
		"callbackURL": "http://localhost:8030/google_callback",
		"clientID":    googleClientID,
	})
	if err != nil {
		panic(err)
	}
}

func CallbackHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got to callback handler!")
	defer req.Body.Close()

	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// The following steps follow
	// https://developers.google.com/identity/gsi/web/guides/verify-google-id-token
	//
	// 1. Verify the CSRF token, which uses the double-submit-cookie pattern and
	//    is added both as a cookie value and post body.
	token, err := req.Cookie("g_csrf_token")
	if err != nil {
		http.Error(w, "token not found", http.StatusBadRequest)
		return
	}

	bodyToken := req.FormValue("g_csrf_token")
	if token.Value != bodyToken {
		http.Error(w, "token mismatch", http.StatusBadRequest)
	}

	// 2. Verify the ID token, which is returned in the `credential` field.
	//    We use the idtoken package for this. `audience` is our client ID.
	ctx := context.Background()
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		panic(err)
	}
	// credential string is the OIDC id token jwt value that can be parsed
	credential := req.FormValue("credential")

	payload, err := validator.Validate(ctx, credential, googleClientID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//TODO(steve): verify that the user is actually in a google group TBD
	// before we set the cookie and proceed
	for key, value := range req.Form {
		fmt.Println("key:", key, "value:", value)
	}

	// 3. Once the token's validity is confirmed, we can use the user identifying
	//    information in the Google ID token.
	for k, v := range payload.Claims {
		fmt.Printf("%v: %v\n", k, v)
	}

	// 4. (steve) set the JWT token value as a new cookie
	c := &http.Cookie{
		Name:     "credential",
		Value:    credential,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   req.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)

	// 5. tells browsers that the response should be exposed to the front-end JavaScript code.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
