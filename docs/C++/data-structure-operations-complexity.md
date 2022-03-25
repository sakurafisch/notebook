# Data Structure Operations Complexity

| Data Structure         | Access | Search | Insertion | Deletion | Comments                                             |
| ---------------------- | ------ | ------ | --------- | -------- | ---------------------------------------------------- |
| **Array**              | 1      | n      | n         | n        |                                                      |
| **Stack**              | n      | n      | 1         | 1        |                                                      |
| **Queue**              | n      | n      | 1         | 1        |                                                      |
| **Linked List**        | n      | n      | 1         | n        |                                                      |
| **Hash Table**         | -      | n      | n         | n        | In case of perfect hash function costs would be O(1) |
| **Binary Search Tree** | n      | n      | n         | n        | In case of balanced tree costs would be O(log(n))    |
| **B-Tree**             | log(n) | log(n) | log(n)    | log(n)   |                                                      |
| **Red-Black Tree**     | log(n) | log(n) | log(n)    | log(n)   |                                                      |
| **AVL Tree**           | log(n) | log(n) | log(n)    | log(n)   |                                                      |
| **Bloom Filter**       | -      | 1      | 1         | -        | False positives are possible while searching         |

