# Python 库

There are two most useful libraries in Python that can help you deal with tabular data:

- **[Pandas](https://pandas.pydata.org/)** allows you to manipulate so-called **Dataframes**, which are analogous to relational tables. You can have named columns, and perform different operations on row, columns and dataframes in general.
- **[Numpy](https://numpy.org/)** is a library for working with **tensors**, i.e. multi-dimensional **arrays**. Array has values of the same underlying type, and it is simpler than dataframe, but it offers more mathematical operations, and creates less overhead.

There are also a couple of other libraries you should know about:

- **[Matplotlib](https://matplotlib.org/)** is a library used for data visualization and plotting graphs
- **[SciPy](https://www.scipy.org/)** is a library with some additional scientific functions. We have already come across this library when talking about probability and statistics

```python
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from scipy import ... # you need to specify exact sub-packages that you need

```

