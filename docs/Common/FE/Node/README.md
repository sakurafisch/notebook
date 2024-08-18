# Node.js

各个异步 API 优先级：

`process.nextTick` > `promise.then` > `setImmediate` > `setTimeout` = `setInterval`

## Common errors

- Uncaught TypeError: undefined is not a promise

If you get the `Uncaught TypeError: undefined is not a promise` error in the console, make sure you use `new Promise()` instead of just `Promise()`

- UnhandledPromiseRejectionWarning

This means that a promise you called rejected, but there was no `catch` used to handle the error. Add a `catch` after the offending `then` to handle this properly.

