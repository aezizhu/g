# Contributing to g

Thank you for your interest in contributing to `g`! This document provides guidelines for contributing to the project.

## Development Setup

1. **Prerequisites**
   - Go 1.21 or higher
   - Git
   - OpenWrt build environment (optional, for testing)

2. **Clone and Build**
   ```bash
   git clone https://github.com/aezizhu/g.git
   cd g
   go build ./cmd/g
   ```

3. **Run Tests**
   ```bash
   go test ./...
   ```

## Code Style

- Follow standard Go conventions and `gofmt`
- Add tests for new functionality
- Update documentation for user-facing changes
- Keep commits focused and atomic

## Security Considerations

- All new command execution paths must use argv arrays (no shell)
- New providers must validate inputs and handle timeouts
- Policy engine changes require security review
- LuCI endpoints must validate user input

## Submitting Changes

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Reporting Issues

- Use GitHub Issues for bugs and feature requests
- Include steps to reproduce for bugs
- Provide system details (OS, Go version, OpenWrt version if applicable)
- Include relevant log output

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
