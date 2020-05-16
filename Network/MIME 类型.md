# MIME 类型

[原文链接](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types)

MIME ：Multipurpose Internet Mail Extensions

**媒体类型**是一种标准，用来表示文档、文件或字节流的性质和格式。

> 它在 [IETF RFC 6838](https://tools.ietf.org/html/rfc6838) 中进行了定义和标准化。
>
> 所有的 `text` JavaScript 类型已经被 [RFC 4329](https://tools.ietf.org/html/rfc4329) 废弃。



- 可以在 [媒体类型](https://www.iana.org/assignments/media-types/media-types.xhtml)页面中找到最新的完整列表。

- 浏览器通常使用 MIME 类型（而不是文件扩展名）来确定如何处理 URL
- 浏览器可以通过请求头 [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) 来设置 [`X-Content-Type-Options`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/X-Content-Type-Options) 以阻止 MIME 嗅探。



## 通用结构

```
type/subtype
```

*type* 表示可以被分多个子类的独立类别。*subtype 表示细分后的每个类型。*

对大小写不敏感，但是传统写法都是小写。

## 默认类型

对于 text 文件类型若没有特定的 subtype，就使用 `text/plain`。类似的，二进制文件没有特定或已知的 subtype，即使用 `application/octet-stream`。

## 独立类型

```
text/plain
text/html
image/jpeg
image/png
audio/mpeg
audio/ogg
audio/*
video/mp4
application/*
application/json
application/javascript
application/ecmascript
application/octet-stream
```

*独立*类型表明了对文件的分类，可以是如下之一：

| 类型          | 描述                                                         | 典型示例                                                     |
| :------------ | :----------------------------------------------------------- | :----------------------------------------------------------- |
| `text`        | 表明文件是普通文本，理论上是人类可读                         | `text/plain`, `text/html`, `text/css, text/javascript`       |
| `image`       | 表明是某种图像。不包括视频，但是动态图（比如动态 gif）也使用 image 类型 | `image/gif`, `image/png`, `image/jpeg`, `image/bmp`, `image/webp`, `image/x-icon`, `image/vnd.microsoft.icon` |
| `audio`       | 表明是某种音频文件                                           | `audio/midi`, `audio/mpeg, audio/webm, audio/ogg, audio/wav` |
| `video`       | 表明是某种视频文件                                           | `video/webm`, `video/ogg`                                    |
| `application` | 表明是某种二进制数据                                         | `application/octet-stream`, `application/pkcs12`, `application/vnd.mspowerpoint`, `application/xhtml+xml`, `application/xml`, `application/pdf` |

## 常见类型

### `application/octet-stream`

这是应用程序文件的默认值。意思是 *未知的应用程序文件 ，*浏览器一般不会自动执行或询问执行。浏览器会像对待 设置了 HTTP 头 [`Content-Disposition`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Disposition) 值为 `attachment` 的文件一样来对待这类文件。

### `text/plain`

文本文件默认值。即使它*意味着未知的文本文件*，但浏览器认为是可以直接展示的。

### `text/css`

在网页中要被解析为 CSS 的任何 CSS 文件必须指定 MIME 为 `text/css`。通常，服务器不识别以.css 为后缀的文件的 MIME 类型，而是将其以 MIME 为 `text/plain` 或 `application/octet-stream` 来发送给浏览器：在这种情况下，大多数浏览器不识别其为 CSS 文件，直接忽略掉。特别要注意为 CSS 文件提供正确的 MIME 类型。

### `text/html`

所有的 HTML 内容都应该使用这种类型。XHTML 的其他 MIME 类型（如 `application/xml+html`）现在基本不再使用（HTML5 统一了这些格式）。

## 图片类型

只有一小部分图片类型是被广泛支持的，Web 安全的，可随时在 Web 页面中使用的：

| MIME 类型       | 图片类型                               |
| :-------------- | :------------------------------------- |
| `image/gif`     | GIF 图片 (无损耗压缩方面被 PNG 所替代) |
| `image/jpeg`    | JPEG 图片                              |
| `image/png`     | PNG 图片                               |
| `image/svg+xml` | SVG 图片 (矢量图)                      |

- 此处的类型划分有一定的争议，有人认为此处应该增加 WebP（`image/webp`），各浏览器对此的支持情况有所不同。
- 尽管 `image/vnd.microsoft.icon` [在 ANA 注册](https://www.iana.org/assignments/media-types/image/vnd.microsoft.icon) , 它仍然不被广泛支持，`image/x-icon` 被作为替代品使用。

## 音频与视频类型

| MIME 类型                                               | 音频或视频类型                                               |
| :------------------------------------------------------ | :----------------------------------------------------------- |
| `audio/wave` `audio/wav` `audio/x-wav` `audio/x-pn-wav` | 音频流媒体文件。一般支持 PCM 音频编码 (WAVE codec "1") ，其他解码器有限支持（如果有的话）。 |
| `audio/webm`                                            | WebM 音频文件格式。Vorbis 和 Opus 是其最常用的解码器。       |
| `video/webm`                                            | 采用 WebM 视频文件格式的音视频文件。VP8 和 VP9 是其最常用的视频解码器。Vorbis 和 Opus 是其最常用的音频解码器。 |
| `audio/ogg`                                             | 采用 OGG 多媒体文件格式的音频文件。 Vorbis 是这个多媒体文件格式最常用的音频解码器。 |
| `video/ogg`                                             | 采用 OGG 多媒体文件格式的音视频文件。常用的视频解码器是 Theora；音频解码器为 Vorbis 。 |
| `application/ogg`                                       | 采用 OGG 多媒体文件格式的音视频文件。常用的视频解码器是 Theora；音频解码器为 Vorbis 。 |
| `application/json`                                      | application/json (MIME_type) https://en.wikipedia.org/wiki/Media_type#Common_examples https://www.iana.org/assignments/media-types/application/json |

## multipart/form-data 简介

`multipart/form-data` 可用于 [HTML 表单](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/Forms)从浏览器发送信息给服务器。作为多部分文档格式，它由边界线（一个由`'--'` 开始的字符串）划分出的不同部分组成。每一部分有自己的实体，以及自己的 HTTP 请求头，[`Content-Disposition`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Disposition)和 [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) 用于文件上传领域，最常用的 ([`Content-Length`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Length) 因为边界线作为分隔符而被忽略）。

```
Content-Type: multipart/form-data; boundary=aBoundaryString
(other headers associated with the multipart document as a whole)

--aBoundaryString
Content-Disposition: form-data; name="myFile"; filename="img.jpg"
Content-Type: image/jpeg

(data)
--aBoundaryString
Content-Disposition: form-data; name="myField"

(data)
--aBoundaryString
(more subparts)
--aBoundaryString--
```

如下所示的表单:

```html
<form action="http://localhost:8000/" method="post" enctype="multipart/form-data">
  <input type="text" name="myTextField">
  <input type="checkbox" name="myCheckBox">Check</input>
  <input type="file" name="myFile">
  <button>Send the file</button>
</form>
```

会发送这样的请求:

```html
POST / HTTP/1.1
Host: localhost:8000
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Connection: keep-alive
Upgrade-Insecure-Requests: 1
Content-Type: multipart/form-data; boundary=---------------------------8721656041911415653955004498
Content-Length: 465

-----------------------------8721656041911415653955004498
Content-Disposition: form-data; name="myTextField"

Test
-----------------------------8721656041911415653955004498
Content-Disposition: form-data; name="myCheckBox"

on
-----------------------------8721656041911415653955004498
Content-Disposition: form-data; name="myFile"; filename="test.txt"
Content-Type: text/plain

Simple file.
-----------------------------8721656041911415653955004498--
```

## multipart/byteranges 简介

`multipart/byteranges` 用于把部分的响应报文发送回浏览器。当发送状态码 [`206`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/206)`Partial Content` 时，这个 MIME 类型用于指出这个文件由若干部分组成，每一个都有其请求范围。就像其他很多类型 [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type)使用分隔符来制定分界线。每一个不同的部分都有 [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type)这样的 HTTP 头来说明文件的实际类型，以及 [`Content-Range`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Range)来说明其范围。

```
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Type: multipart/byteranges; boundary=3d6b6a416f9b5
Content-Length: 385

--3d6b6a416f9b5
Content-Type: text/html
Content-Range: bytes 100-200/1270

eta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="vieport" content
--3d6b6a416f9b5
Content-Type: text/html
Content-Range: bytes 300-400/1270

-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: "Open Sans", "Helvetica
--3d6b6a416f9b5--
```