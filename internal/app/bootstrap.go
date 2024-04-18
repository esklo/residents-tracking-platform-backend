package app

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/pkg/errors"
)

func (a *App) bootstrapDatabase(ctx context.Context) error {
	repo := a.serviceProvider.UserRepository()
	_, err := repo.GetByEmail(ctx, a.serviceProvider.AppConfig().AdminEmail())
	if errors.Is(err, sql.ErrNoRows) {
		user := &model.User{
			Email:     a.serviceProvider.AppConfig().AdminEmail(),
			FirstName: "Администратор",
			Role:      model.UserRoleAdmin,
		}
		if err := user.SetPassword(a.serviceProvider.AppConfig().AdminPassword()); err != nil {
			return errors.Wrap(err, "bootstrap: can not set admin user password")
		}
		if _, err := repo.Create(ctx, user); err != nil {
			return errors.Wrap(err, "bootstrap: can not get create admin user")
		}
	} else if err != nil {
		return errors.Wrap(err, "bootstrap: can not get user by email")
	}
	return nil
}
