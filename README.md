# go-bloom

The Bloom Filter Implementation in Golang with Redis.

- [Bloom Filter Calculator][ref-bloom-calculator]
- [Bloom Filter Calculator2][ref-bloom-calculator2]

## Definition
n: the number of items in filter.
p: the probability of false positives. (0.01 means 1%)
k: the number of hash functions.
m: the number of bits in the filter.(the needed memory).

## Confirm the parameters
can use the link blow the estimate the n,p,k,m.
1. Can use the n,p as the input, then you can get hte k,m.
```
func EstimateParameters(n uint, p float64) (m int64, k int64)
```
2. Can use the n,p,k as the input, then you can get the m.
For the solution, you can directly call
```
func New(m int64, k int64, bitSet BitSetProvider) *BloomFilter
```
The difference is you can set the k yourself or use the value calculated by the system.

## Usage
```go
var client *redis.Client
client = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "",
    DB:       0,
    PoolSize: 256,
})
//use the estimate m and k.
m, k := bloom.EstimateParameters(100000, 0.001)
//new a Bloom Filter
bitSet := bloom.NewRedisBitSet("test_key", m, client)
b := bloom.New(m, k, bitSet)

//check exist
data := []byte("some key")
exists, err := b.Exists(data)
err = b.Add(data)
exists, err := b.Exists(data)
```

[ref-bloom-calculator]: https://krisives.github.io/bloom-calculator/
[ref-bloom-calculator2]: https://hur.st/bloomfilter/
