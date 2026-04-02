# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is `github.com/pgavlin/text`, a Go package that provides generic versions of Go's `strings` package functions. It uses Go generics with the `String` constraint (`~string | ~[]byte`) so the same functions work on both string and byte slice types, avoiding the duplication between `strings` and `bytes` packages.

## Build and Test

```bash
go test ./...           # run all tests
go test -run TestFoo    # run a single test
go test -bench .        # run benchmarks
```

There are no linters, code generators, or build steps beyond `go test`.

## Architecture

The core type constraint is `String interface { ~string | ~[]byte }`, defined in `strings.go`. All public functions and generic types are parameterized on this constraint.

**Package structure:**
- Root package `text` — generic versions of `strings` functions (Contains, Split, Replace, Trim, etc.), plus `Builder[S]`, `Reader[S]`, `Replacer[S]`, and `Writer[S]` types
- `utf8/` — generic versions of `unicode/utf8` functions, parameterized on `~string | ~[]byte`
- `internal/bytealg/` — single helper `AsString` that uses `unsafe` to view any `String` as a `string` without allocation

**Key pattern:** Most functions convert their generic `String` input to `string` via `bytealg.AsString` (zero-copy unsafe cast), delegate to the stdlib `strings` package for the core logic, then convert results back to the generic type `S`. This avoids reimplementing algorithms while staying generic.

**Origin:** The code is forked/adapted from the Go standard library (`strings`, `unicode/utf8`). Copyright headers and structure reflect this.
