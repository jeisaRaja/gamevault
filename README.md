# GameVault
GameVault is a RESTful API designed to manage and store information about video games, including their titles, release years, genres, platforms, developers, publishers, and pricing. The API supports creating, retrieving, and validating game entries with a flexible and efficient data model.

## Features

- Game Management: Create, read, and validate games.
- Validation: Validates game fields such as title, genres, and platforms.
- Flexible Data Model: Supports multiple platforms and genres in a string-based format.
- Docker Support: Includes Docker and Docker Compose configurations for easy setup and deployment.

## Project Structure
```
├── cmd
│   └── gamevault          # Main application logic
├── docs                   # Documentation
├── internal
│   ├── data               # Data models and database logic
│   └── validator          # Input validation logic
├── migrations             # Database migrations
├── .gitignore             # Git ignore file
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose configuration
├── .env                   # Environment variables
├── go.mod                 # Go module dependencies
├── go.sum                 # Go checksum dependencies
├── run_compose.sh         # Script to run Docker Compose
├── run_gamevault.sh       # Script to start the API
├── stop_compose.sh        # Script to stop Docker Compose
└── README.md              # Project documentation
```

