package log

import "go.uber.org/zap"

type Logger = zap.SugaredLogger

func New() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
