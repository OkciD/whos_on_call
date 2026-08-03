package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (c *apiClient) GetUser(ctx context.Context) (*models.User, error) {
	resp, err := c.genClient.GetUserWithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from api: %w", err)
	}

	statusCode := resp.StatusCode()

	if statusCode == 200 {
		return resp.JSON200.ToAppModel(), nil
	} else {
		return nil, mapper.RespToError(statusCode, *resp.JSONDefault)
	}
}
