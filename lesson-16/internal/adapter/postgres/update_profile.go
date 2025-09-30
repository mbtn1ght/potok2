package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/transaction"
)

func (p *Postgres) UpdateProfile(ctx context.Context, profile domain.Profile) error {
	ctx, span := tracer.Start(ctx, "adapter postgres UpdateProfile")
	defer span.End()

	const sql = `UPDATE profile SET name = $1, age = $2, contacts = $3, updated_at = NOW()
                     WHERE id = $4`

	contacts, err := json.Marshal(profile.Contacts)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	args := []any{
		profile.Name,
		profile.Age,
		contacts,
		profile.ID,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err = txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}

		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
