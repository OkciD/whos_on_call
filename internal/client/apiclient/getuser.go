package apiclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (c *apiClient) GetUser(ctx context.Context) (*models.User, error) {
	resp, err := c.genClient.GetUserWithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from api: %w", err)
	}

	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		return resp.JSON200.ToAppModel(), nil
	}
	return nil, mapper.RespToError(statusCode, *resp.JSONDefault)
}
