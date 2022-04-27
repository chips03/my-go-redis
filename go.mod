module github.com/chips03/my-go-redis

go 1.18

require (
	github.com/chips03/my-go-redis/database v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
)

replace github.com/chips03/my-go-redis/database => ./database
