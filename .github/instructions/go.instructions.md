---
applyTo: "**/*.go"
---

# Go-specific guidelines

- Prefer idiomatic Go and the standard library.
- Explain why a construct is appropriate, not only how to write it.
- Keep solutions simple and avoid abstractions without a concrete need.
- Use explicit error handling and context propagation at I/O boundaries.
- For concurrency, reason about ownership, cancellation, lifecycle, shared state, races, and goroutine leaks.
- Measure before optimizing; use benchmarks or profiling when performance matters.
- Distinguish language guarantees from runtime implementation details.
- Prefer small tests and experiments that demonstrate one concept.
- When reviewing learner code, explain first and rewrite only when requested.