# redis-client

A Redis client for demo. It can be configured with Redis server hostname and port number via environment variables `REDIS_SERVER` and `REDIS_SERVER_PORT`. It has 100 most frequent English words built in. The client will increment the count of a random word from the 100-word list at the rate of 1 request per second. The 
random word is picked using exponential distribution, meaning that the most frequent word will have higher chance to be picked.

