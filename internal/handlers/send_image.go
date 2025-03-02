package handlers

import (
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

	bytes, err := img.ToBytes()
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
