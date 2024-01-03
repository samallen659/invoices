package test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func SetupDB(t testing.TB, ctx context.Context, dbPass string, dbUser string, dbName string) string {
	t.Helper()

	pgContainer, err := createDBContainer(ctx, dbPass, dbUser, dbName)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}
	ports, err := pgContainer.Container.Ports(ctx)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	hostPort := ports["5432/tcp"][0].HostPort

	err = migrateDB(dbPass, dbUser, dbName, hostPort)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	return connStr
}

func createDBContainer(ctx context.Context, dbPass string, dbUser string, dbName string) (*postgres.PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase(dbName), postgres.WithUsername(dbUser), postgres.WithPassword(dbPass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	return pgContainer, nil
}

func migrateDB(dbPass string, dbUser string, dbName string, hostPort string) error {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}
	pathToMigrationFile := filepath.Clean(filepath.Join(filepath.Dir(path), "../../db/migrations/"))

	databaseURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUser, dbPass, hostPort, dbName)
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFile), databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	log.Println("migration done")
	return nil
}

func CreateHTTPRequest(t testing.TB, method string, route string, body *[]byte) *http.Request {
	t.Helper()

	var req *http.Request
	var err error
	if body != nil {
		b := bytes.NewReader(*body)
		req, err = http.NewRequest(method, route, b)
		if err != nil {
			t.Fatalf("failed creating http request: %s", err.Error())
		}
	} else {
		req, err = http.NewRequest(method, route, nil)
		if err != nil {
			t.Fatalf("failed creating http request: %s", err.Error())
		}
	}

	return req
}

func AssertHTTPStatusCode(t testing.TB, desiredCode int, returnedCode int) {
	t.Helper()

	if returnedCode != desiredCode {
		t.Errorf("Unexpected error code, received: %d expected: %d", returnedCode, desiredCode)
	}
}
