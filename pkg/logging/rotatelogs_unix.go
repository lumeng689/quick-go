// +build !windows

package logging

import (
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join("logs", "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}
