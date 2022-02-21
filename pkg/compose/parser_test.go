package compose_test

import (
	"github.com/kaiaverkvist/dployr/pkg/compose"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	composeFile = `
version: '3'
services:
  app:
    container_name: test-data
    build: .
    environment:
      - SERVER_PORT=${SERVER_PORT}
      - AUTO_TLS=${AUTO_TLS}
      - FRIENDLY_LOGGING=${FRIENDLY_LOGGING}
      - DOMAIN=${DOMAIN}
      - OPERATOR_USER_ID=${OPERATOR_USER_ID}
      - GITHUB_CLIENT_ID=${GITHUB_CLIENT_ID}
      - GITHUB_CLIENT_SECRET=${GITHUB_CLIENT_SECRET}
      - GITHUB_REDIRECT_URL=${GITHUB_REDIRECT_URL}
      - SESSION_SECRET=${SESSION_SECRET}
      - DATABASE_HOST=${DB_HOST}
      - DATABASE_USER=${DB_USER}
      - DATABASE_NAME=${DB_NAME}
      - DATABASE_PASSWORD=${DB_PASSWORD}
    ports:
      - ${SERVER_PORT}:${SERVER_PORT}
    restart: on-failure
    volumes:
      - api:/usr/src/app/
    depends_on:
      - fullstack-postgres
    networks:
      - fullstack
`
)

var (
	expectedEnvs = []string{
		"SERVER_PORT",
		"AUTO_TLS",
		"FRIENDLY_LOGGING",
		"DOMAIN",
		"OPERATOR_USER_ID",
		"GITHUB_CLIENT_ID",
		"GITHUB_CLIENT_SECRET",
		"GITHUB_REDIRECT_URL",
		"SESSION_SECRET",
		"DB_HOST",
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
	}
)

func TestParseEnvironmentVariables(t *testing.T) {
	envs := compose.GetExpectedEnvironmentVariables(composeFile)

	assert.Len(t, envs, 13)
	assert.Equal(t, expectedEnvs, envs)
}
