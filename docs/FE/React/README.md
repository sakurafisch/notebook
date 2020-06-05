# React

## 生命周期

```jsx
class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {date: new Date()};
  }

  componentDidMount() {

  }

  componentWillUnmount() {

  }

  componentDidUpdate() {
  
  }

  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
      </div>
    );
  }
}
```

上面代码中，`componentDidMount()`、`componentWillUnmount()`和`componentDidUpdate()`就是三个最常用的生命周期方法。其中，`componentDidMount()`会在组件挂载后自动调用，`componentWillUnmount()`会在组件卸载前自动调用，`componentDidUpdate()`会在 UI 每次更新后调用（即组件挂载成功以后，每次调用 render 方法，都会触发这个方法）。

还有三个生命周期方法，不是经常使用。

- `shouldComponentUpdate(nextProps, nextState)`：每当`this.props`或`this.state`有变化，在`render`方法执行之前，就会调用这个方法。该方法返回一个布尔值，表示是否应该继续执行`render`方法，即如果返回`false`，UI 就不会更新，默认返回`true`。组件挂载时，`render`方法的第一次执行，不会调用这个方法。

- `static getDerivedStateFromProps(props, state)`：该方法在`render`方法执行之前调用，包括组件的第一次记载。它应该返回一个新的 state 对象，通常用在组件状态依赖外部输入的参数的情况。

- `getSnapshotBeforeUpdate()`：该方法在每次 DOM 更新之前调用，用来收集 DOM 信息。它返回的值，将作为参数传入`componentDidUpdate()`方法。