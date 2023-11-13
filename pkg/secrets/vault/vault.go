package vault

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/secrets"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type (
	wrapper struct {
		client *vault.Client
	}
)

func New() secrets.Client {
	return &wrapper{}
}

func (c *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	vaultAddress := config.GetString("vault.url")

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress(vaultAddress),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return err
	}

	// authenticate with a root token (insecure)
	// TODO: change to use app role login
	/*
		resp, err := client.Auth.AppRoleLogin(
			ctx,
			schema.AppRoleLoginRequest{
				RoleId:   os.Getenv("MY_APPROLE_ROLE_ID"),
				SecretId: os.Getenv("MY_APPROLE_SECRET_ID"),
			},
			vault.WithMountPath("my/approle/path"), // optional, defaults to "approle"
		)
	*/
	if err := client.SetToken("myroot"); err != nil {
		return err
	}
	c.client = client

	return nil
}

func (c *wrapper) Get(ctx context.Context, key string) (string, error) {
	path, attribute := splitKey(key)
	s, err := c.client.Secrets.KvV2Read(ctx, path, vault.WithMountPath("secret"))
	if err != nil {
		return "", err
	}
	value, ok := s.Data.Data[attribute].(string)
	if !ok {
		return "", fmt.Errorf("unable to read secret")
	}
	return value, nil
}

func (c *wrapper) Set(ctx context.Context, key string, value string) error {
	path, attribute := splitKey(key)
	_, err := c.client.Secrets.KvV2Write(ctx, path, schema.KvV2WriteRequest{
		Data: map[string]any{
			attribute: value,
		}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *wrapper) Delete(ctx context.Context, key string) error {
	_, err := c.client.Secrets.KvV2Delete(ctx, key, vault.WithMountPath("secret"))
	if err != nil {
		return err
	}
	return nil
}

// HELPERS

func splitKey(key string) (path string, attribute string) {
	segments := strings.Split(key, "/")
	if len(segments) == 0 {
		return "", ""
	}
	// path is everything but the last segment
	path = strings.Join(segments[:len(segments)-1], "/")
	// attribute is the last segment
	attribute = segments[len(segments)-1]
	return path, attribute
}
