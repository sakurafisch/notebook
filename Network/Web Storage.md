# Web Storage

# 概念

Web Storage API 提供了存储机制，通过该机制，浏览器可以安全地存储键值对，比使用 cookie 更加直观。

Web Storage 包含如下两种机制：

- `sessionStorage` 为每一个给定的源（given origin）维持一个独立的存储区域，该存储区域在页面会话期间可用（即只要浏览器处于打开状态，包括页面重新加载和恢复）。
- `localStorage` 同样的功能，但是在浏览器关闭，然后重新打开后数据仍然存在。

这两种机制是通过 [`Window.sessionStorage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/sessionStorage) 和 [`Window.localStorage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/localStorage) 属性使用（更确切的说，在支持的浏览器中 `Window` 对象实现了 `WindowLocalStorage` 和 `WindowSessionStorage` 对象并挂在其 `localStorage` 和 `sessionStorage` 属性下）—— 调用其中任一对象会创建 [`Storage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage) 对象，通过 [`Storage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage) 对象，可以设置、获取和移除数据项。对于每个源（origin）`sessionStorage` 和 `localStorage` 使用不同的 Storage 对象 —— 独立运行和控制。

## Web Storage 接口

- [`Storage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage)

  允许你在一个特定域中设置，检索和删除数据和储存类型 (session or local.)

- [`Window`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window)

  Web Storage API 继承于 [`Window`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window) 对象，并提供两个新属性  — [`Window.sessionStorage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/sessionStorage) 和 [`Window.localStorage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/localStorage) — 它们分别地提供对当前域的会话和本地 [`Storage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage) 对象的访问。

- [`StorageEvent`](https://developer.mozilla.org/zh-CN/docs/Web/API/StorageEvent)

   当一个存储区更改时，存储事件从文档的 `Window` 对象上被发布。

## 参考文献

- [使用 Web Storage API](https://developer.mozilla.org/zh-CN/docs/Web/API/Web_Storage_API/Using_the_Web_Storage_API)

- [Web Storage API](https://developer.mozilla.org/zh-CN/docs/Web/API/Web_Storage_API)



