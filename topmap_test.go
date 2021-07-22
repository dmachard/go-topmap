package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var datamap = map[string]int{
	"apple":  1,
	"orange": 10,
	"red":    2,
	"blue":   50,
}

var datamap_sorted = []TopMapItem{
	{"blue", 50},
	{"orange", 10},
	{"red", 2},
	{"apple", 1},
}

var benchmarks = []struct{ MaxItems int }{
	{10},
	{100},
	{1000},
	{10000},
}

func TestTopMapRecord(t *testing.T) {
	tmap := NewTopMap(5)

	for name, hit := range datamap {
		tmap.Record(name, hit)
	}

	sorted := tmap.Get()
	if len(sorted) != len(datamap) {
		t.Error("invalid top map size", sorted)
	}
}

func TestTopMapGetSorted(t *testing.T) {
	tmap := NewTopMap(5)

	for name, hit := range datamap {
		tmap.Record(name, hit)
	}

	for index, item := range tmap.Get() {
		if datamap_sorted[index].Name != item.Name {
			t.Error("unexpected item name")
		}
		if datamap_sorted[index].Hit != item.Hit {
			t.Error("unexpected item value")
		}
	}
}

func TestTopMapRecordMax(t *testing.T) {
	tmap := NewTopMap(2)

	for name, hit := range datamap {
		tmap.Record(name, hit)
	}

	sorted := tmap.Get()
	if len(sorted) != 2 {
		t.Errorf("invalid top map size, should be 2 not %d", len(sorted))
	}
}

func BenchmarkTopMapRecord(b *testing.B) {
	rand.Seed(int64(b.N))
	for _, bm := range benchmarks {
		desc := fmt.Sprintf("TopMapSize: %d\n", bm.MaxItems)
		b.Run(desc, func(b *testing.B) {
			tmap := NewTopMap(bm.MaxItems)
			for n := 0; n < b.N; n++ {
				k := strconv.Itoa(n)
				tmap.Record(k, rand.Intn(bm.MaxItems))
			}
		})

	}
}

func BenchmarkTopMapGet(b *testing.B) {
	rand.Seed(int64(b.N))
	for _, bm := range benchmarks {
		desc := fmt.Sprintf("TopMapSize: %d\n", bm.MaxItems)
		b.Run(desc, func(b *testing.B) {
			tmap := NewTopMap(bm.MaxItems)
			for i := 0; i < bm.MaxItems; i++ {
				k := strconv.Itoa(i)
				tmap.Record(k, rand.Intn(bm.MaxItems))
			}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				r := tmap.Get()
				if len(r) != bm.MaxItems {
					break
				}
			}
		})

	}
}
