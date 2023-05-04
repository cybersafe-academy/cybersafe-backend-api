package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var ExportedSettings *settings

func Config(envPrefix string, files []string) *settings {

	s := settings{source: config(envPrefix, files)}

	return &s
}

func config(envPrefix string, files []string) *viper.Viper {
	v := viper.New()

	v.SetConfigFile(files[0])
	v.AddConfigPath(".")

	if envPrefix != "" {
		v.SetEnvPrefix(envPrefix)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("could not read config file \"%s\"", files[0]))
	}

	for i := 1; i < len(files); i++ {
		f, err := os.Open(files[i])
		if err != nil {
			panic(fmt.Errorf("could not read config file \"%s\"", files[i]))
		}

		if err := v.MergeConfig(f); err != nil {
			_ = f.Close()
			fmt.Println(err.Error())
			panic(fmt.Errorf("could not merge config file \"%s\"", files[i]))
		}

		_ = f.Close()
	}

	return v
}
