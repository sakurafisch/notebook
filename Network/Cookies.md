# Cookies

**`Cookie`** 是一个请求首部，其中含有先前由服务器通过 [`Set-Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie) 首部投放并存储到客户端的 [HTTP cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)。

这个首部可能会被完全移除，例如在浏览器的隐私设置里面设置为禁用 cookie。

| Header type                                                  | [Request header](https://developer.mozilla.org/en-US/docs/Glossary/Request_header) |
| :----------------------------------------------------------- | ------------------------------------------------------------ |
| [Forbidden header name](https://developer.mozilla.org/en-US/docs/Glossary/Forbidden_header_name) | yes                                                          |

## 语法

```html
Cookie: <cookie-list>
Cookie: name=value
Cookie: name=value; name2=value2; name3=value3
```

<cookie-list>

一系列的名称 / 值对，形式为 `=。`名称 / 值对之间用分号和空格 ('; ') 隔开。

## 示例

```html
Cookie: PHPSESSID=298zf09hf012fh2; csrftoken=u32t4o3tb3gg43; _gat=1;
```

## 分类

Cookie 总是保存在客户端中，按在客户端中的存储位置，可分为内存 Cookie 和硬盘 Cookie。

内存 Cookie 由[浏览器](https://zh.wikipedia.org/wiki/浏览器)维护，保存在[内存](https://zh.wikipedia.org/wiki/内存)中，浏览器关闭后就消失了，其存在时间是短暂的。硬盘 Cookie 保存在[硬盘](https://zh.wikipedia.org/wiki/硬盘)里，有一个过期时间，除非用户手工清理或到了过期时间，硬盘 Cookie 不会被删除，其存在时间是长期的。所以，按存在时间，可分为非持久 Cookie 和持久 Cookie。

## 用途

因为 [HTTP 协议](https://zh.wikipedia.org/wiki/HTTP)是无状态的，即[服务器](https://zh.wikipedia.org/wiki/服务器)不知道用户上一次做了什么，这严重阻碍了[交互式 Web 应用程序](https://zh.wikipedia.org/wiki/交互式Web应用程序)的实现。在典型的网上购物场景中，用户浏览了几个页面，买了一盒饼干和两瓶饮料。最后结帐时，由于 HTTP 的无状态性，不通过额外的手段，服务器并不知道用户到底买了什么，所以 Cookie 就是用来绕开 HTTP 的无状态性的 “额外手段” 之一。服务器可以设置或读取 Cookies 中包含信息，借此维护用户跟服务器[会话](https://zh.wikipedia.org/wiki/会话_(计算机科学))中的状态。

在刚才的购物场景中，当用户选购了第一项商品，服务器在向用户发送网页的同时，还发送了一段 Cookie，记录着那项商品的信息。当用户访问另一个页面，浏览器会把 Cookie 发送给服务器，于是服务器知道他之前选购了什么。用户继续选购饮料，服务器就在原来那段 Cookie 里追加新的商品信息。结帐时，服务器读取发送来的 Cookie 就行了。

Cookie 另一个典型的应用是当登录一个网站时，网站往往会请求用户输入用户名和密码，并且用户可以勾选 “下次自动登录”。如果勾选了，那么下次访问同一网站时，用户会发现没输入用户名和密码就已经登录了。这正是因为前一次登录时，服务器发送了包含登录凭据（用户名加密码的某种[加密](https://zh.wikipedia.org/wiki/加密)形式）的 Cookie 到用户的硬盘上。第二次登录时，如果该 Cookie 尚未到期，浏览器会发送该 Cookie，服务器验证凭据，于是不必输入用户名和密码就让用户登录了。

## 识别功能

如果在一台计算机中安装多个浏览器，每个浏览器都会以独立的空间存放 Cookie。因为 Cookie 中不但可以确认用户信息，还能包含计算机和浏览器的信息，所以一个用户使用不同的浏览器登录或者用不同的计算机登录，都会得到不同的 Cookie 信息，另一方面，对于在同一台计算机上使用同一浏览器的多用户群，Cookie 不会区分他们的身份，除非他们使用不同的用户名登录。

## 偷窃 Cookies 和脚本攻击

- Cookie 盗贼：搜集用户 Cookie 并发给攻击者的黑客，攻击者将利用 Cookie 消息通过合法手段进入用户帐户。
- Cookie 投毒：一般认为，Cookie 在储存和传回服务器期间没有被修改过，而攻击者会在 Cookie 送回服务器之前对其进行修改，达到自己的目的。例如，在一个购物网站的 Cookie 中包含了顾客应付的款项，攻击者将该值改小，达到少付款的目的。

## 缺陷

1. Cookie 会被附加在每个 HTTP 请求中，所以无形中增加了流量。
2. 由于在 HTTP 请求中的 Cookie 是明文传递的，所以安全性成问题，除非用 [HTTPS](https://zh.wikipedia.org/wiki/HTTPS)。
3. Cookie 的大小限制在 4KB 左右，对于复杂的存储需求来说是不够用的。

## 替代品

- [Brownie 方案](http://sourceforge.net/projects/brownie)，是一项[开放源代码](https://zh.wikipedia.org/wiki/开放源代码)工程，由 [SourceForge](https://zh.wikipedia.org/wiki/SourceForge) 发起。Brownie 曾被用以共享在不同域中的接入，而 Cookies 则被构想成单一域中的接入。这项方案已经停止开发。
- [P3P](https://zh.wikipedia.org/wiki/P3P)，用以让用户获得更多控制个人隐私权利的 [协议](https://zh.wikipedia.org/wiki/网络传输协议)。在浏览网站时，它类似于 Cookie。
- 在与服务器传输数据时，通过在地址后面添加唯一[查询串](https://zh.wikipedia.org/w/index.php?title=查询串&action=edit&redlink=1)，让服务器识别是否合法用户，也可以避免使用 Cookie。