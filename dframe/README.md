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

### dframe

The main type should be:

```go
package dframe

type Frame struct {
	// contains filtered or unexported fields
}

// Err returns the first error encountered during operations on a Frame.
func (df *Frame) Err() error { ... }

// NumRows returns the number of rows of this Frame.
func (df *Frame) NumRows() int { ... }

// NumCols returns the number of columns of this Frame.
func (df *Frame) NumCols() int { ... }

// Column returns the i-th column of this Frame.
func (df *Frame) Column(i int) *array.Column { ... }

// ColumnNames returns the list of column names of this Frame.
func (df *Frame) ColumnNames() []string { ... }
```

It is expected to build `dframe.Frame` on top of `arrow/array.Interface` and/or `arrow/tensor.Interface` to re-use the SIMD optimized operations and zero-copy optimization that are implemented within these packages.
Using Arrow should also allow seamless interoperability with other data wrangling systems, possibly written in other languages than Go.

`tobgu/qframe` presents a `QFrame` type that is essentially immutable.
Operations on a `QFrame`, such as copying columns, dropping columns, sorting them or applying some kind of operation on columns, return a new `QFrame`, leaving the original untouched.

Arrow uses a ref-counting mechanism for all the types that involve memory allocation (mainly to address workloads involving memory allocated on a GPGPU, by a SQL database or a mmap-file.)
This ref-counting mechanism is presented to the user as a pair of methods `Retain/Release` that resp. increment and decrement that reference count.
At first, it would seem this mechanism would prevent to expose an API with "chained methods", as the intermediate `Frame` would be "leaked":

```go
o := df.Slice(0, 10).Select("col1", "col2").Apply("col1 + col2")
```

If we want an immutable `Frame`, the code above should be rewritten as:

```go
sli := df.Slice(0, 10)
defer sli.Release()

sel := sli.Select("col1", "col2")
defer sel.Release()

o := sel.Apply("col1 + col2")
defer o.Release()
```
It is not clear (to me!) yet whether an immutable `Frame` makes much sense in Go and with this ref-counting mechanism coming from Arrow.

But, immutable or not, one could recoup the nice "chained methods" API by introducing a `dframe.Tx` transaction:

```go
// Exec runs the provided function inside an atomic read/write transaction,
// applied on this Frame.
func (df *Frame) Exec(f func(tx *Tx) error) error { ... }

func example(df *dframe.Frame) {
	err := df.Exec(func(tx *dframe.Tx) error {
		tx.Slice(0, 10).Select("col1", "col2").Apply("col1 + col2")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

Or, without a "chained methods" API:

```go
func example(df *dframe.Frame) {
	err := df.Exec(func(tx *dframe.Tx) error {
		tx.Slice(0, 10)
		tx.Select("col1", "col2")
		tx.Apply("col1 + col2")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

```go
// Open opens an already existing Frame using the provided driver technology,
// located at the provided source.
//
// Possible drivers: hdf5, npyio, csv, json, hdfs, spark, sql, ...
func Open(drv, src string) (*Frame, error) { ... }

// Create creates a new Frame, using the provided driver technology
func Create(drv, dst string, schema *arrow.Schema, opts ...Option) (*Frame, error) { ... }

// New creates a new in-memory data frame with the provided memory schema.
func New(schema *arrow.Schema, opts ...Option) (*Frame, error) { ... }

// FromMem creates a new data frame from the provided in-memory data.
func FromMem(dict Dict, opts ...Option) (*Frame, error) { ... }

// FromArrays creates a new data frame from the provided schema and arrays.
func FromArrays(schema *arrow.Schema, arrs []array.Interface, opts ...Option) (*Frame, error) { ... }

// FromCols creates a new data frame from the provided schema and columns.
func FromCols(cols []array.Column, opts ...Option) (*Frame, error) { ... }

// FromTable creates a new data frame from the provided arrow table.
func FromTable(tbl array.Table, opts ...Option) (*Frame, error) { ... }

// FromFrame returns a new data frame created by applying the provided
// transaction on the provided frame.
func FromFrame(df *Frame, f func(tx *Tx) error) (*Frame, error) { ... }

// Exec runs the provided function inside an atomic read/write transaction,
// applied on this Frame.
func (df *Frame) Exec(f func(tx *Tx) error) error { ... }

// RExec runs the provided function inside an atomic read-only transaction,
// applied on this Frame.
func (df *Frame) RExec(f func(tx *Tx) error) error { ... }
```

### Operations

One should be able to carry the following operations on a `dframe.Frame`:

- retrieve the list of columns that a `Frame` is made of,
- create new columns that are the result of an operation on a set of already existing columns of that `Frame`,
- drop columns from a `Frame`
- append new data to a `Frame`, (either a new column or a new row)
- select a subset of columns from a `Frame`
- create different versions of a `Frame`: _e.g._ create `sub` from `Frame` `df` where `sub` is a subset of `df`.

