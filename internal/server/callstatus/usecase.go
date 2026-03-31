package callstatus

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type UseCase interface {
	Calculate(ctx context.Context) (models.CallStatus, error)
}
