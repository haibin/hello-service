package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "SALES : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(logger); err != nil {
		logger.Print(err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	var cfg struct {
		Web struct {
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}
	fmt.Println(cfg)
	return nil
}