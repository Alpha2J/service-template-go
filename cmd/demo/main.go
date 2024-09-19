package main

import logger "service-template-go/internal/pkg/demo/log"

func main() {
	logger.Debug("hello from debug logger")
	logger.Debugf("hello %d", 1)
	logger.Info("hello world from logger")
	logger.Infof("hello world %d from sugaredlogger", 2)
}
