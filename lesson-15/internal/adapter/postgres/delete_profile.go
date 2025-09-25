package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/transaction"
)

func (p *Postgres) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	ctx, span := tracer.Start(ctx, "adapter postgres DeleteProfile")
	defer span.End()

	const sql = `UPDATE profile SET deleted_at = NOW()
                     WHERE id = $1`

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}

		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
