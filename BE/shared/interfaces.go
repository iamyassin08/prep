package shared

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type identityManager interface {
	CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error)
	LoginUser(ctx context.Context, username string, password string) (*gocloak.JWT, error)
}
