Go Learning Path (Hands-On)
===========================

Run any example
---------------
- From repo root: `go run guide/examples/01_vars_and_types.go`
- Swap the filename to run others (they are standalone `package main` files).

Recommended order
-----------------
1) Setup & basics  
   - `package`, imports, `main` function, comments  
   - Declarations: `var x int = 1`, `x := 1`, `const Pi = 3.14`, zero values  
   - Types: numeric, string/rune/byte, bool, arrays vs slices, maps  
   - Type conversion: `float64(i)`, not implicit casts  
   - Visibility: identifiers starting with capital letters are exported

2) Control flow  
   - `if` with optional short statement, `for` (3-clause, while-style, range)  
   - `switch` (value, tagless), `defer`, `break`/`continue`/`goto`

3) Functions  
   - Multiple return values, named results, variadic params, higher-order funcs  
   - Methods (value vs pointer receiver) and method sets  
   - First-class functions and closures

4) Data structures  
   - Structs, embedding, tags, literals, comparability  
   - Slices (len/cap, append, copy, subslices, make), arrays  
   - Maps (make, lookup, ok-idiom, delete), strings as bytes vs runes

5) Interfaces & generics  
   - Implicit satisfaction, small interfaces, empty interface `any`  
   - Type assertions and type switches  
   - Type parameters: constraints, `comparable`, custom constraints, instantiation

6) Errors & resources  
   - `error` type, wrapping: `fmt.Errorf("... %w", err)`  
   - Sentinel errors vs typed errors, `errors.Is` / `errors.As`  
   - `defer` for cleanup; `context.Context` for cancellation

7) Concurrency  
   - Goroutines, channels (unbuffered/buffered), `select`, timeouts  
   - Worker patterns, `sync.WaitGroup`, `sync.Mutex`, `atomic` basics  
   - Avoid data races: never share writable data without sync

Examples to run (in `guide/examples`)
-------------------------------------
- Basics/types: `01_vars_and_types.go`
- Control/functions: `02_control_and_funcs.go`
- Structs/interfaces: `03_structs_interfaces.go`
- Collections: `04_collections_maps_slices.go`
- Concurrency intro: `05_concurrency.go`
- Goroutine basics: `06_goroutines_basics.go`
- Channels patterns: `07_channels_patterns.go`

8) Testing & tooling  
   - `go test`, table tests, benchmarks, examples  
   - `go fmt`, `go vet`, `golangci-lint` (if used), `go mod tidy`  
   - Modules: `go mod init`, `go get`, versioning

Reading list (official & concise)
---------------------------------
- Tour of Go: https://go.dev/tour
- Effective Go: https://go.dev/doc/effective_go
- Go blog: https://go.dev/blog/
- Standard library index: https://pkg.go.dev/std

Next steps
----------
- Skim `guide/examples/*.go` in order; run them.  
- Modify values and print statements to cement understanding.  
- Try writing small utilities (file copier, HTTP client/server, CLI flag parser) using the same patterns.

