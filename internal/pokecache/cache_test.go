package pokecache

import(
	"testing"
	"time"
	"bytes"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key:	"https://example.com",
			val:	[]byte("testdata"),
		},
		{
			key:	"https://example.com/path",
			val:	[]byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache := NewCache(interval)
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("key not found")
		}
		if !bytes.Equal(val, c.val) {
			t.Errorf("value: %s doesn't match expected: %s", string(val), string(c.val))
		}
	}
}

func TestReapLoop(t *testing.T) {
	
	cache := NewCache(5 * time.Millisecond)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep((5 * time.Millisecond) * 2)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
