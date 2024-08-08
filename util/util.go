package util

import (
	"flag"
	"os"

	"github.com/VijayVPatil/Helping-Hand/model"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {

	Logger = *logrus.New()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		// Customizing delimiters
	})

	Logger.SetOutput(os.Stdout)
}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level(debug , error , info, warning)")

	flag.Parse()
	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelWarn:
		Logger.SetLevel(logrus.WarnLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func Log(logLevel, packageLevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameter != nil {
			Logger.Debugf("package level: %s, functionName : %s , message: %v , parameter : %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debugf("package level: %s, functionName : %s , message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelWarn:
		if parameter != nil {
			Logger.Warnf("package level: %s, functionName : %s , message: %v , parameter : %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("package level: %s, functionName : %s , message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("package level: %s, functionName : %s , message: %v , parameter : %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("package level: %s, functionName : %s , message: %v\n", packageLevel, functionName, message)
		}
	default:
		if parameter != nil {
			Logger.Infof("package level: %s, functionName : %s , message: %v , parameter : %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("package level: %s, functionName : %s , message: %v\n", packageLevel, functionName, message)
		}
	}
}
