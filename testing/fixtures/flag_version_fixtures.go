package fixtures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FeatureFlagVersionTestData() []Test {
	feature := "version"
	return []Test{
		shouldPrintCurrentVersion(feature),
	}
}

func shouldPrintCurrentVersion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintCurrentVersion",
		CmdArgs: []string{"--version"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "1build version 1.1.2")
		},
	}
}
