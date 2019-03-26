package config

import (
	"github.com/newrelic/go-agent"
	"strconv"
	"strings"
)

func newrelicConfig() newrelic.Config {
	appName := mustGetString("NEW_RELIC_APP_NAME")
	licenceKey := mustGetString("NEW_RELIC_LICENSE_KEY")
	config := newrelic.NewConfig(appName, licenceKey)
	config.Enabled = mustGetBool("NEW_RELIC_ENABLED")
	newRelicIgnoreStatusCodes := mustGetString("NEW_RELIC_IGNORE_STATUS_CODES")
	ignoreStatusCodes := strings.Split(newRelicIgnoreStatusCodes, ",")
	config.ErrorCollector.IgnoreStatusCodes = getNewRelicErrorIgnoreCodes(ignoreStatusCodes)
	return config
}

func getNewRelicErrorIgnoreCodes(ignoreStatusCodes []string) []int {
	var errorIgnoreStatusCodes []int
	for _, statusCode := range ignoreStatusCodes {
		statusCode, err := strconv.Atoi(statusCode)
		if err == nil {
			errorIgnoreStatusCodes = append(errorIgnoreStatusCodes, statusCode)
		}
	}
	return errorIgnoreStatusCodes
}
