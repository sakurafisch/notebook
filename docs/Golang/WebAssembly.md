# WebAssembly

## Go 1.11 Release Notes

Go 1.11 adds an experimental port to [WebAssembly](https://webassembly.org/) (js/wasm).

Go programs currently compile to one WebAssembly module that includes the Go runtime for goroutine scheduling, garbage collection, maps, etc. As a result, the resulting size is at minimum around 2 MB, or 500 KB compressed. Go programs can call into JavaScript using the new experimental syscall/js package. Binary size and interop with other languages has not yet been a priority but may be addressed in future releases.

As a result of the addition of the new GOOS value "js" and GOARCH value "wasm", Go files named *_js.go or *_wasm.go will now be ignored by Go tools except when those GOOS/GOARCH values are being used. If you have existing filenames matching those patterns, you will need to rename them.

More information can be found on the [WebAssembly wiki page](https://go.dev/wiki/WebAssembly).

## Go 1.12 Release Notes

[syscall/js](https://go.dev/pkg/syscall/js/)

The `Callback` type and `NewCallback` function have been renamed; they are now called [`Func`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Func) and [`FuncOf`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#FuncOf), respectively. This is a breaking change, but WebAssembly support is still experimental and not yet subject to the [Go 1 compatibility promise](https://go.dev/doc/go1compat). Any code using the old names will need to be updated.

If a type implements the new [`Wrapper`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Wrapper) interface, [`ValueOf`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#ValueOf) will use it to return the JavaScript value for that type.

The meaning of the zero [`Value`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value) has changed. It now represents the JavaScript `undefined` value instead of the number zero. This is a breaking change, but WebAssembly support is still experimental and not yet subject to the [Go 1 compatibility promise](https://go.dev/doc/go1compat). Any code relying on the zero [`Value`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value) to mean the number zero will need to be updated.

The new [`Value.Truthy`](https://go.dev/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value.Truthy) method reports the [JavaScript "truthiness"](https://developer.mozilla.org/en-US/docs/Glossary/Truthy) of a given value.

## Go 1.13 Release Notes

For `GOARCH=wasm`, the new environment variable `GOWASM` takes a comma-separated list of experimental features that the binary gets compiled with. The valid values are documented [here](https://go.dev/cmd/go/#hdr-Environment_variables).

## Go 1.14 Release Notes

JavaScript values referenced from Go via `js.Value` objects can now be garbage collected.

`js.Value` values can no longer be compared using the `==` operator, and instead must be compared using their `Equal` method.

`js.Value` now has `IsUndefined`, `IsNull`, and `IsNaN` methods.

