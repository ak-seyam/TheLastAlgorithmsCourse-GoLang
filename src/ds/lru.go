package ds

type keyAware[K Hashable, T interface{}] struct {
	value T
	key   K
}

type LRU[K Hashable, T interface{}] struct {
	dll    *DLinkedList[keyAware[K, T]]
	m      map[K]*DLinkedListNode[keyAware[K, T]]
	cap    int
	Length int
}

func NewLRU[K Hashable, T interface{}](cap int) LRU[K, T] {
	return LRU[K, T]{
		cap: cap,
		dll: &DLinkedList[keyAware[K, T]]{
			Head: nil,
			Tail: nil,
		},
		m:      make(map[K]*DLinkedListNode[keyAware[K, T]]),
		Length: 0,
	}
}

func (cache LRU[K, T]) Get(key K) (*T, bool) {
	v, found := cache.m[key]
	if found {
		prev := v.Prev
		nxt := v.Next
		if prev != nil {
			prev.Next = nxt
		} else {
			return &(v.Value.value), true
		}
		if nxt != nil {
			nxt.Prev = prev
		}
		v.Next = cache.dll.Head
		v.Prev = nil
		cache.dll.Head = v
		return &(v.Value.value), true
	}
	return nil, false
}

func (cache *LRU[K, T]) Put(key K, value T) {
	// check if it exist
	_, found := cache.m[key]
	if !found {
		n := cache.dll.Prepend(keyAware[K, T]{
			value: value,
			key:   key,
		})
		cache.m[key] = n
		if cache.Length == cache.cap {
			oldTail := cache.dll.Tail
			if oldTail != nil {
				delete(cache.m, oldTail.Value.key)
				cache.dll.Pop()
			}
		} else {
			cache.Length++
		}
	} else {
		cache.dll.Prepend(keyAware[K, T]{
			key:   key,
			value: value,
		})
		cache.dll.Pop()
	}
}
