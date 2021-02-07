package go_bloom_test

import (
	"testing"

	"github.com/go-redis/redis"
	bloom "github.com/left-pocket/go-bloom"
)

func TestRedisBloomFilter(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 256,
	})

	m, k := bloom.EstimateParameters(100000, 0.0001)
	t.Logf("EstimateParameters, m=%d, k=%d", m, k)
	bitSet := bloom.NewRedisBitSet("test_key", m, client)
	b := bloom.New(m, k, bitSet)
	testBloomFilter(t, b)
}

func testBloomFilter(t *testing.T, b *bloom.BloomFilter) {
	data := []byte("some key")
	existsBefore, err := b.Exists(data)
	if err != nil {
		t.Fatal("Error checking for existence in bloom filter")
	}
	if existsBefore {
		t.Fatal("Bloom filter should not contain this data")
	}
	err = b.Add(data)
	if err != nil {
		t.Fatal("Error adding to bloom filter")
	}
	existsAfter, err := b.Exists(data)
	if err != nil {
		t.Fatal("Error checking for existence in bloom filter")
	}
	if !existsAfter {
		t.Fatal("Bloom filter should contain this data")
	}
}
