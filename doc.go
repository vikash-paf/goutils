// Package goutils is a modern, zero-dependency, generic utility library for Go 1.21+.
//
// It provides highly-tested, idiomatic utilities for collections, synchronization,
// error handling, and data structures, designed to eliminate boilerplate in Go projects.
//
// The library is organized into specialized sub-packages:
//
//   - algo: Binary search and Top-K algorithms.
//   - async: Concurrent execution and retry logic.
//   - cache: Concurrent-safe LRU and LFU caches.
//   - control: Ternary operators and panic-handling utilities.
//   - cronx: Simple task scheduling.
//   - cryptox: Secure AES-GCM encryption.
//   - dict: Map manipulation helpers (Keys, Values, Merge).
//   - ds: Advanced data structures (PriorityQueue, RingBuffer, BloomFilter, etc.).
//   - encodingx: JSON and encoding utilities.
//   - errx: Multi-error aggregation.
//   - fsx: File system utilities.
//   - httpx: Safe HTTP client with timeouts.
//   - id: Unique ID generation (UUID, NanoID).
//   - iterx: Go 1.23+ iterator utilities.
//   - mathx: Generic math operations.
//   - opt: Null-safe Optional types.
//   - parallel: High-level parallel processing.
//   - poolx: Type-safe sync.Pool wrappers.
//   - ptr: Pointer-to-value conversion helpers.
//   - rate: Token bucket rate limiting.
//   - resilience: Fault-tolerance patterns (CircuitBreaker).
//   - result: Functional error handling (Result[T]).
//   - set: Generic Set data structure.
//   - slice: Comprehensive slice manipulation.
//   - str: String manipulation algorithms.
//   - syncx: Advanced concurrency primitives (Batcher, Debounce).
//   - timex: Time and business day arithmetic.
//   - tuple: Generic pairs and zipping.
package goutils
