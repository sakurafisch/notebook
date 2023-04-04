# Tokio

[官方Tutorial](https://tokio.rs/tokio/tutorial)

[Rustt中文翻译](https://github.com/rustlang-cn/Rustt/tree/main/Books/Tokio-Tutorial)

## channel

- mpsc, 多生产者，单消费者模式

- oneshot, 单生产者单消费，一次只能发送一条消息

- broadcast，多生产者，多消费者，其中每一条发送的消息都可以被所有接收者收到，因此是广播

- watch，单生产者，多消费者，只保存一条最新的消息，因此接收者只能看到最近的一条消息，例如，这种模式适用于配置文件变化的监听

## 文件io

Module tokio::[io](https://docs.rs/tokio/latest/tokio/io/index.html)


