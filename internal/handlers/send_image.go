package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/utils"
	"splitflap-backend/pkg/fluent"
)

func (a *Application) SendImage(url string) error {
	img, err := utils.ConvertUrlToImage(url)
	if err != nil {
		logger.Error().Err(err).Msg("failed to convert url to image")
		return err
	}

	bytes, err := json.Marshal(img)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get image bytes")
		return err
	}

	err = fluent.Post(fmt.Sprintf("http://%s:8080/image", a.ExternalLCDDisplayIP), bytes).
		OnSuccess(func(bytes []byte) error {
			return nil
		}).
		OnError(func(bytes []byte) error {
			return errors.New(string(bytes))
		}).
		Execute()

	if err != nil {
		logger.Error().Err(err).Msg("failed to send image")
		return err
	}

	return nil
}

type ImageRequest struct {
	Image [][]Color `json:"image"`
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}
