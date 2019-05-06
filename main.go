package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func main() {
	// Get config path
	logger := log.New()
	if len(os.Args) <= 1 {
		logger.Println(fmt.Errorf("config path argument not provided"))
		os.Exit(1)
	}

	// Load app configs
	err := loadAppConfig(os.Args[1])
	if err != nil {
		logger.Println(fmt.Errorf("failed to load config file. Error: %s", err.Error()))
		os.Exit(2)
	}


	logPath := viper.GetString("log_path")
	if len(logPath) <= 0 {
		logger.Println(fmt.Errorf("unable to find 'log_path configuration"))
		os.Exit(3)
	}
	// Open a file
	out, err := os.OpenFile(viper.GetString("log_path"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logger.Println(fmt.Errorf("error opening file: %v", err))
		os.Exit(4)
	}

	// close it
	defer out.Close()

	initLog(out)

	err = DoWork()
	if err != nil {
		log.Error("Processing error: ", err.Error())
		os.Exit(5)
	}
}

// Work entry point
func DoWork() error {
	return fmt.Errorf("nothing implemented yet")
}

//Loads config settings via viper lib
func loadAppConfig(configPathArg string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPathArg)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error loading config file: %s", err)
	}

	// Can do a config test right her if desired
	if !viper.IsSet("my_setting") {
		return fmt.Errorf("error 'my_setting' not set in config")
	}
	return nil
}


// Log Initialization
func initLog(out io.Writer) {
	// Use default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)

	log.SetOutput(out)
}