package arangodb

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/apito-cms/arangodb-lite/types"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtCredentials struct {
	jwt            string `json:"jwt"`
	mustChangePass bool   `json:"must_change_pass"`
}

func (c *Connection) authenticate(user, pass string) error {
	creds, err := json.Marshal(&credentials{
		Username: user,
		Password: pass,
	})
	if err != nil {
		return err
	}

	body, err := c.post("_open/auth", creds)
	if err != nil {
		return err
	}

	// Assign token to connection for future use
	auth := new(types.Auth)
	if err = json.Unmarshal(body, auth); err != nil {
		return err
	}
	c.token = fmt.Sprintf("bearer %s", auth.Jwt)
	c.header.Set("Authorization", c.token)

	if c.config.DebugMode {
		log.Infof("connected to: %s using token: %s", c.host, auth.Jwt)
	}

	return nil
}
