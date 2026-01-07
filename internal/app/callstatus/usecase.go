package callstatus

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type UseCase interface {
	Calculate(ctx context.Context) (models.CallStatus, error)
}
