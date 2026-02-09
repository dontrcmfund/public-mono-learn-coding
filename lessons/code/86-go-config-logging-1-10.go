package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
GO CONFIG + LOGGING (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/86-go-config-logging-1-10.go
2) Try env overrides:
   - APP_PORT=9090 go run lessons/code/86-go-config-logging-1-10.go
   - APP_DEBUG=true go run lessons/code/86-go-config-logging-1-10.go

Extra context:
- lessons/notes/173-go-config-first-principles.md
- lessons/notes/174-go-observability-first-principles.md
*/

// LESSON 1: Runtime config struct
// Why this matters: explicit config keeps runtime behavior understandable.
type Config struct {
	AppName      string
	Port         int
	Debug        bool
	ReadTimeoutS int
}

// LESSON 2: Helper to read env with default
// Why this matters: defaults keep local onboarding simple.
func envOrDefault(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

// LESSON 3: Parse typed env values
// Why this matters: config values must be validated, not assumed.
func parseIntEnv(key string, fallback int) (int, error) {
	raw := envOrDefault(key, strconv.Itoa(fallback))
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be int, got %q", key, raw)
	}
	return n, nil
}

func parseBoolEnv(key string, fallback bool) (bool, error) {
	rawDefault := "false"
	if fallback {
		rawDefault = "true"
	}
	raw := strings.ToLower(envOrDefault(key, rawDefault))
	switch raw {
	case "true", "1", "yes":
		return true, nil
	case "false", "0", "no":
		return false, nil
	default:
		return false, fmt.Errorf("%s must be bool, got %q", key, raw)
	}
}

// LESSON 4: Central config loader
// Why this matters: one place for startup validation.
func loadConfig() (Config, error) {
	port, err := parseIntEnv("APP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}
	timeout, err := parseIntEnv("APP_READ_TIMEOUT_S", 5)
	if err != nil {
		return Config{}, err
	}
	debug, err := parseBoolEnv("APP_DEBUG", false)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{
		AppName:      envOrDefault("APP_NAME", "go-learning-service"),
		Port:         port,
		Debug:        debug,
		ReadTimeoutS: timeout,
	}
	if cfg.Port <= 0 {
		return Config{}, fmt.Errorf("APP_PORT must be positive")
	}
	return cfg, nil
}

// LESSON 5: Structured log helper
// Why this matters: stable keys make logs searchable and machine-friendly.
func logEvent(level string, msg string, fields map[string]any) {
	parts := []string{
		"ts=" + time.Now().Format(time.RFC3339),
		"level=" + level,
		"msg=" + strconv.Quote(msg),
	}
	for k, v := range fields {
		parts = append(parts, k+"="+fmt.Sprint(v))
	}
	log.Println(strings.Join(parts, " "))
}

// LESSON 6: Startup log pattern
// Why this matters: startup visibility reduces deployment confusion.

// LESSON 7: Error log pattern
// Why this matters: failures should include actionable context.

// LESSON 8: Debug-conditional logs
// Why this matters: useful detail in dev without noisy prod logs.
func maybeDebug(cfg Config, msg string, fields map[string]any) {
	if !cfg.Debug {
		return
	}
	logEvent("DEBUG", msg, fields)
}

// LESSON 9: Simulated request handling logs
// Why this matters: observability should exist around normal flows too.
func handleRequest(cfg Config, method string, path string) {
	start := time.Now()
	maybeDebug(cfg, "request.received", map[string]any{"method": method, "path": path})
	time.Sleep(10 * time.Millisecond)
	durationMS := time.Since(start).Milliseconds()
	logEvent("INFO", "request.completed", map[string]any{
		"method":      method,
		"path":        path,
		"status_code": 200,
		"duration_ms": durationMS,
	})
}

// LESSON 10: End-to-end startup demonstration
// Why this matters: config + logging create operational confidence.
func main() {
	cfg, err := loadConfig()
	if err != nil {
		logEvent("ERROR", "startup.config_invalid", map[string]any{"error": err})
		return
	}

	logEvent("INFO", "startup.ready", map[string]any{
		"app":            cfg.AppName,
		"port":           cfg.Port,
		"debug":          cfg.Debug,
		"read_timeout_s": cfg.ReadTimeoutS,
	})

	handleRequest(cfg, "GET", "/health")
	handleRequest(cfg, "POST", "/tasks")
}

// End of Go Config + Logging 1-10
