// Package dict provides generic utility functions for map/dictionary manipulation.
package dict

// Keys returns a slice containing all the keys of the given map.
// The order of keys is not guaranteed.
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all the values of the given map.
// The order of values is not guaranteed.
func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Merge combines multiple maps into a single map. If duplicate keys exist,
// the value from the map appearing later in the arguments will overwrite the earlier one.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Invert creates a new map where the keys and values are swapped.
// If there are duplicate values in the original map, the resulting key will map
// to the last key encountered during iteration.
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	if m == nil {
		return nil
	}
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Omit creates a new map omitting the specified keys from the original map.
func Omit[K comparable, V any](m map[K]V, keysToOmit ...K) map[K]V {
	if m == nil {
		return nil
	}

	omitSet := make(map[K]struct{}, len(keysToOmit))
	for _, k := range keysToOmit {
		omitSet[k] = struct{}{}
	}

	result := make(map[K]V)
	for k, v := range m {
		if _, shouldOmit := omitSet[k]; !shouldOmit {
			result[k] = v
		}
	}
	return result
}
