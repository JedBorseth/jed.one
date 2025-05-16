# API Documentation - api.jed.one

[![API Status](https://img.shields.io/badge/status-active-success.svg)](https://api.jed.one/health)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

This API service provides a straightforward and efficient REST API built with Go (Gin framework) and documented using a static site generated with AstroJS. Designed with simplicity and performance in mind, api.jed.one delivers reliable endpoints for your application needs.

## Features

- ⚡️ High-performance Go (Gin) backend
- 📚 Static documentation with AstroJS
- 🔄 RESTful architecture
- 📝 Comprehensive logging
- 🔍 Robust error handling

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Node.js (v16 or higher) for documentation site
- pnpm or yarn

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/api.jed.one.git

# Install Go dependencies
go mod download

# Install documentation site dependencies
cd docs
pnpm install

# Start the API server
go run main.go

# Start the documentation site (in a separate terminal)
cd docs
pnpm run dev
```

## API Documentation

Browse our documentation at [api.jed.one](https://api.jed.one)

### Base URL

```
https://api.jed.one/v1
```

## Environment Variables

```
PORT=8080
GIN_MODE=release
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please email jedborseth@gmail.com or open an issue in the GitHub repository.

## Authors

- **Jed Borseth** - _Initial work_ - [JedBorseth](https://github.com/JedBorseth)

© 2025 api.jed.one. All rights reserved.
