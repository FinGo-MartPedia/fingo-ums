package helpers

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Logger Initialized Using Logrus")

	return log
}
