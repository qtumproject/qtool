package log

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus/hooks/writer"
)

type Option func(logger *logrus.Logger) error

func WithDebugLevel(debug bool) Option {
	return func(logger *logrus.Logger) error {
		if debug {
			logger.SetLevel(logrus.DebugLevel)
			logger.SetReportCaller(true)
			logger.Formatter.(*logrus.TextFormatter).FullTimestamp = true
		}
		return nil
	}
}

func WithWriter(output io.Writer) Option {
	return func(logger *logrus.Logger) error {
		logger.SetOutput(output)
		return nil
	}
}

func WithFiles(outputFile string, errorFile string) Option {
	return func(logger *logrus.Logger) error {
		if _, err := os.Stat(outputFile); err == nil {
			os.Remove(outputFile)
		}
		if _, err := os.Stat(errorFile); err == nil {
			os.Remove(errorFile)
		}
		pathMap := lfshook.PathMap{
			logrus.ErrorLevel: errorFile,
			logrus.DebugLevel: outputFile,
			logrus.InfoLevel:  outputFile,
			logrus.WarnLevel:  outputFile,
		}
		logger.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{
				TimestampFormat: "02-01-2006 15:04:05",
				CallerPrettyfier: func(f *runtime.Frame) (string, string) {
					return fmt.Sprintf("%s ", formatFilePath(f.Function)), fmt.Sprintf(" %s:%d ", formatFilePath(f.File), f.Line)
				},
			},
		))
		return nil
	}
}
