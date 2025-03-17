# Contributing to mcp-kafka

Thank you for your interest in contributing to mcp-kafka! This document provides guidelines and instructions for contributing to the project.

## Development Setup

1. Fork and clone the repository:
   ```bash
   git clone https://github.com/YOUR_USERNAME/mcp-kafka.git
   cd mcp-kafka
   ```

2. Set up your development environment:
   - Ensure you have Go 1.24 or higher installed
   - Install required dependencies:
     ```bash
     go mod download
     ```

3. Create a new branch for your feature:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Guidelines

### Code Style

- Follow the standard Go formatting guidelines
- Run `go fmt` before committing your changes
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and single-purpose

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting a PR:
  ```bash
  go test ./...
  ```
- Maintain or improve test coverage

### Commit Messages

- Use clear and descriptive commit messages
- Follow the conventional commits format:
  ```
  type(scope): description

  [optional body]
  [optional footer]
  ```
- Types: feat, fix, docs, style, refactor, test, chore

## Pull Request Process

1. Update your branch with the latest changes from main:
   ```bash
   git fetch origin
   git rebase origin/main
   ```

2. Push your changes:
   ```bash
   git push origin feature/your-feature-name
   ```

3. Create a Pull Request:
   - Provide a clear description of your changes
   - Reference any related issues
   - Include screenshots or GIFs for UI changes
   - List any breaking changes

4. Respond to review comments and make requested changes

## Feature Requests

- Use the GitHub Issues feature to request new features
- Provide a clear description of the feature
- Explain the use case and benefits
- Include any relevant examples

## Bug Reports

When reporting bugs, please include:
- A clear description of the issue
- Steps to reproduce
- Expected behavior
- Actual behavior
- Environment details (OS, Go version, etc.)
- Any relevant logs or error messages

## Documentation

- Update README.md for significant changes
- Add inline documentation for new functions
- Update any relevant configuration examples
- Include usage examples for new features

## Questions?

Feel free to open an issue for any questions or concerns. We're here to help!

## License

By contributing, you agree that your contributions will be licensed under the project's MIT License. 