package configs

import "github.com/spf13/viper"

func initConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("configs")
	config.SetConfigName("settings.dev.yml") // 配置文件的名称
	config.SetConfigType("yaml")             // 配置文件的扩展名，这里除了json还可以有yaml等格式
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}
