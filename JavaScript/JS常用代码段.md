# JS常用代码段

## 在网页中添加日期

```js
let today = new Date()
let formatDate = today.toDateString()
let selectElement = document.getElementById('date')
selectElement.innerHTML = formatDate
```

