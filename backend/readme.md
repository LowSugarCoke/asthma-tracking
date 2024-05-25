
## Folder Structure 
Recommended by [Go - The Ultimate Folder Structure](https://gist.github.com/ayoubzulfiqar/9f1a34049332711fddd4d4b2bfd46096)

```
main/                       # project root
│
├── api/
│   └── routes.go       # OpenAPI/Swagger specs
│
├── cmd/
│   └── server/
│       └── main.go         # Application entry point
│
├── internal/
│   │
│   ├── app/
│   │   └── server.go      # Application server logic
│   │   └── config.go   # Configuration logic  
│   │   └── database.go    # Database connection logic
│   │
│   ├── dependencies/
│   │   └── dependencies.go    # Dependencies struct needed by the server
│   │
│   ├── handlers/
│   │   ├── users.go    # User request handlers
│   │   ├── validations.go   
│   │   └── daily_information.go  # Daily information request handlers
│   │
│   │
│   └── repository/
│       ├── user_repository.go     # UserRepository interface and implementations 
│       ├── postgres_user_repository.go     # Concrete implementation of UserRepository
│       ├── daily_information_repository.go  # DailyInformationRepository interface and implementations
│       └── postgres_daily_information_repository.go     # Concrete implementation of DailyInformationRepository
│
│
│
├── pkg/
│   ├── types/
│   │   └── types.go       # Common data types
│   └── util/
│       └── util.go        # Utility functions
│   
├── .env
│
├── docker-compose.yml
│
└── Dockerfile
```

## Commands


Run:
```
docker-compose up
```


Stop and remove containers:
```
docker-compose down 
```

Remove the my-go-server image:
```
docker rmi my-go-server
```


Execute SQL queries against the PostgreSQL container:
```
docker exec -it <container_id_or_name> psql -U <username> <database_name>
```

