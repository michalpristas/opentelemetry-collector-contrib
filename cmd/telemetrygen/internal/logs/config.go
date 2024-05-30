// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	"encoding/hex"

	"github.com/spf13/pflag"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/internal/common"
)

// Config describes the test scenario.
type Config struct {
	common.Config
	NumLogs        int
	Body           string
	SeverityText   string
	SeverityNumber int32
	TraceID        string
	SpanID         string
}

// Flags registers config flags.
func (c *Config) Flags(fs *pflag.FlagSet) {
	c.CommonFlags(fs)

	fs.StringVar(&c.HTTPPath, "otlp-http-url-path", "/v1/logs", "Which URL path to write to")

	fs.IntVar(&c.NumLogs, "logs", 1, "Number of logs to generate in each worker (ignored if duration is provided)")
	fs.StringVar(&c.Body, "body", "the message", "Body of the log")
	fs.StringVar(&c.SeverityText, "severity-text", "Info", "Severity text of the log")
	fs.Int32Var(&c.SeverityNumber, "severity-number", 9, "Severity number of the log, range from 1 to 24 (inclusive)")
	fs.StringVar(&c.TraceID, "trace-id", "", "TraceID of the log")
	fs.StringVar(&c.SpanID, "span-id", "", "SpanID of the log")
}

// Validate validates the test scenario parameters.
func (c *Config) Validate() error {
	if c.TraceID != "" {
		if len(c.TraceID) != 32 {
			return errInvalidTraceIDLenght
		}

		_, err := hex.DecodeString(c.TraceID)
		if err != nil {
			return errInvalidTraceID
		}
	}

	if c.SpanID != "" {
		if len(c.SpanID) != 16 {
			return errInvalidSpanIDLenght
		}
		_, err := hex.DecodeString(c.SpanID)
		if err != nil {
			return errInvalidSpanID
		}
	}

	return nil
}
