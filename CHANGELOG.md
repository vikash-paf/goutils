# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-03-23

### Added
- **28 Cleaned and Audited Packages**: Full review of the entire library for v1.0 compliance.
- **Verified Examples**: Added `example_test.go` for all 28 packages, providing verified GoDoc examples.
- **New syncx Primitives**: Added `Batcher` (time/size based flushing), `Debounce`, and `Throttle`.
- **Consistency**: Standardized GoDoc comments across the library.

### Changed
- **Documentation**: Removed AI-generated filler text from all READMEs and source files.
- **async**: Improved `RetryWithContext` to correctly handle pre-canceled contexts.
- **cryptox**: Simplified AES encryption error messages.
- **ds**: Optimized `Queue` memory management.

## [0.2.0] - 2024-05-15 (Approx)

### Added
- **New Packages**: `result`, `fsx`, `iterx`, `id`, `encodingx`.
- **Data Structures**: Added `Trie` and `DAG` to the `ds` package.
- **Concurrency**: Added `KeyMutex` and `MapBatched`/`ForEachBatched` to `parallel`.
- **Algorithms**: Added `algo.BinarySearch`.
- **Documentation**: Added detailed package-level READMEs for all existing modules.

## [0.1.0] - 2024-01-10 (Approx)

### Added
- **Initial Release**: Core generic utility library for Go 1.21+.
- **Slice Manipulation**: `Map`, `Filter`, `Reduce`, `Unique`, `Chunk`, `GroupBy`.
- **Basic Utilities**: `control`, `mathx`, `tuple`, `opt`.
- **Advanced Systems**: `PriorityQueue`, `RingBuffer`, `TokenBucket`, `CircuitBreaker`.
- **Worker Pool**: Initial `WorkerPool` implementation.
- **Time Extensions**: `timex` package for business day arithmetic.
- **Cache**: Initial `LRU` cache implementation.
