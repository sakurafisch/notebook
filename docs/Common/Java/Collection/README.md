# Collection

## 接口

参见 [接口](https://java.winnerwinter.com/collections/interfaces/index.html)

## 聚合操作

参见 [聚合操作](https://java.winnerwinter.com/collections/streams/index.html)

## 实现

参见 [实现](https://java.winnerwinter.com/collections/implementations/index.html)

- General-purpose implementations (通用实现) 是最常用的实现，专为日常使用而设计。它们在标题为通用实现的表中进行了总结。

- Special-purpose implementations (专用实现) 旨在用于特殊情况，并显示非标准性能特征、使用限制或行为。

- Concurrent implementations (并发实现) 旨在支持高并发性，通常以牺牲单线程性能为代价。这些实现是 java.util.concurrent 包的一部分。

- Wrapper implementations (包装实现) 与其他类型的实现(通常是通用实现)结合使用，以提供增加或限制的功能。

- Convenience implementations (便利实现) 是通常通过静态工厂方法提供的迷你实现，为特殊集合(例如，单例集)的通用实现提供方便，有效的替代方案。

- Abstract implementations (抽象实现) 是骨架实现，有助于构建自定义实现  稍后将在 Custom Collection Implementations 部分中进行介绍。一个高级主题，并不是特别困难，但相对较少的人需要这样做。

### Java 集合框架提供了几个核心接口的通用实现：

- 对于 Set 接口，HashSet 是最常用的实现。

- 对于 List 接口，ArrayList 是最常用的实现。

- 对于 Map 接口，HashMap 是最常用的实现。

- 对于 Queue 接口，LinkedList 是最常用的实现。

- 对于 Deque 接口，ArrayDeque 是最常用的实现。

每个通用实现都提供其接口中包含的所有可选操作。

Java 集合框架还为需要非标准性能、使用限制或其他不寻常行为的情况提供了几种特殊用途的实现。

java.util.concurrent 包中包含多个集合实现，这些实现是线程安全的，但不受单个排除锁的控制。

Collections 类(与 Collection 接口相对)提供了对集合进行操作或返回集合的静态方法，这些方法称为包装器实现。

最后，有几种便利实现，当你不需要它们的全部功能时，它可以比通用实现更有效。通过静态工厂方法提供便利实现。


### Set 实现

Set 实现分为通用和专用实现。

#### 通用 Set 实现

有三种通用的 [Set](https://docs.oracle.com/javase/8/docs/api/java/util/Set.html) 实现 [HashSet](https://docs.oracle.com/javase/8/docs/api/java/util/HashSet.html)，[TreeSet](https://docs.oracle.com/javase/8/docs/api/java/util/TreeSet.html)，[LinkedHashSet](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedHashSet.html)。

使用这三种中的哪一种通常是直截了当的。HashSet 比 TreeSet 快得多(大多数操作的常量时间与 log 时间)，但不提供排序保证。如果需要使用 SortedSet 接口中的操作，或者需要按值进行迭代，请使用 TreeSet;否则，请使用 HashSet。可以肯定的是，大多数时候你最终都会使用 HashSet。

LinkedHashSet 在某种意义上介于 HashSet 和 TreeSet 之间。作为一个具有链表的哈希表来实现，它提供 insertion-ordered (插入顺序) 迭代(最早插入到最近插入)并且运行几乎与 HashSet 一样快。LinkedHashSet 实现使其客户端免于由 HashSet 提供的未指定的、通常是混乱的排序影响，而不会导致与 TreeSet 相关的增加的成本。

#### 专用 Set 实现
有两个专用的 Set 实现 [EnumSet](https://docs.oracle.com/javase/8/docs/api/java/util/EnumSet.html) 和 [CopyOnWriteArraySet](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/CopyOnWriteArraySet.html)。

EnumSet 是枚举类型的高性能 Set 实现。枚举集的所有成员必须具有相同的枚举类型。在内部，它由位向量表示，通常是单个 long。枚举集支持枚举类型范围内的迭代。例如，给定星期几的枚举声明，你可以迭代工作日。EnumSet 类提供了一个简单的静态工厂。

```java
for (Day d : EnumSet.range(Day.MONDAY, Day.FRIDAY))
    System.out.println(d);
```

枚举集还为传统的位标志提供了丰富的，类型安全的替代品。

```java
EnumSet.of(Style.BOLD, Style.ITALIC)
```

CopyOnWriteArraySet 是一个由写时复制数组支持的 Set 实现。所有的可变操作，例如 add，set 和 remove，都是通过制作数组的新副本来实现的；不需要锁定。甚至迭代也可以安全地与元素插入和删除同时进行。与大多数 Set 实现不同，add，remove 和 contains 方法需要的时间与集合的大小成正比。此实现 仅 适用于很少修改但经常迭代的集。它非常适合维护必须防止重复的事件处理程序列表。

### List 实现
List 实现分为通用和专用实现。

#### 通用 List 实现

有两个通用的 List 实现 [ArrayList](https://docs.oracle.com/javase/8/docs/api/java/util/ArrayList.html) 和 [LinkedList](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedList.html)。

大多数情况下，你可能会使用 ArrayList，它提供了常量时间的位置访问，而且速度非常快。它不必为 List 中的每个元素分配节点对象，并且当它必须同时移动多个元素时，它可以利用 System.arraycopy 的优点。将 ArrayList 视为 Vector 而没有同步开销。

如果经常将元素添加到 List 的开头或迭代 List 以从其内部删除元素，则应考虑使用 LinkedList。这些操作需要 LinkedList 中的常量时间和 ArrayList 中的线性时间。但是你的性能要付出很大的代价。位置访问需要 LinkedList 中的线性时间和 ArrayList 中的常量时间。此外，LinkedList 的常数因子要差得多。如果你认为要使用 LinkedList，请在做出选择之前使用 LinkedList 和 ArrayList 来衡量应用程序的性能；ArrayList 通常更快。

ArrayList 有一个调整参数  initial capacity (初始容量)，指的是 ArrayList 在必须增长之前可以容纳的元素数。LinkedList 没有调整参数，有七个可选操作，其中一个是 clone。其他六个是 addFirst，getFirst，removeFirst，addLast，getLast，以及 removeLast。LinkedList 还实现了 Queue 接口。

#### 专用 List 实现

[CopyOnWriteArrayList](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/CopyOnWriteArrayList.html) 是由写时复制数组支持的 List 实现。此实现在本质上类似于 CopyOnWriteArraySet。即使在迭代期间也不需要同步，并且保证迭代器永远不会抛出 ConcurrentModificationException。此实现非常适合维护事件处理程序列表，其中更改很少发生，并且遍历频繁且可能耗时。

如果需要同步，Vector 将比用 Collections.synchronizedList 同步的 ArrayList 稍快。但是 Vector 有大量的遗留操作，所以要小心总是使用 List 接口操作 Vector，否则你以后可能将无法替换实现。

如果你的 List 的大小是固定的  也就是说，你永远不会使用 remove，add，或 containsAll 以外的任何批量操作  你有第三个选择，绝对值得考虑。有关详细信息，请参阅 [Convenience Implementations](https://java.winnerwinter.com/collections/implementations/convenience.html) 部分中的 Arrays.asList。

### Map 实现

Map 实现分为通用，专用和并发实现。

#### 通用 Map 实现

三个通用 [Map](https://docs.oracle.com/javase/8/docs/api/java/util/Map.html) 实现是 [HashMap](https://docs.oracle.com/javase/8/docs/api/java/util/HashMap.html)，[TreeMap](https://docs.oracle.com/javase/8/docs/api/java/util/TreeMap.html) 和 [LinkedHashMap](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedHashMap.html)。如果需要 SortedMap 操作或以键排序的 Collection-视图迭代，请使用 TreeMap;如果你想要最大速度而不关心迭代顺序，请使用 HashMap;如果你想要近似 HashMap 性能和插入顺序迭代，请使用 LinkedHashMap。在这方面，Map 的情况类似于 Set。同样，[Set Implementations](https://java.winnerwinter.com/collections/implementations/set.html) 部分中的所有其他内容也适用于 Map 实现。

LinkedHashMap 提供了 LinkedHashSet 不具备的两项功能。创建 LinkedHashMap 时，可以基于 key 的访问而不是插入顺序来对其进行排序。换句话说，仅查找与键相关联的值会将该键带到 map 的末尾。此外，LinkedHashMap 提供了 removeEldestEntry 方法，可以重写该方法，以便在将新映射添加到 map 时自动移除过时映射的策略。这使得实现自定义缓存变得非常容易。

例如，此覆盖将允许 map 多达 100 个条目，然后每次添加新条目时它将删除最旧条目，从而保持 100 个条目的稳定状态。

private static final int MAX_ENTRIES = 100;

protected boolean removeEldestEntry(Map.Entry eldest) {
    return size() > MAX_ENTRIES;
}

#### 专用 Map 实现

有三个专用 Map 实现 [EnumMap](https://docs.oracle.com/javase/8/docs/api/java/util/EnumMap.html)，[WeakHashMap](https://docs.oracle.com/javase/8/docs/api/java/util/WeakHashMap.html) 和 [IdentityHashMap](https://docs.oracle.com/javase/8/docs/api/java/util/IdentityHashMap.html)。EnumMap，内部实现为 array，是一个与枚举键一起使用的高性能 Map 实现。此实现将 Map 接口的丰富性和安全性与接近数组的速度相结合。如果要将枚举映射到值，则应始终使用 EnumMap 而不是数组。

WeakHashMap 是 Map 接口的一个实现，它只存储对其键的弱引用。仅存储弱引用允许在其键不再在 WeakHashMap 之外引用时对其进行垃圾回收。此类提供了利用弱引用功能的最简单方法。它对于实现“类似注册表”的数据结构很有用，其中当任何线程不再可以访问其键时，条目的实用程序就会消失。

IdentityHashMap 是基于哈希表的基于身份的 Map 实现。此类对于拓扑保留对象图转换非常有用，例如序列化或深度复制。要执行此类转换，你需要维护一个基于身份的“节点表”，以跟踪已经看到的对象。基于身份的 map 还用于维护动态调试器和类似系统中的对象到元信息映射。最后，基于身份的地图有助于阻止故意不正常的 equals 方法导致的“欺骗攻击”，因为 IdentityHashMap 永远不会调用键的 equals 的方法。这种实现的另一个好处是它很快。

#### 并发 Map 实现
[java.util.concurrent](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/package-summary.html) 包中包含 [ConcurrentMap](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/ConcurrentMap.html) 接口，该接口继承了 Map 原子的 putIfAbsent，remove，replace 方法，且 ConcurrentHashMap 实现那个接口。

ConcurrentHashMap 是一个由哈希表支持的高度并发、高性能的实现。执行获取时，此实现永远不会阻塞，并允许客户端选择更新的并发级别。它旨在作为 Hashtable 的替代品：除了实现 ConcurrentMap 之外，它还支持 Hashtable 特有的所有遗留方法。同样，如果你不需要遗留操作，请小心使用 ConcurrentMap 接口对其进行操作。

### Queue 实现

Queue 实现分为通用和并发实现。

#### 通用 Queue 实现
如前一节所述，LinkedList 实现 Queue 接口，为 add poll 等提供先进先出(FIFO)队列操作。

[PriorityQueue](https://docs.oracle.com/javase/8/docs/api/java/util/PriorityQueue.html) 类是基于 heap (堆) 数据结构的优先级队列。此队列根据构造时指定的顺序对元素进行排序，这可以是元素的自然顺序或由显式 Comparator 强加的排序。

队列获取操作  poll，remove，peek 和 element  访问队列头部的元素。队列的头 是关于指定排序的最小元素。如果多个元素被绑定为最小值，则头部是这些元素之一；关系被任意打破。

PriorityQueue 及其迭代器实现 Collection 和 Iterator 接口的所有可选方法。方法 iterator 中提供的迭代器不保证以任何特定顺序遍历 PriorityQueue 的元素。对于有序遍历，请考虑使用 Arrays.sort(pq.toArray())。

#### 并发 Queue 实现
java.util.concurrent 包中包含一组同步的 Queue 接口和类。[BlockingQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/BlockingQueue.html) 继承 Queue，其操作在获取元素时等待队列变为非空，并且在存储元素时等待队列的空间可用。此接口由以下类实现：

[LinkedBlockingQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/LinkedBlockingQueue.html) 由链接节点支持的可选有界 FIFO 阻塞队列
[ArrayBlockingQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/ArrayBlockingQueue.html) 由数组支持的有界 FIFO 阻塞队列
[PriorityBlockingQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/PriorityBlockingQueue.html) 由堆支持的无界阻塞优先级队列
[DelayQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/DelayQueue.html) 由堆支持的基于时间的调度队列
[SynchronousQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/SynchronousQueue.html) 一个简单的集合点机制，它使用 BlockingQueue 接口
在 JDK 7 中，[TransferQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/TransferQueue.html) 是一个专用的 BlockingQueue，其中向队列添加元素的代码可以选择等待(阻塞)代码在另一个线程中获取元素。TransferQueue 只有一个实现：

[LinkedTransferQueue](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/LinkedTransferQueue.html) 基于链接节点的无界 TransferQueue

### Deque 实现

Deque 接口，发音为 "deck"，表示双端队列。Deque 接口可以实现为各种类型的 Collections。Deque 接口实现分为通用和并发实现。

#### 通用 Deque 实现

通用实现包括 LinkedList 和 ArrayDeque 类。Deque 接口支持在两端插入，移除和获取元素。[ArrayDeque](https://docs.oracle.com/javase/8/docs/api/java/util/ArrayDeque.html) 类是 Deque 接口的可调整大小的数组实现，而 [LinkedList](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedList.html) 类是列表实现。

Deque 接口的基本插入，移除和获取操作 addFirst，addLast，removeFirst，removeLast，getFirst 和 getLast。方法 addFirst 在头部添加元素，而 addLast 在 Deque 实例的尾部添加元素。

LinkedList 实现比 ArrayDeque 实现更灵活。LinkedList 实现所有可选列表操作。null 元素在 LinkedList 实现中是允许的，但在 ArrayDeque 实现中不允许。

在效率方面，ArrayDeque 比 LinkedList 更有效，可以在两端添加和移除操作。LinkedList 实现中的最佳操作是在迭代期间移除当前元素。LinkedList 实现不是迭代的理想结构。

LinkedList 实现比 ArrayDeque 实现消耗更多内存。对于 ArrayDeque 实例遍历，请使用以下任何一种方法：

##### foreach

foreach 速度很快，可用于各种列表。

```java
ArrayDeque<String> aDeque = new ArrayDeque<String>();

for (String str : aDeque) {
    System.out.println(str);
}
```

##### Iterator

Iterator 可用于各类数据的各种列表的正向遍历。

```java
ArrayDeque<String> aDeque = new ArrayDeque<String>();

for (Iterator<String> iter = aDeque.iterator(); iter.hasNext();  ) {
    System.out.println(iter.next());
}
```

本教程中使用 ArrayDeque 类来实现 Deque 接口。本教程中使用的示例的完整代码可在 [ArrayDequeSample](https://java.winnerwinter.com/collections/interfaces/examples/ArrayDequeSample.java) 中找到。LinkedList 和 ArrayDeque 类都不支持多个线程的并发访问。

#### 并发 Deque 实现

[LinkedBlockingDeque](https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/LinkedBlockingDeque.html) 类是 Deque 接口的并发实现。如果 deque 为空，那么诸如 takeFirst 和 takeLast 之类的方法会等到元素变为可用，然后获取并移除相同的元素。

#### 包装器实现

参见 [包装器实现](https://java.winnerwinter.com/collections/implementations/wrapper.html)

#### 便利实现

参加 [便利实现](https://java.winnerwinter.com/collections/implementations/convenience.html)
