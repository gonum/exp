# dframe

`dframe` is a work-in-progress [Data Frame](https://en.wikipedia.org/wiki/Pandas_%28software%29) a-la [pandas](https://pandas.pydata.org/pandas-docs/stable/index.html).

`dframe` is leveraging [Apache Arrow](https://arrow.apache.org/) and its [Go backend](https://godoc.org/github.com/apache/arrow/go/arrow).

## Proposal

We propose to introduce a new `Frame` type inside the `dframe` package: a 2-dim data structure to handle:

- tabular data with heterogeneous columns (like a `SQL` table)
- arbitrary matrix data with row and column labels
- any other form of observational/statistical dataset.

For a good cross-pollination and integration with the Gonum and Go scientific ecosystem, it is expected for other "companion" packages tailored for a few focused operations to appear:

- integration with `gonum/plot`,
- integration with `gonum/stat`,
- integration with `gonum/mat` (_e.g.:_ creation of `dframe.Frame`s from `gonum/mat.Vector` or `gonum/mat.Matrix`, and vice versa)
- `hdf5` loading/saving of `dframe.Frame`s,
- integration with `encoding/csv` or `npyio`,
- integration with `database/sql`,
- etc...

### Previous work

The data frame concept comes from `R`'s `data.frame` and Python's `pandas.DataFrame`:

- https://www.rdocumentation.org/packages/base/versions/3.4.3/topics/data.frame
- https://pandas.pydata.org/pandas-docs/stable/generated/pandas.DataFrame.html

A few data frame-like implementations in Go have also been investigated:

- [kniren/gota](https://github.com/kniren/gota)
- [tobgu/qframe](https://github.com/tobgu/qframe)

Some inspiration from this previous body of work will be drawn, both in terms of API and performance hindsight.

### API

The main type should be:

```go
package dframe

type Frame struct {
	// contains filtered or unexported fields
}
```

It is expected to build `dframe.Frame` on top of `arrow/array.Interface` and/or `arrow/tensor.Interface` to re-use the SIMD optimized operations and zero-copy optimization that are implemented within these packages.
Using Arrow should also allow seamless interoperability with other data wrangling systems, possibly written in other languages than Go.

```go
// Open opens an already existing Frame using the provided driver technology,
// located at the provided source.
//
// Possible drivers: hdf5, npyio, csv, json, hdfs, spark, sql, ...
func Open(drv, src string) (*Frame, error) { ... }

// Create creates a new Frame, using the provided driver technology
func Create(drv, dst string, schema *arrow.Schema) (*Frame, error) { ... }

// New creates a new in-memory data frame with the provided memory schema.
func New(schema *arrow.Schema) (*Frame, error) { ... }
```
