# Web Page Analyzer

A modern web application built with Go for analyzing web pages and extracting key information about their structure, content and links.

## Project Overview

The Web Page Analyzer is a robust web application that performs comprehensive analysis of web pages. It provides detailed insights into HTML structure, link analysis, and form detection, making it useful for users.

### Key Features

- **HTML Version Detection**: Automatically identifies HTML Version
- **Page Title Extraction**: Retrieves and displays page titles
- **Heading Analysis**: Counts headings by level
- **Link Analysis**: Distinguishes between internal and external links
- **Accessibility Check**: Identifies inaccessible external links
- **Form Detection**: Detects login forms based on field patterns
- **Error Handling**: Comprehensive HTTP error reporting with user friendly messages
- **Simple UI**: Clean and responsive interface

## Architecture

The application follows Go best practices with a clean, modular architecture:

```
web-page-analyzer/
├── cmd/                   # Application entry points
├── internal/              # Private application code
│   ├── app/               # Application layer
│   ├── api/               # API layer
│   ├── services/          # Business logic layer
│   ├── models/            # Data models
│   └── utils/             # Utility functions
├── web/                   # Web assets
│   ├── static/            # Static assets (CSS, JS)
│   └── templates/         # HTML templates
├── docs/                  # Documentation
├── Dockerfile             # Container configuration
├── Makefile               # Build and deployment automation
└── README.md
```

### Technology Stack

- **Backend**: Go 1.25
- **Web Framework**: Gorilla Mux
- **HTML Parsing**: GoQuery
- **Logging**: Logrus
- **Frontend**: Vanilla JavaScript, CSS3
- **Containerization**: Docker
- **Build System**: Make

## Prerequisites

Before running this application, ensure you have:

- **Go 1.25 or higher**
- **Git** - For version control
- **Docker** - For containerized deployment
- **Make** - For build automation

### System Requirements

- **Memory**: Minimum 1GB RAM
- **Storage**: 100MB free space
- **Network**: Internet access for analyzing external URLs

## Installation

### Local Setup Using Makefile

1. **Install development tools**
   ```bash
   make install
   ```

2. **Build and run**
   ```bash
   make run
   ```

## Testing

### Run Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Logging

The application uses structured logging with Logrus:

- **Format**: JSON
- **Levels**: Debug, Info, Warn, Error, Fatal
- **Fields**: Request method, path, IP, error details