package auth

import (
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func Routes(client *ent.Client, jwt *JWTProvider) http.Handler {
	r := chi.NewRouter()
	r.Get("/testAuthDONOTUSE/{username}", TestAuthDONOTUSE(client, jwt))
	return r
}

func TestAuthDONOTUSE(client *ent.Client, jwt *JWTProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var staticUser *ent.User
		var err error
		username := chi.URLParam(r, "username")
		adminViewerCtx := viewer.NewSystemAdminContext(ctx)
		staticUser, err = client.User.Query().Where(user.Username(username)).First(adminViewerCtx)
		if err != nil {
			permsCreate := client.UserPermissions.Create()
			if r.URL.Query().Get("admin") == "true" {
				permsCreate.SetAdmin(true)
			}
			perms, err := permsCreate.Save(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create static user permissions")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			staticUser, err = client.User.Create().
				SetUsername(username).
				SetEmail(fmt.Sprintf("%s@notarealemail.com", username)).
				SetUserPermissions(perms).
				Save(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create static user")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		perms := staticUser.QueryUserPermissions().FirstX(adminViewerCtx)
		token, err := jwt.CreateToken(
			staticUser.ID.String(),
			staticUser.Username,
			permissions.From(perms).StringSlice(),
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create token")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token.String(),
			Expires: token.Claims().ExpiresAt.Time,
		})
		w.Header().Add("X-Api-Token", token.String())
		w.Header().Add("X-Api-Expires", token.Claims().ExpiresAt.Format(time.RFC3339))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, staticUser)
	}
}
