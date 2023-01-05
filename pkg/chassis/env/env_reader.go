package env

import (
	l "github.com/jgkawell/galactus/pkg/logging"

	"github.com/spf13/viper"
)

// ReadEnvironmentConfigurations reads the configuration file on the local filesystem and sets viper config
func ReadEnvironmentConfigurations(logger l.Logger, baseDirectory string) l.Error {
	logger = logger.WithField("func", "readLocalConfiguration")
	logger.Trace("NewLocalConfig")

	// grab env variables
	viper.AutomaticEnv()
	// check whether to read values.json (remote) or local.yaml (local)
	if viper.GetBool("REMOTE") {
		logger.Info("reading remote configuration")
		viper.SetConfigName("values")
	} else {
		logger.Info("reading local configuration")
		viper.SetConfigName("local")
	}
	viper.AddConfigPath(baseDirectory)

	if err := viper.ReadInConfig(); err != nil {
		return logger.WrapError(l.NewError(err, "unable to read config"))
	}

	return nil
}
