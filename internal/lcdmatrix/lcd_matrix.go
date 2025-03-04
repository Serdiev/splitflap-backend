package lcdmatrix

import (
	"encoding/json"
	"errors"
	"fmt"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/utils"
	"splitflap-backend/pkg/fluent"
)

type LctMatrixClient struct {
	IpAddress string
}

func NewLcdMatrixClient() *LctMatrixClient {
	return &LctMatrixClient{
		IpAddress: "192.168.1.127",
	}
}

func (client *LctMatrixClient) UpdateIp(ipAddress string) {
	client.IpAddress = ipAddress
}

func (client *LctMatrixClient) SendImage(image utils.Image) error {
	payload, err := json.Marshal(image)
	if err != nil {
		return err
	}

	err = fluent.
		Post(fmt.Sprintf("http://%s/image", client.IpAddress), payload).
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
