package seeds

import (
	"behealth-api/infrastructure"
	"context"

	"firebase.google.com/go/auth"
)

// SuperAdminSeed -> super admin
type SuperAdminSeed struct {
	logger   infrastructure.Logger
	firebase *auth.Client
	env      infrastructure.Env
}

//NewSuperAdminUser -> creates admin credential
func NewSuperAdminUser(
	firebase *auth.Client,
	logger infrastructure.Logger,
	env infrastructure.Env,
) SuperAdminSeed {
	return SuperAdminSeed{
		logger:   logger,
		firebase: firebase,
		env:      env,
	}
}

func (c SuperAdminSeed) Run() {
	c.logger.Zap.Info("🌱 seeding super admin ........")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	EMAIL := c.env.AdminEmail
	PASSWORD := c.env.AdminPassword
	DISPLAYNAME := c.env.AdminDisplay

	if EMAIL != "" {
		c.logger.Zap.Info("Checking Firebase user ........")
		_, err := c.firebase.GetUserByEmail(ctx, EMAIL)

		if err != nil {
			params := (&auth.UserToCreate{}).
				Email(EMAIL).
				Password(PASSWORD).
				EmailVerified(true).
				DisplayName(DISPLAYNAME).
				Disabled(false)

			u, err := c.firebase.CreateUser(ctx, params)

			if err != nil {
				c.logger.Zap.Fatalf("Error Creating Admin in firebase: %v\n", err)
				return
			}

			//Add admin claims to the user
			claims := map[string]interface{}{"role": "admin"}
			err = c.firebase.SetCustomUserClaims(ctx, u.UID, claims)

			if err != nil {
				c.logger.Zap.Fatalf("error adding admin claim to user %v in firebase: %v\n", DISPLAYNAME, err)
				return
			}
			c.logger.Zap.Info("Admin User Created.")
			return
		}
		c.logger.Zap.Info("Admin user already exists")
	}
}
