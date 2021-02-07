package go_bloom

import (
	"hash/fnv"
	"math"
)

type BitSetProvider interface {
	Set([]int64) error
	Test([]int64) (bool, error)
}

type BloomFilter struct {
	m      int64 // the size(bit) for the BloomFilter
	k      int64 // the hash function count
	bitSet BitSetProvider
}

func New(m int64, k int64, bitSet BitSetProvider) *BloomFilter {
	return &BloomFilter{m: m, k: k, bitSet: bitSet}
}

// EstimateParameters estimates requirements for m and k.
// Input: n: number of items, p: the err_rate
// Output: m: the total Size(bit), k: the hash function number.
// https://krisives.github.io/bloom-calculator/
func EstimateParameters(n uint, p float64) (int64, int64) {
	m := math.Ceil(float64(n) * math.Log(p) / math.Log(1.0/math.Pow(2.0, math.Ln2)))
	k := math.Ln2*m/float64(n) + 0.5

	return int64(m), int64(k)
}

func (f *BloomFilter) Add(data []byte) error {
	locations := f.getLocations(data)
	err := f.bitSet.Set(locations)
	if err != nil {
		return err
	}
	return nil
}

func (f *BloomFilter) Exists(data []byte) (bool, error) {
	locations := f.getLocations(data)
	isSet, err := f.bitSet.Test(locations)
	if err != nil {
		return false, err
	}
	if !isSet {
		return false, nil
	}

	return true, nil
}

func (f *BloomFilter) getLocations(data []byte) []int64 {
	locations := make([]int64, f.k)
	hasher := fnv.New64()
	hasher.Write(data)
	a := make([]byte, 1)
	for i := int64(0); i < f.k; i++ {
		a[0] = byte(i)
		hasher.Write(a)
		hashValue := hasher.Sum64()
		locations[i] = int64(hashValue % uint64(f.m))
	}
	return locations
}
