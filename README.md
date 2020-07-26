<p align="center">
  <b>Randr</b>
</p>

Randr is a `golang` library to render HTML templates for `server-side rendering`, highly inspired by [`react`](https://github.com/facebook/react), [`preact/htm`](https://github.com/developit/htm) and [`lit-element`](https://github.com/polymer/lit-html). It is still a proof-of-concept, but should perform way better compared to `html/template` because it compiles expressions to static code in golang, which means it also has *0 runtime overhead*(similarly to [`quicktemplate`](https://github.com/valyala/quicktemplate)).

## Built with `randr`:
  - [lucat1/helmet](https://github.com/lucat1/helmet) - A `react-helmet` like component
  - [lucat1/css](https://github.com/lucat1/css) - A CSS styling library

## Installation

To install the compiler and the library it is suggested to use the standard `go get` command:

```sh
  $ go get -u github.com/lucat1/randr/rcc # get the compiler
  $ go get -u github.com/lucat1/randr # get the library (inside a project root)
```

## Usage

> NOTE: we assume that `$GOPATH/bin` is available inside your `$PATH`.
To compile a file or a folder of files you should use the `rcc`(`randr code compiler`) tool as follows:

```sh
  $ rcc <input> <output>
```

`<input>` can be either a file or a folder

`<output>` is the output location, and supports currenly only the `[name]` placeholder

Here's an example:

```sh
  $ rcc src "dist/[name].randr.go"
```
