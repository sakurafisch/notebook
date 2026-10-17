# HTTP Cookies

HTTP Cookie（也叫 Web Cookie 或浏览器 Cookie）是服务器发送到用户浏览器并保存在本地的一小块数据，它会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上。通常，它用于告知服务端两个请求是否来自同一浏览器，如保持用户的登录状态。Cookie 使基于[无状态](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview#HTTP_is_stateless_but_not_sessionless)的 HTTP 协议记录稳定的状态信息成为了可能。

Cookie 主要用于以下三个方面：

- 会话状态管理（如用户登录状态、购物车、游戏分数或其它需要记录的信息）
- 个性化设置（如用户自定义设置、主题等）
- 浏览器行为跟踪（如跟踪分析用户行为等）

Cookie 曾一度用于客户端数据的存储，因当时并没有其它合适的存储办法而作为唯一的存储手段，但现在随着现代浏览器开始支持各种各样的存储方式，Cookie 渐渐被淘汰。由于服务器指定 Cookie 后，浏览器的每次请求都会携带 Cookie 数据，会带来额外的性能开销（尤其是在移动环境下）。新的浏览器 API 已经允许开发者直接将数据存储到本地，如使用 [Web storage API](https://developer.mozilla.org/zh-CN/docs/Web/API/Web_Storage_API) （本地存储和会话存储）或 [IndexedDB](https://developer.mozilla.org/zh-CN/docs/Web/API/IndexedDB_API) 。

## 创建 Cookie

当服务器收到 HTTP 请求时，服务器可以在响应头里面添加一个 [`Set-Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie)选项。浏览器收到响应后通常会保存下 Cookie，之后对该服务器每一次请求中都通过 [`Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Cookie)请求头部将 Cookie 信息发送给服务器。另外，Cookie 的过期时间、域、路径、有效期、适用站点都可以根据需要来指定。

### `Set-Cookie响应头部` 和 `Cookie请求头部`

服务器使用 [`Set-Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie)响应头部向用户代理（一般是浏览器）发送 Cookie 信息。一个简单的 Cookie 可能像这样：

```
Set-Cookie: <cookie名>=<cookie值>
```

服务器通过该头部告知客户端保存 Cookie 信息。

Tips：如何在以下几种服务端程序中设i置 `Set-Cookie` 响应头信息 :

- [PHP](https://secure.php.net/manual/en/function.setcookie.php)
- [Node.JS](https://nodejs.org/dist/latest-v8.x/docs/api/http.html#http_response_setheader_name_value)
- [Python](https://docs.python.org/3/library/http.cookies.html)
- [Ruby on Rails](http://api.rubyonrails.org/classes/ActionDispatch/Cookies.html)

```html
HTTP/1.0 200 OK
Content-type: text/html
Set-Cookie: yummy_cookie=choco
Set-Cookie: tasty_cookie=strawberry

[页面内容]
```

现在，对该服务器发起的每一次新请求，浏览器都会将之前保存的 Cookie 信息通过 [`Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Cookie) 请求头部再发送给服务器。

```html
GET /sample_page.html HTTP/1.1
Host: www.example.org
Cookie: yummy_cookie=choco; tasty_cookie=strawberry
```

### 会话期 Cookie

会话期 Cookie 是最简单的 Cookie：

- 浏览器关闭之后它会被自动删除，也就是说它仅在会话期内有效。

- 不需要指定过期时间（`Expires`）或者有效期（Max-Age）。

有些浏览器提供了会话恢复功能，这种情况下即使关闭了浏览器，会话期 Cookie 也会被保留下来，就好像浏览器从来没有关闭一样。

### 持久性 Cookie

可以指定一个特定的过期时间（`Expires`）或有效期（`Max-Age`）。

```html
Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT;
```

当 Cookie 的过期时间被设定时，设定的日期和时间只与客户端相关，而不是服务端。

### `Secure` 标记

- 标记为 `Secure` 的 Cookie 只应通过被 HTTPS 协议加密过的请求发送给服务端。
- 即便设置了 `Secure` 标记，敏感信息也不应该通过 Cookie 传输，因为 Cookie 有其固有的不安全性，`Secure `标记也无法提供确实的安全保障。
- 从 Chrome 52 和 Firefox 52 开始，不安全的站点（`http:`）无法使用 Cookie 的 `Secure` 标记。

###  `HttpOnly` 标记

- 为避免跨域脚本 ([XSS](https://developer.mozilla.org/en-US/docs/Glossary/XSS)) 攻击，通过 JavaScript 的 [`Document.cookie`](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/cookie) API 无法访问带有 `HttpOnly` 标记的 Cookie，它们只应该发送给服务端。

- 如果包含服务端 Session 信息的 Cookie 不想被客户端 JavaScript 脚本调用，那么就应该为其设置 `HttpOnly` 标记。

### 作用域

`Domain` 和 `Path` 标识定义了 Cookie 的*作用域：*即 Cookie 应该发送给哪些 URL。

`Domain` 标识指定了哪些主机可以接受 Cookie。如果不指定，默认为[当前文档的主机](https://developer.mozilla.org/en-US/docs/Web/API/Document/location)（**不包含子域名**）。如果指定了 `Domain`，则一般包含子域名。

`Path` 标识指定了主机下的哪些路径可以接受 Cookie（该 URL 路径必须存在于请求 URL 中）。以字符 `%x2F` ("/") 作为路径分隔符，**子路径也会被匹配。**

### `SameSite` Cookies 

`SameSite` Cookie 允许服务器要求某个 cookie 在跨站请求时不会被发送，从而可以阻止跨站请求伪造攻击（[CSRF](https://developer.mozilla.org/en-US/docs/Glossary/CSRF)）。

```js
Set-Cookie: key=value; SameSite=Strict
```

`SameSite` 参数的可选值：

- `None` ：浏览器会在同站请求、跨站请求下继续发送 cookies，不区分大小写。

- `Strict` ：浏览器将只发送相同站点请求的 cookie (即当前网页 URL 与请求目标 URL 完全一致)。如果请求来自与当前 location 的 URL 不同的 URL，则不包括标记为 Strict 属性的 cookie。

- `Lax` ：在新版本浏览器中，为默认选项，Same-site cookies 将会为一些跨站子请求保留，如图片加载或者 frames 的调用，但只有当用户从外部站点导航到 URL 时才会发送。如 link 链接。

### JavaScript 通过 Document.cookie 访问 Cookie

通过 [`Document.cookie`](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/cookie)属性可创建新的 Cookie，也可通过该属性访问非 `HttpOnly` 标记的 Cookie。

```js
document.cookie = "yummy_cookie=choco"; 
document.cookie = "tasty_cookie=strawberry"; 
console.log(document.cookie); 
// logs "yummy_cookie=choco; tasty_cookie=strawberry"
```

JavaScript 可以通过跨站脚本攻击（XSS）的方式来窃取 Cookie。

