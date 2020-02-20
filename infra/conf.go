package infra

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CobraInitialization() {
	viper.AutomaticEnv()

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No configuration file found")
	}
}

func LoggerConfig(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", "text", "one of text or json")
	rootCmd.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func ServerConfig(cmd *cobra.Command) {
	cmd.Flags().String("server.host", "127.0.0.1", "host on which the server should listen")
	cmd.Flags().Int("server.port", 8080, "port on which the server should listen")
	cmd.Flags().Bool("server.debug", false, "debug mode for the server")
	cmd.Flags().String("server.allowedOrigins", "*", "allowed origins for the server")
	cmd.Flags().String("server.token", "", "authorization token to use if any")
	cmd.Flags().String("jwt.salt", "", "used to sign the JWTs")
	viper.BindPFlags(cmd.Flags())
}

func DatabaseConfig(cmd *cobra.Command) {
	cmd.Flags().Bool("populate", false, "used to populate databases in order to run integration tests")
	cmd.Flags().String("mysql.host", "db", "mysql host")
	cmd.Flags().String("mysql.port", "3306", "mysql port")
	cmd.Flags().String("mysql.user", "root", "mysql user")
	cmd.Flags().String("mysql.password", "testpassword", "mysql user password")
	cmd.Flags().String("mysql.database", "clean", "mysql database")
	cmd.Flags().String("redis.host", "redis", "redis host")
	cmd.Flags().String("redis.port", "6379", "redis port")

	viper.BindPFlags(cmd.Flags())
}
