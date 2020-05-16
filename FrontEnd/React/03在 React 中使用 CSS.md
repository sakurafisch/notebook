# React 中使用CSS 的几种方式

## 在组件中直接使用style

不需要组件从外部引入css文件，直接在组件中书写。

```jsx
import React, { Component } from "react";
const div1 = {
  width: "300px",
  margin: "30px auto",
  backgroundColor: "#44014C",  //驼峰法
  minHeight: "200px",
  boxSizing: "border-box"
};
class Test extends Component {
  constructor(props, context) {
    super(props);
  }
  render() {
    return (
     <div>
       <div style={div1}>123</div>
       <div style="background-color:red;">
     </div>
    );
  }
}
export default Test;
```

background-color，box-sizing 等属性要转换成驼峰命名法 backgroundColor，boxSizing。属性值写在双引号内。

## 在组件中引入[name].css文件

需要在当前组件开头使用import引入css文件。

```jsx
import React, { Component } from "react";
import TestChidren from "./TestChidren";
import "@/assets/css/index.css";
class Test extends Component {
  constructor(props, context) {
    super(props);
  }
  render() {
    return (
      <div>
        <div className="link-name">123</div>
        <TestChidren>测试子组件的样式</TestChidren>
      </div>
    );
  }
}
export default Test;
```

这种方式引入的css样式**会作用于当前组件及其所有后代组件。**

## 在组件中引入[name].scss文件

详见 [sass](https://www.sass.hk/)

```cmd
yarn add node-sass  # 安装node-sass
```

编写scss文件

```scss
//index.scss
.App{
  background-color: #282c34;
  .header{
    min-height: 100vh;
    color: white;
  }
}
```

这种方式引入的css样式**同样会作用于当前组件及其所有后代组件。**

## 在组件中引入[name].module.css文件

将 css 文件作为一个模块引入，这个模块中的所有 css ，只作用于当前组件。不会影响其后代组件。

```jsx
import React, { Component } from "react";
import TestChild from "./TestChild";
import moduleCss from "./test.module.css";
class Test extends Component {
  constructor(props, context) {
    super(props);
  }  
  render() {
    return (
     <div>
       <div className={moduleCss.linkName}>321321</div>
       <TestChild></TestChild>
     </div>
    );
  }
}
export default Test;
```

## 在组件中引入 [name].module.scss文件

```jsx
import React, { Component } from "react";
import TestChild from "./TestChild";
import moduleCss from "./test.module.scss";
class Test extends Component {
  constructor(props, context) {
    super(props);
  }  
  render() {
    return (
     <div>
       <div className={moduleCss.linkName}>321321</div>
       <TestChild></TestChild>
     </div>
    );
  }
}
export default Test;
```

## 使用styled-components

详见 [styled-components](https://www.styled-components.com/)

```cmd
yarn add styled-components  # 安装 styledcomponents
```

创建一个 js 文件，在此编写样式

```js
//style.js
import styled, { createGlobalStyle } from "styled-components";
export const SelfLink = styled.div`
  height: 50px;
  border: 1px solid red;
  color: yellow;
`;
export const SelfButton = styled.div`
  height: 150px;
  width: 150px;
  color: ${props => props.color};
  background-image: url(${props => props.src});
  background-size: 150px 150px;
`;
```

组件中使用 styled-components 样式

```jsx
import React, { Component } from "react";
import { SelfLink, SelfButton } from "./style";
class Test extends Component {
  constructor(props, context) {
    super(props);
  }  
  render() {
    return (
     <div>
       <SelfLink title="People's Republic of China">app.js</SelfLink>
       <SelfButton color="palevioletred" style={{ color: "pink" }} src={fist}>
          SelfButton
        </SelfButton>
     </div>
    );
  }
}
export default Test;
```

这种方式的样式也只对当前组件有效。

## 使用radium

详见 [radium](https://github.com/FormidableLabs/radium)

```cmd
yarn add radium  # 安装 radium
```

在react组件中引入 radium 并使用：

```jsx
import React, { Component } from "react";
import Radium from 'radium';
let styles = {
  base: {
    color: '#fff',
    ':hover': {
      background: '#0074d9'
    }
  },
  primary: {
    background: '#0074D9'
  },
  warning: {
    background: '#FF4136'
  }
};
class Test extends Component {
  constructor(props, context) {
    super(props);
  }  
  render() {
    return (
     <div>
      <button style={[ styles.base, styles.primary ]}>
        this is a primary button
      </button>
     </div>
    );
  }
}
export default Radium(Test); 
```

在export之前，必须用Radium包裹。