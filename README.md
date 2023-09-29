## How to reproduce the issue

Please run `go test`.

## What Did I Do?

I attempted to read all cells using the `*excelize.File.GetRows()` method.

I wrote test code to reproduce the issue, which can be found [here](https://github.com/nobishino/reproduce-excelize)

## What Did I Expect to See?

I expected the result slice to contain all numbers.

## What Did I See Instead?

Instead, I observed that it returned an empty string `""` for cells with the format `#,#####.`

## my environment

### ` go version`

<pre>
go version
go version go1.21.0 darwin/amd64
</pre>

### `excelize` version

- The latest `v2.8.0` release reproduces the issue.
- `v2.7.1` does not reproduce the issue.
- https://github.com/qax-os/excelize/commit/7c221cf29531fcd38871d3295f4b511029cb4282 seems first commit which reproduces this issue.

### OS

<details><summary><code>go env</code> Output</summary><br><pre>
go env
GO111MODULE=''
GOARCH='amd64'
GOHOSTARCH='amd64'
GOHOSTOS='darwin'
</pre></details>
