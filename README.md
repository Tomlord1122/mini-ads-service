# Dcard-Backend-Intern

````markdown
# Ads Management API

## Overview

This project is 2024 Dcard backend intern assignment.

## Features

- **Create Ads**: Add new advertisements with detailed targeting options.
- **List Ads**: Retrieve active ads with flexible filtering, optimized for performance.
- **Cache Integration**: Leverages Redis for caching, significantly reducing response times for frequent queries.
- **Database Optimization**: Utilizes indexed queries for efficient data retrieval.
- **Scalability**: Designed for high throughput, supporting auto-scaling and load balancing.

## Getting Started

### Prerequisites

- Go 1.15+
- PostgreSQL 12+
- Redis 6+

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Tomlord1122/Dcard-Backend-Intern.git
   ```
````

2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up the environment variables as per the `.env.example` file.

### Running the API

1. Start the PostgreSQL and Redis servers.
2. Run the following command to start the API server:
   ```bash
   go run ./api/server.go
   ```

## API Endpoints

- **POST /ads**: Create a new ad.
- **GET /ads**: List ads based on query parameters.
