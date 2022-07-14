# More

## Rust By Example

[Rust By Example](https://doc.rust-lang.org/rust-by-example/)

## 黑魔法

[nomicon](https://doc.rust-lang.org/nightly/nomicon/) 编写 unsafe Rust 所需要注意的问题。

## LeetCode

[zhangyuan 的仓库](https://github.com/zhangyuang/leetcode)

## 异步编程

[Asynchronous Programming in Rust](https://rust-lang.github.io/async-book/) 

[async-std](https://book.async.rs/) 

[tokio](https://tokio.rs/) 

## 并行编程

[rayon](https://github.com/rayon-rs/rayon) Rayon: A data parallelism library for Rust

## 网络编程

[hyper](https://hyper.rs/)
Fast and safe HTTP for the Rust language

[tonic](https://docs.rs/tonic/latest/tonic/)
A Rust implementation of gRPC, a high performance, open source, general RPC framework that puts mobile and HTTP/2 first.

[Tower](https://github.com/tower-rs/tower)
Tower is a library of modular and reusable components for building robust networking clients and servers. Tower aims to make it as easy as possible to build robust networking clients and servers. It is protocol agnostic, but is designed around a request / response pattern.

[axum](https://github.com/tokio-rs/axum)
Ergonomic and modular web framework built with Tokio, Tower, and Hyper.

[Actix-web](https://actix.rs/)
A powerful, pragmatic, and extremely fast web framework for Rust.
与 axum 对比：actix-web 自己加了一层runtime将线程作为Actor来管理多个线程，每个线程实际跑的都是tokio的单线程block_on，这样线程之间就没法任务窃取了，失去了tokio任务调度的优势，换取了无线程上下文切换的性能。这是actix-web和其他框架的主要区别。Axum则是完全利用tokio。

## 区块链

[substrate](https://substrate.io) 跨链

[solana](https://solana.com) 吞吐量大的公链

## UI

[yew](https://github.com/yewstack/yew) Rust / Wasm framework for building client web apps

[iced](https://github.com/iced-rs/iced) A cross-platform GUI library for Rust, inspired by Elm

# FFI

[rust-bindgen](https://github.com/rust-lang/rust-bindgen) Automatically generates Rust FFI bindings to C (and some C++) libraries.

[cbindgen](https://github.com/eqrion/cbindgen) A project for generating C bindings from Rust code.

[android-ndk-rs](Rust bindings to the Android NDK)

[deno_bindgen](https://github.com/littledivy/deno_bindgen) is an external tool to simplify glue code generation of Deno FFI libraries written in Rust.

[wasm_bindgen](https://github.com/rustwasm/wasm-bindgen) Import JavaScript things into Rust and export Rust things to JavaScript.

[node_bindgen](https://github.com/infinyon/node-bindgen) Easy way to write Node.js module using Rust
