package configuration

import (
	"fmt"
	"os"
)

const AUTOPILOT_BASE_URL = "AUTOPILOT_BASE_URL"
const AUTOPILOT_API_KEY = "AUTOPILOT_API_KEY"
const ADDRESS_BOOK_REDIS_ADDR = "ADDRESS_BOOK_REDIS_ADDR"

type APIConfig struct {
	BaseUrl string
	ApiKey  string
}

type RedisConfig struct {
	Addr string
}

type Values struct {
	APIConfig   APIConfig
	RedisConfig RedisConfig
}

func readAPIConfig() (APIConfig, error) {
	url := os.Getenv(AUTOPILOT_BASE_URL)
	if len(url) == 0 {
		return APIConfig{}, fmt.Errorf("Unable to read env variable: %s", AUTOPILOT_BASE_URL)
	}
	apiKey := os.Getenv(AUTOPILOT_API_KEY)
	if len(apiKey) == 0 {
		return APIConfig{}, fmt.Errorf("Unable to read env variable: %s", AUTOPILOT_API_KEY)
	}
	return APIConfig{
		BaseUrl: url,
		ApiKey:  apiKey,
	}, nil
}

func readRedisConfig() (RedisConfig, error) {
	addr := os.Getenv(ADDRESS_BOOK_REDIS_ADDR)
	if len(addr) == 0 {
		return RedisConfig{}, fmt.Errorf("Unable to read env variable: %s", ADDRESS_BOOK_REDIS_ADDR)
	}

	return RedisConfig{
		Addr: addr,
	}, nil
}

func ReadConfig() (Values, error) {
	redisConfig, err := readRedisConfig()
	if err != nil {
		return Values{}, err
	}

	apiConfig, err := readAPIConfig()
	if err != nil {
		return Values{}, err
	}

	return Values{
		APIConfig:   apiConfig,
		RedisConfig: redisConfig,
	}, nil
}
