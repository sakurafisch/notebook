# Flask

参考资料：

[安装](https://dormousehole.readthedocs.io/en/latest/installation.html#installation)

[快速上手](https://dormousehole.readthedocs.io/en/latest/quickstart.html)

[教程](https://dormousehole.readthedocs.io/en/latest/tutorial/index.html#tutorial)

[Flask 方案](https://dormousehole.readthedocs.io/en/latest/patterns/index.html#patterns)

[文档](https://dormousehole.readthedocs.io/en/latest/)

[API](https://dormousehole.readthedocs.io/en/latest/api.html#api)

------

```bash
pip install Flask # 安装 Flask，建议安装在 venv 中
```

```python
import flask
```

不要使用 `flask.py` 作为应用名称，这会与 Flask 本身发生冲突。

## 依赖 

当安装 Flask 时，以下配套软件会被自动安装。

- [Werkzeug](https://palletsprojects.com/p/werkzeug/) 用于实现 WSGI ，应用和服务之间的标准 Python 接口。
- [Jinja](https://palletsprojects.com/p/jinja/) 用于渲染页面的模板语言。
- [MarkupSafe](https://palletsprojects.com/p/markupsafe/) 与 Jinja 共用，在渲染页面时用于避免不可信的输入，防止注入攻击。
- [ItsDangerous](https://palletsprojects.com/p/itsdangerous/) 保证数据完整性的安全标志数据，用于保护 Flask 的 session cookie.
- [Click](https://palletsprojects.com/p/click/) 是一个命令行应用的框架。用于提供 `flask` 命令，并允许添加自定义 管理命令。

### 可选依赖 

以下配套软件不会被自动安装。如果安装了，那么 Flask 会检测到这些软件。

- [Blinker](https://pythonhosted.org/blinker/) 为 [信号](https://dormousehole.readthedocs.io/en/latest/signals.html#signals) 提供支持。
- [SimpleJSON](https://simplejson.readthedocs.io/) 是一个快速的 JSON 实现，兼容 Python’s `json` 模块。如果安装 了这个软件，那么会优先使用这个软件来进行 JSON 操作。
- [python-dotenv](https://github.com/theskumar/python-dotenv#readme) 当运行 `flask` 命令时为 [通过 dotenv 设置环境变量](https://dormousehole.readthedocs.io/en/latest/cli.html#dotenv) 提供支持。
- [Watchdog](https://pythonhosted.org/watchdog/) 为开发服务器提供快速高效的重载。

## 最小的 Flask 应用

```python
# hello.py
from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'
```

1. 首先我们导入了 [`Flask`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.Flask) 类。 该类的实例将会成为我们的 WSGI 应用。
2. 接着我们创建一个该类的实例。第一个参数是应用模块或者包的名称。如果你使用 一个单一模块（就像本例），那么应当使用 `__name__` ，因为名称会根据这个 模块是按应用方式使用还是作为一个模块导入而发生变化（可能是 ‘__main__’ ， 也可能是实际导入的名称）。这个参数是必需的，这样 Flask 才能知道在哪里可以 找到模板和静态文件等东西。更多内容详见 [`Flask`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.Flask) 文档。
3. 然后我们使用 [`route()`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.Flask.route) 装饰器来告诉 Flask 触发函数的 URL 。
4. 函数名称被用于生成相关联的 URL 。函数最后返回需要在用户浏览器中显示的信息。

## 运行应用

```bash
# 导出 FLASK_APP 环境变量
export FLASK_APP=hello.py # Linux系统 
set FLASK_APP=hello.py # cmd
$env:FLASK_APP = "hello.py" # Power Shell
```

`FLASK_APP` 环境变量中储存的是模块的名称，运行 `flask run` 命令就 会导入这个模块

```bash
# 运行应用（服务器仅本地可见）
flask run # 或者 python -m flask run
```

```cmd
# 运行应用（可公开访问）
flask run --host=0.0.0.0  # 这行代码告诉你的操作系统监听所有公开的 IP
```

## 调试模式

（只需要记录出错信息和追踪堆栈？参见 [应用错误处理](https://dormousehole.readthedocs.io/en/latest/errorhandling.html#application-errors) ）

服务器会在修改应用代码之后自动重启，并且当应用出错时还会提供一个有用的调试器。

```bash
# 打开所有开发功能（包括调试模式）
export FLASK_ENV=development # 在 Windows 下需要使用 set 来代替 export
flask run
```

这样可以实现以下功能：

1. 激活调试器。
2. 激活自动重载。
3. 打开 Flask 应用的调试模式。

还可以通过导出 `FLASK_DEBUG=1` 来单独控制调试模式的开关。

## 路由

使用 [`route()`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.Flask.route) 装饰器把函数绑定到 URL:

```python
@app.route('/')
def index():
    return 'Index Page'

@app.route('/hello')
def hello():
    return 'Hello, World'
```

可以动态变化 URL 的某些部分， 还可以为一个函数指定多个规则。

### 变量规则 

通过把 URL 的一部分标记为 `` 就可以在 URL 中添加变量。标记的 部分会作为关键字参数传递给函数。通过使用 `` ，可以 选择性的加上一个转换器，为变量指定规则。请看下面的例子:

```python
@app.route('/user/<username>')
def show_user_profile(username):
    # show the user profile for that user
    return 'User %s' % escape(username)

@app.route('/post/<int:post_id>')
def show_post(post_id):
    # show the post with the given id, the id is an integer
    return 'Post %d' % post_id

@app.route('/path/<path:subpath>')
def show_subpath(subpath):
    # show the subpath after /path/
    return 'Subpath %s' % escape(subpath)
```

转换器类型：

| `string` | （缺省值） 接受任何不包含斜杠的文本 |
| -------- | ----------------------------------- |
| `int`    | 接受正整数                          |
| `float`  | 接受正浮点数                        |
| `path`   | 类似 `string` ，但可以包含斜杠      |
| `uuid`   | 接受 UUID 字符串                    |

### 唯一的 URL / 重定向行为

以下两条规则的不同之处在于是否使用尾部的斜杠。:

```
@app.route('/projects/')
def projects():
    return 'The project page'

@app.route('/about')
def about():
    return 'The about page'
```

`projects` 的 URL 是中规中矩的，尾部有一个斜杠，看起来就如同一个文件夹。 访问一个没有斜杠结尾的 URL 时 Flask 会自动进行重定向，帮你在尾部加上一个斜杠。

`about` 的 URL 没有尾部斜杠，因此其行为表现与一个文件类似。如果访问这个 URL 时添加了尾部斜杠就会得到一个 404 错误。这样可以保持 URL 唯一，并帮助 搜索引擎避免重复索引同一页面。