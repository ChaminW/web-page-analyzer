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
├── internal/              
│   ├── app/               # Application layer
│   ├── api/               # API layer
│   ├── services/          # Business logic layer
│   ├── models/            # Data models
│   ├── middleware/        # Middleware layer
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
- **Metrics**: Prometheus
- **Profiling**: pprof
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

## Observability

### Logging

The application uses structured logging with Logrus:

- **Format**: JSON
- **Levels**: Debug, Info, Warn, Error, Fatal
- **Fields**: Request method, path, IP, error details

### Metrics

- Request count and response times
- Error rates and types
- Analysis duration statistics

## Error Handling

The application provides comprehensive error handling:

- **HTTP Errors**: Detailed status codes with user-friendly descriptions
- **Network Errors**: Connection timeout and network failure handling
- **Validation Errors**: URL format and input validation
- **Graceful Degradation**: Continues operation even with partial failures

## Technical Challenges & Solutions

### 1. HTML Version Detection
- **Challenge**: Accurately determining HTML versions without explicit DOCTYPE declarations
- **Solution**: Implemented a multi-layered detection system

### 2. Concurrent Link Analysis
- **Challenge**: Performance bottleneck during external link validation
- **Solution**: Implemented parallel processing using Goroutines for concurrent link checking

## Future Improvements

### 1. Performance Optimizations
- Implement caching layer for frequently analyzed URLs
- Add connection pooling for HTTP requests
- Optimize HTML parsing for large documents
- Implement request rate limiting

### 2. Feature Enhancements
- Add support for JavaScript rendered content analysis
- Enable custom timeout configurations
- Add batch URL processing capability

### 3. Security Improvements
- Implement advanced input validation
- Add rate limiting per IP
- Enhance error handling for security-sensitive information
- Add support for authentication and authorization

### 4. Monitoring & Analytics
- Implement detailed performance metrics
- Add real-time analysis dashboard