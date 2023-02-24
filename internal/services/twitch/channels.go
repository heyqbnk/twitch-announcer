package twitch

import (
	"context"
	"fmt"

	"github.com/qbnk/twitch-announcer/pkg/twitch/helix"

	"github.com/qbnk/twitch-announcer/internal/models"
)

// GetChannel returns information about specified Twitch channel by its
// identifier.
func (s *Service) GetChannel(ctx context.Context, channelID int) (helix.Channel, error) {
	accessToken, err := s.getAccessToken(ctx)
	if err != nil {
		return helix.Channel{}, fmt.Errorf("%w: %v", models.ErrNoAccessToken, err)
	}

	channel, err := s.api.GetChannel(ctx, accessToken, channelID)
	if err != nil {
		return helix.Channel{}, fmt.Errorf("get channel from API: %v", err)
	}

	return channel, nil
}
