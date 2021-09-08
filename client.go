package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	serverName string
	port       int
)

func initialize() {
	var err error
	serverName = os.Getenv("REDIS_SERVER")
	if serverName == "" {
		serverName = "localhost"
	}

	portStr := os.Getenv("REDIS_SERVER_PORT")
	if portStr == "" {
		port = 6379
	} else {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			fmt.Errorf("Env var REDIS_SERVER_PORT of value %s is invalid\n", portStr)
		}
	}

	fmt.Printf("Redis server: %s\nPort number: %d\n", serverName, port)
}

func main() {
	initialize()

	words := []string{"as", "I", "his", "that", "he", "was", "for", "on", "are", "with", "they", "be", "at", "one", "have", "this", "from", "by", "hot", "word", "but", "what", "some", "is", "it", "you", "or", "had", "the", "of", "to", "and", "a", "in", "we", "can", "out", "other", "were", "which", "do", "their", "time", "if", "will", "how", "said", "an", "each", "tell", "does", "set", "three", "want", "air", "well", "also", "play", "small", "end", "put", "home", "read", "hand", "port", "large", "spell", "add", "even", "land", "here", "must", "big", "high", "such", "follow", "act", "why", "ask", "men", "change", "went", "light", "kind", "off", "need", "house", "picture", "try", "us", "again", "animal", "point", "mother", "world", "near", "build", "self", "earth", "father"}

	fmt.Println(len(words))

	const leng = 100

	ctx := context.Background()

	// connect redis server
	var rdb *redis.Client

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", serverName, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rand.Seed(time.Now().Unix())
	for {
		// get a random word using modified exponential distribution
		number := int(rand.ExpFloat64()*leng/2) % leng
		word := words[number]

		countStr, err := rdb.Get(ctx, word).Result()
		if err == redis.Nil {
			err := rdb.Set(ctx, word, 0, 0).Err()
			if err != nil {
				fmt.Printf("Error: cannot init [%s]\n", word)
			}
			fmt.Printf("[%s] 0\n", word)
		} else if err != nil {
			fmt.Printf("Error: cannot get [%s]\n", word)
		} else {
			count, _ := strconv.Atoi(countStr)
			err := rdb.Set(ctx, word, count+1, 0).Err()
			if err != nil {
				fmt.Printf("Error: cannot set [%s] count to %s\n", word)
			} else {
				fmt.Printf("[%s] %d\n", word, count+1)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
