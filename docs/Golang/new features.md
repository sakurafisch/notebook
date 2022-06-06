# New Features

## Go 1.5 Release Notes

### Map literals

Due to an oversight, the rule that allowed the element type to be elided from slice literals was not applied to map keys. This has been [corrected](https://go.dev/cl/2591) in Go 1.5. An example will make this clear. As of Go 1.5, this map literal,

```go
m := map[Point]string{
    Point{29.935523, 52.891566}:   "Persepolis",
    Point{-25.352594, 131.034361}: "Uluru",
    Point{37.422455, -122.084306}: "Googleplex",
}
```

may be written as follows, without the `Point` type listed explicitly:

```go
m := map[Point]string{
    {29.935523, 52.891566}:   "Persepolis",
    {-25.352594, 131.034361}: "Uluru",
    {37.422455, -122.084306}: "Googleplex",
}
```

## Go 1.4 Release Notes

### For-range loops

Up until Go 1.3, `for`-`range` loop had two forms

```
for i, v := range x {
	...
}
```

and

```
for i := range x {
	...
}
```

If one was not interested in the loop values, only the iteration itself, it was still necessary to mention a variable (probably the [blank identifier](https://go.dev/ref/spec#Blank_identifier), as in `for` `_` `=` `range` `x`), because the form

```
for range x {
	...
}
```

was not syntactically permitted.

This situation seemed awkward, so as of Go 1.4 the variable-free form is now legal. The pattern arises rarely but the code can be cleaner when it does.

*Updating*: The change is strictly backwards compatible to existing Go programs, but tools that analyze Go parse trees may need to be modified to accept this new form as the `Key` field of [`RangeStmt`](https://go.dev/pkg/go/ast/#RangeStmt) may now be `nil`.

### Method calls on **T

Given these declarations,

```go
type T int
func (T) M() {}
var x **T
```

both `gc` and `gccgo` accepted the method call

```go
x.M()
```

which is a double dereference of the pointer-to-pointer `x`. The Go specification allows a single dereference to be inserted automatically, but not two, so this call is erroneous according to the language definition. It has therefore been disallowed in Go 1.4, which is a breaking change, although very few programs will be affected.

*Updating*: Code that depends on the old, erroneous behavior will no longer compile but is easy to fix by adding an explicit dereference.

## Go 1.7 Release Notes

There is one tiny language change in this release. The section on [terminating statements](https://go.dev/ref/spec#Terminating_statements) clarifies that to determine whether a statement list ends in a terminating statement, the “final non-empty statement” is considered the end, matching the existing behavior of the gc and gccgo compiler toolchains. In earlier releases the definition referred only to the “final statement,” leaving the effect of trailing empty statements at the least unclear. The [`go/types`](https://go.dev/pkg/go/types/) package has been updated to match the gc and gccgo compiler toolchains in this respect. This change has no effect on the correctness of existing programs.

## Go 1.17 Release Notes

Go 1.17 includes three small enhancements to the language.

- [Conversions from slice to array pointer](https://go.dev/ref/spec#Conversions_from_slice_to_array_pointer): An expression `s` of type `[]T` may now be converted to array pointer type `*[N]T`. If `a` is the result of such a conversion, then corresponding indices that are in range refer to the same underlying elements: `&a[i] == &s[i]` for `0 <= i < N`. The conversion panics if `len(s)` is less than `N`.
- [`unsafe.Add`](https://go.dev/pkg/unsafe#Add): `unsafe.Add(ptr, len)` adds `len` to `ptr` and returns the updated pointer `unsafe.Pointer(uintptr(ptr) + uintptr(len))`.
- [`unsafe.Slice`](https://go.dev/pkg/unsafe#Slice): For expression `ptr` of type `*T`, `unsafe.Slice(ptr, len)` returns a slice of type `[]T` whose underlying array starts at `ptr` and whose length and capacity are `len`.

The package unsafe enhancements were added to simplify writing code that conforms to `unsafe.Pointer`'s [safety rules](https://go.dev/pkg/unsafe/#Pointer), but the rules remain unchanged. In particular, existing programs that correctly use `unsafe.Pointer` remain valid, and new programs must still follow the rules when using `unsafe.Add` or `unsafe.Slice`.

Note that the new conversion from slice to array pointer is the first case in which a type conversion can panic at run time. Analysis tools that assume type conversions can never panic should be updated to consider this possibility.

## Go 1.18 Release Notes

### Generics

[泛型的官方教程](https://go.dev/doc/tutorial/generics)

Go 1.18 includes an implementation of generic features as described by the [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md). This includes major - but fully backward-compatible - changes to the language.

These new language changes required a large amount of new code that has not had significant testing in production settings. That will only happen as more people write and use generic code. We believe that this feature is well implemented and high quality. However, unlike most aspects of Go, we can't back up that belief with real world experience. Therefore, while we encourage the use of generics where it makes sense, please use appropriate caution when deploying generic code in production.

While we believe that the new language features are well designed and clearly specified, it is possible that we have made mistakes. We want to stress that the [Go 1 compatibility guarantee](https://tip.golang.org/doc/go1compat) says "If it becomes necessary to address an inconsistency or incompleteness in the specification, resolving the issue could affect the meaning or legality of existing programs. We reserve the right to address such issues, including updating the implementations." It also says "If a compiler or library has a bug that violates the specification, a program that depends on the buggy behavior may break if the bug is fixed. We reserve the right to fix such bugs." In other words, it is possible that there will be code using generics that will work with the 1.18 release but break in later releases. We do not plan or expect to make any such change. However, breaking 1.18 programs in future releases may become necessary for reasons that we cannot today foresee. We will minimize any such breakage as much as possible, but we can't guarantee that the breakage will be zero.

The following is a list of the most visible changes. For a more comprehensive overview, see the [proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md). For details see the [language spec](https://tip.golang.org/ref/spec).

- The syntax for [function](https://tip.golang.org/ref/spec#Function_declarations) and [type declarations](https://tip.golang.org/ref/spec#Type_declarations) now accepts [type parameters](https://tip.golang.org/ref/spec#Type_parameter_declarations).
- Parameterized functions and types can be instantiated by following them with a list of type arguments in square brackets.
- The new token `~` has been added to the set of [operators and punctuation](https://tip.golang.org/ref/spec#Operators_and_punctuation).
- The syntax for [Interface types](https://tip.golang.org/ref/spec#Interface_types) now permits the embedding of arbitrary types (not just type names of interfaces) as well as union and `~T` type elements. Such interfaces may only be used as type constraints. An interface now defines a set of types as well as a set of methods.
- The new [predeclared identifier](https://tip.golang.org/ref/spec#Predeclared_identifiers) `any` is an alias for the empty interface. It may be used instead of `interface{}`.
- The new [predeclared identifier](https://tip.golang.org/ref/spec#Predeclared_identifiers) `comparable` is an interface that denotes the set of all types which can be compared using `==` or `!=`. It may only be used as (or embedded in) a type constraint.

There are three experimental packages using generics that may be useful. These packages are in x/exp repository; their API is not covered by the Go 1 guarantee and may change as we gain more experience with generics.

- [`golang.org/x/exp/constraints`](https://pkg.go.dev/golang.org/x/exp/constraints)

  Constraints that are useful for generic code, such as [`constraints.Ordered`](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered).

- [`golang.org/x/exp/slices`](https://pkg.go.dev/golang.org/x/exp/slices)

  A collection of generic functions that operate on slices of any element type.

- [`golang.org/x/exp/maps`](https://pkg.go.dev/golang.org/x/exp/maps)

  A collection of generic functions that operate on maps of any key or element type.



The current generics implementation has the following known limitations:

- The Go compiler cannot handle type declarations inside generic functions or methods. We hope to provide support for this feature in a future release.
- The Go compiler does not accept arguments of type parameter type with the predeclared functions `real`, `imag`, and `complex`. We hope to remove this restriction in a future release.
- The Go compiler only supports calling a method `m` on a value `x` of type parameter type `P` if `m` is explicitly declared by `P`'s constraint interface. Similarly, method values `x.m` and method expressions `P.m` also are only supported if `m` is explicitly declared by `P`, even though `m` might be in the method set of `P` by virtue of the fact that all types in `P` implement `m`. We hope to remove this restriction in a future release.
- The Go compiler does not support accessing a struct field `x.f` where `x` is of type parameter type even if all types in the type parameter's type set have a field `f`. We may remove this restriction in a future release.
- Embedding a type parameter, or a pointer to a type parameter, as an unnamed field in a struct type is not permitted. Similarly, embedding a type parameter in an interface type is not permitted. Whether these will ever be permitted is unclear at present.
- A union element with more than one term may not contain an interface type with a non-empty method set. Whether this will ever be permitted is unclear at present.



Generics also represent a large change for the Go ecosystem. While we have updated several core tools with generics support, there is much more to do. It will take time for remaining tools, documentation, and libraries to catch up with these language changes.

### Bug fixes

The Go 1.18 compiler now correctly reports `declared but not used` errors for variables that are set inside a function literal but are never used. Before Go 1.18, the compiler did not report an error in such cases. This fixes long-outstanding compiler issue [#8560](https://golang.org/issue/8560). As a result of this change, (possibly incorrect) programs may not compile anymore. The necessary fix is straightforward: fix the program if it was in fact incorrect, or use the offending variable, for instance by assigning it to the blank identifier `_`. Since `go vet` always pointed out this error, the number of affected programs is likely very small.

The Go 1.18 compiler now reports an overflow when passing a rune constant expression such as `'1' << 32` as an argument to the predeclared functions `print` and `println`, consistent with the behavior of user-defined functions. Before Go 1.18, the compiler did not report an error in such cases but silently accepted such constant arguments if they fit into an `int64`. As a result of this change, (possibly incorrect) programs may not compile anymore. The necessary fix is straightforward: fix the program if it was in fact incorrect, or explicitly convert the offending argument to the correct type. Since `go vet` always pointed out this error, the number of affected programs is likely very small.