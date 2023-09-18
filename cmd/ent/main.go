package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"entgo.io/ent/dialect"

	// postgres driver.
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/otakakot/go-orm-s/gen/ent"
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		panic(err)
	}

	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	ctx := context.Background()

	tx, err := client.Tx(ctx)
	if err != nil {
		panic(err)
	}

	now := time.Now().UTC()

	user := ent.EntUser{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Deleted:   false,
	}

	userName := ent.EntUserName{
		ID:        uuid.New(),
		UserID:    user.ID,
		Value:     uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		Deleted:   false,
	}

	if _, err := tx.EntUser.Create().
		SetID(user.ID).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		SetDeleted(user.Deleted).
		Save(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			slog.Error("failed to ollback")
		}
		panic(err)
	}

	if _, err := tx.EntUserName.Create().
		SetID(userName.ID).
		SetUserID(userName.UserID).
		SetValue(userName.Value).
		SetCreatedAt(userName.CreatedAt).
		SetUpdatedAt(userName.UpdatedAt).
		SetDeleted(userName.Deleted).
		Save(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			slog.Error("failed to ollback")
		}
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	if us, err := client.EntUser.Get(ctx, user.ID); err == nil {
		slog.Info(fmt.Sprintf("user %+v", us))
	} else {
		panic(err)
	}

	if un, err := client.EntUserName.Query().Where(entusername.UserIDEQ(user.ID)).First(ctx); err == nil {
		slog.Info(fmt.Sprintf("user name %+v", un))
	} else {
		panic(err)
	}
}
