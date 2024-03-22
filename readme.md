# Trace-Torch

## Overview
Trace-Torch is a logging library designed to facilitate easy and efficient tracking of application events. It provides a multi-level logging system that allows developers to capture detailed information about their application's behavior in real-time.

## Features
- **Multi-Level Logging**: Supports various levels of logging, including Debug, Info, and Error.
- **Customizable Output**: Allows users to define custom outputs for log messages.
- **Easy Integration**: Designed for simple integration with Go applications.

## Getting Started
To use Trace-Torch in your Go application, start by importing the library:

```go
import "github.com/smark-ark/Trace-Torch/pocketlog"
```


Create a new logger instance with the desired threshold level:
```go
logger := pocketlog.New(pocketlog.LevelInfo, os.Stdout)
```
Now you can use the logger to log messages at different levels.
```go
logger.Infof("Payment of amount %v from account %v received", amount, accountID)
logger.Errorf("Error processing request: %v", err)
```


## Contributing
Contributions to Trace-Torch are welcome! Please feel free to submit issues, pull requests, or enhancements to improve the library.