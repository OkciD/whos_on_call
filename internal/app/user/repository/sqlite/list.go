package sqlite

import (
	"context"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	dbModels "github.com/OkciD/whos_on_call/internal/app/models/db"
)

func (r *Repository) List(ctx context.Context) ([]appModels.User, error) {
	dbUsers := make([]dbModels.User, 0, 4)

	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM users")
	if err != nil {
		return nil, fmt.Errorf("select all users failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		dbUser := dbModels.User{}
		if err := rows.Scan(&dbUser.ID, &dbUser.Name); err != nil {
			return nil, fmt.Errorf("failed to scan db users: %w", err)
		}
		dbUsers = append(dbUsers, dbUser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in list users query: %w", err)
	}

	appUsers := make([]appModels.User, 0, len(dbUsers))
	for _, dbUser := range dbUsers {
		appUser := dbUser.ToAppModel()
		appUsers = append(appUsers, *appUser)
	}

	return appUsers, nil
}
