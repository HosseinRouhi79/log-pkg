
***log-pkg***:
A Flexible Logging Interface for Go
log-pkg provides a unified interface for interacting with various logging packages in Go, such as zap and zerolog. This simplifies the process of switching between logging backends while maintaining consistent logging behavior in your application.

Features:

Flexible logging backend selection: Choose the logging package that best suits your needs (e.g., zap for structured logging, zerolog for performance) by specifying it in the configuration.
Consistent logging API: Use the same logging.NewLogger function and logging methods (e.g., Infof, Errorf) regardless of the underlying logging backend.
Reduced boilerplate code: Focus on your application logic without worrying about the intricate details of different logging packages.
Installation:

Bash
go get github.com/HosseinRouhi79/log-pkg
Use code with caution.
Usage:

Import the package:
```
Go
import (
  "github.com/HosseinRouhi79/log-pkg"
)
Use code with caution.
Define the logging configuration:
Go
func main() {
  cfg := config.LogConfig()
  // (Optional) Set the desired logger (defaults to zap)
  // cfg.Logger = "zerolog"
  logger := logging.NewLogger(&cfg)

  // Use the provided logging methods
  logger.Infof("Starting-%s", "test")
}
```
Use code with caution.
Default Logger:

By default, log-pkg uses zap as the logging backend.
Customizing the Logger:

You can specify the desired logging backend (e.g., "zerolog") in the config.LogConfig struct. This allows you to switch between logging packages seamlessly.
Supported Log Levels:

log-pkg provides various logging levels for different severities (e.g., Debug, Info, Warn, Error, Fatal). Refer to the documentation of the underlying logging backend (zap or zerolog) for specific details.
