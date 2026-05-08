# AGENTS.md - Code Review & Improvement Tracking

## Project Overview

Go backend for controlling split-flap display hardware on Raspberry Pi 4. Integrates with Spotify (now playing), Avanza stocks, and WebSocket clients. Uses Protocol Buffers for serial communication with ESP32-based modules.

**Tech Stack:** Go 1.26.1, Gin, Protocol Buffers, WebSockets, Spotify API, Avanza API

---

## Code Review Findings

### Security Issues
- [x] WebSocket `CheckOrigin` accepts all origins (`internal/websocket/main.go`)
- [x] Log endpoint (`/log`) lacks input validation/sanitization

### Code Quality Issues
- [x] ~~Duplicate `getPlayingText()` function~~ (removed from `internal/utils/text.go`)
- [ ] Mixed logging (`fmt.Println` vs structured logger in `spotify_handler.go`, `spotify.go`)
- [ ] Error return values ignored (e.g., `message_handler.go` ignores `SendMessage` error)
- [ ] Hardcoded magic numbers (`24`, `12`) instead of using config values
- [ ] Unused `validOrigins` variable in `routes.go`

### Concurrency Issues
- [ ] `SpotifyClients` and `LcdDisplays` maps accessed without mutex protection
- [ ] WebSocket broadcast modifies map during iteration (potential race condition)

### Testing & Documentation
- [ ] Very limited test coverage - only 4 test files
- [ ] Critical paths untested: Spotify integration, serial communication, state machine
- [ ] No API documentation
- [ ] Missing GoDoc comments on exported functions

### Other Issues
- [ ] 1300+ line generated protobuf file (`internal/generated/gen.go`) checked into repo
- [ ] Potential nil pointer in `handleSplitflapState` (`application.go:72-87`)
- [ ] State machine incomplete (`idleState()` always returns false, clock state not implemented)
- [ ] Serial read loop lacks buffer size limits
- [ ] No CI/CD pipeline

---

## Improvement Roadmap

### High Priority

1. **Add Comprehensive Testing**
   - Unit tests for handlers, spotify client, serial communication
   - Integration tests for critical paths
   - Use test mocks for `MessageSender` and `StocksClient` interfaces

2. **Fix Security Issues**
   - Restrict WebSocket `CheckOrigin` to known origins
   - Add input validation/sanitization to log endpoint
   - Consider adding rate limiting

3. **Fix Concurrency Issues**
   - Add mutex protection for `SpotifyClients` and `LcdDisplays` maps in `Application`
   - Review all shared state access in goroutines

4. **Improve Error Handling**
   - Don't ignore errors (e.g., `SendMessage` return value in `message_handler.go`)
   - Replace `fmt.Println` with structured logging consistently
   - Remove debugging statements

### Medium Priority

5. **Use Configuration for Magic Numbers**
   - Replace hardcoded `24`, `12` with config values
   - `ROW_LENGTH` should be derived from `cfg.GetRowLength()`

6. **Add API Documentation**
   - Document all endpoints with request/response examples
   - Consider using OpenAPI/Swagger

7. **Generated Code Management**
   - Either regenerate protobuf during build or add to `.gitignore`
   - Document the protobuf generation process

### Low Priority

8. **Code Cleanup**
   - Remove unused variables (e.g., `validOrigins` in `routes.go`)
   - Add GoDoc comments to exported functions
   - Remove commented-out code blocks

9. **CI/CD Pipeline**
   - Add Makefile targets for testing, linting, building
   - Add `golangci-lint` to CI
   - Add GitHub Actions or similar

10. **Complete Incomplete Features**
    - Implement clock state in state machine
    - Review and complete `idleState()` logic
    - Add proper nil checks in `handleSplitflapState`

---

## Build & Run

```bash
go build ./...
go run cmd/splitflap/main.go
```

## Testing

```bash
go test ./...
```

---

*Last reviewed: 2026-05-08* (WebSocket CheckOrigin fixed)
