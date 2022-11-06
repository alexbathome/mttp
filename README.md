# mttp
### HTTP Server with in-built Prometheus Metrics Middleware

[![Go Report Card](https://goreportcard.com/badge/github.com/alexbathome/mttp/mttp)](https://goreportcard.com/report/github.com/alexbathome/mttp/mttp)

The MTTP Library is a simple HTTP Router, that is constructed in a clear, and declarative manner by leveraging the builder pattern.

The MTTP Server allows you to very easily inject common prometheus metrics to a HTTP route, simply by using the `WithMetrics()` method on your Server.
