.
├── cmd/
│   └── main.go
│
├── internal/
│   ├── app/
│   │   ├── user/
│   │   │   ├── usecase/
│   │   │   │   ├── create.go
│   │   │   │   ├── update.go
│   │   │   │   ├── delete.go
│   │   │   │   └── get.go
│   │   │   ├── repository/
│   │   │   │   ├── mysql.go
│   │   │   │   └── redis.go
│   │   │   ├── entity/
│   │   │   │   └── user.go
│   │   │   └── delivery/
│   │   │       ├── http/
│   │   │       │   ├── handler.go
│   │   │       │   ├── router.go
│   │   │       │   └── middleware.go
│   │   │       └── grpc/
│   │   │           ├── handler.go
│   │   │           ├── server.go
│   │   │           └── middleware.go
│   │   ├── auth/
│   │   │   ├── usecase/
│   │   │   │   └── auth.go
│   │   │   └── delivery/
│   │   │       └── http/
│   │   │           ├── handler.go
│   │   │           ├── router.go
│   │   │           └── middleware.go
│   │   └── ...
│   │
│   ├── domain/
│   │   ├── entity/
│   │   │   └── user.go
│   │   ├── repository/
│   │   │   └── user.go
│   │   ├── usecase/
│   │   │   ├── user.go
│   │   │   └── auth.go
│   │   └── ...
│   │
│   ├── infra/
│   │   ├── persistence/
│   │   │   ├── mysql/
│   │   │   │   ├── mysql.go
│   │   │   │   └── migrate/
│   │   │   │       └── ...
│   │   │   └── redis/
│   │   │       └── redis.go
│   │   ├── delivery/
│   │   │   ├── http/
│   │   │   │   └── ...
│   │   │   └── grpc/
│   │   │       └── ...
│   │   ├── config/
│   │   │   └── config.go
│   │   └── ...
│   └── ...
│
├── pkg/
│   ├── middleware/
│   │   └── ...
│   ├── logger/
│   │   └── ...
│   └── ...
│
├── config/
│   └── ...
│
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── README.md
