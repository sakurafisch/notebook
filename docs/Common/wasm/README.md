# wasm

[2022 wasm 现状](https://harshal.sheth.io/2022/01/31/webassembly.html)

> 有充分的[理由](https://kubesphere.io/blogs/will-cloud-native-webassembly-replace-docker_/)相信 Wasm 代表了容器化的未来。与 Docker 相比，它的冷启动时间快了 10-100 倍，占用空间更小，并使用了更好的基于能力的约束安全模型。使 Wasm 模块（而不是容器）成为计算和部署的标准单元将实现更好的可扩展性和安全性。


[wapm 包管理器](https://wapm.io/)

[krustlet 部署到 k8s](https://krustlet.dev/)

[spin 微服务](https://github.com/fermyon/spin)

[Knative](https://knative.dev/)

## 社区样例

[dapr-wasm](https://github.com/second-state/dapr-wasm) A template project to demonstrate how to run WebAssembly functions as sidecar microservices in dapr

## Cloud Native wasm

[Dapr](https://dapr.io/) wasm sidecar 模式就靠它了。社区实践可以参考[此文章](https://www.secondstate.io/articles/dapr-wasmedge-webassembly/)

[WasmEdge](https://github.com/WasmEdge/WasmEdge) WasmEdge is a lightweight, high-performance, and extensible WebAssembly runtime for cloud native, edge, and decentralized applications. It powers serverless apps, embedded functions, microservices, smart contracts, and IoT devices.

有关社区实践，建议阅读 [Manage WebAssembly Apps Using Container and Kubernetes Tools](https://www.secondstate.io/articles/manage-webassembly-apps-in-wasmedge-using-docker-tools/)

`WasmEdge` 有可能有利于实现微服务的 Side Car 模式。

[wasmCloud](https://github.com/wasmCloud/wasmCloud) wasmCloud allows for simple, secure, distributed application development using WebAssembly actors and capability providers.

## GraalVM Implementation

[Graal wasm](https://www.graalvm.org/22.1/reference-manual/wasm/) 官方实现，但截止写笔记时还不支持 `wasi`。(2022年6月22日)

[Truffle Wasm](https://www.research.manchester.ac.uk/portal/files/160212054/salim_TruffleWasm_vee_2020_authorversion.pdf) 第三方实现，首次支持 `wasi`。

## runtime

[wasmer](https://github.com/wasmerio/wasmer) 🚀 The leading WebAssembly Runtime supporting WASI and Emscripten.

我比较看好这个运行时。

[wasmtime](https://github.com/bytecodealliance/wasmtime) Standalone JIT-style runtime for WebAssembly, using Cranelift

我不看好这个运行时。既然都 JIT 了，那还不如只用 JavaScript。

