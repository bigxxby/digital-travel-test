package connection

import (
	"fmt"

	"github.com/bigxxby/digital-travel-test/internal/config"
	"github.com/go-redis/redis"
)

//get redis connection
// type Config struct {
// 	AppPort string

// 	DBUser     string
// 	DBPassword string
// 	DBHost     string
// 	DBPort     string
// 	DBName     string
// 	DBSSLMode  string

// 	//redis
// 	RedisHost string
// 	RedisPort string

// 	JwtSecret string
// }

func GetRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test the connection
	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return rdb, nil
}
