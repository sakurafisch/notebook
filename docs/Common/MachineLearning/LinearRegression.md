# Linear Regression

## Linear regression equation

In algebraic terms, the model would be defined as y=mx+b, where

- y is miles per gallon—the value we want to predict.
- m is the slope of the line.
- x is pounds—our input value.
- b is the y-intercept.

In ML, we write the equation for a linear regression model as follows:

$$ y′=b+w_1x_1 $$

where:

- y′ is the predicted label—the output.
- b is the **bias** of the model. Bias is the same concept as the y-intercept in the algebraic equation for a line. In ML, bias is sometimes referred to as w0. Bias is a **parameter** of the model and is calculated during training.
- w1 is the **weight** of the feature. Weight is the same concept as the slope m in the algebraic equation for a line. Weight is a **parameter** of the model and is calculated during training.
- x1 is a **feature**—the input.

During training, the model calculates the weight and bias that produce the best model.



## Types of loss

In linear regression, there are four main types of loss, which are outlined in the following table.

| Loss type                     | Definition                                                   | Equation                             |
| :---------------------------- | :----------------------------------------------------------- | :----------------------------------- |
| **L<sub></sub>1 loss**                   | The sum of the absolute values of the difference between the predicted values and the actual values. | $$∑\|actual value−predicted value\|$$    |
| **Mean absolute error (MAE)** | The average of L1 losses across a set of examples.           | $$\frac{1}{N}∑\|actual value−predicted value\|$$ |
| **L<sub>2</sub> loss**                   | The sum of the squared difference between the predicted values and the actual values. | $$∑(actual value−predicted value)^2$$     |
| **Mean squared error (MSE)**  | The average of L2 losses across a set of examples.           | $$\frac{1}{N}∑(actual value−predicted value)^2$$  |

The functional difference between L1 loss and L2 loss (or between MAE and MSE) is squaring. When the difference between the prediction and label is large, squaring makes the loss even larger. When the difference is small (less than 1), squaring makes the loss even smaller.

When processing multiple examples at once, we recommend averaging the losses across all the examples, whether using MAE or MSE.

When choosing the best loss function, consider how you want the model to treat outliers. For instance, MSE moves the model more toward the outliers, while MAE doesn't. L2 loss incurs a much higher penalty for an outlier than L1 loss.

## Gradient descent

Gradient descent is a mathematical technique that iteratively finds the weights and bias that produce the model with the lowest loss. Gradient descent finds the best weight and bias by repeating the following process for a number of user-defined iterations.

The model begins training with randomized weights and biases near zero, and then repeats the following steps:

1. Calculate the loss with the current weight and bias.
2. Determine the direction to move the weights and bias that reduce loss.
3. Move the weight and bias values a small amount in the direction that reduces loss.
4. Return to step one and repeat the process until the model can't reduce the loss any further.

## Convergence and convex functions

The loss functions for linear models always produce a convex surface. As a result of this property, when a linear regression model converges, we know the model has found the weights and bias that produce the lowest loss.

If we graph the loss surface for a model with one feature, we can see its convex shape. 

A linear model converges when it's found the minimum loss. Therefore, additional iterations only cause gradient descent to move the weight and bias values in very small amounts around the minimum. If we graphed the weights and bias points during gradient descent, the points would look like a ball rolling down a hill, finally stopping at the point where there's no more downward slope.

It's important to note that the model almost never finds the exact minimum for each weight and bias, but instead finds a value very close to it. It's also important to note that the minimum for the weights and bias don't correspond to zero loss, only a value that produces the lowest loss for that parameter.

## Hyperparameters

Three common hyperparameters are:

- Learning rate
- Batch size
- Epochs

### Learning Rate

Learning rate is a floating point number you set that influences how quickly the model converges. If the learning rate is too low, the model can take a long time to converge. However, if the learning rate is too high, the model never converges, but instead bounces around the weights and bias that minimize the loss. The goal is to pick a learning rate that's not too high nor too low so that the model converges quickly.

The learning rate determines the magnitude of the changes to make to the weights and bias during each step of the gradient descent process. The model multiplies the gradient by the learning rate to determine the model's parameters (weight and bias values) for the next iteration. In the third step of gradient descent, the "small amount" to move in the direction of negative slope refers to the learning rate.

The difference between the old model parameters and the new model parameters is proportional to the slope of the loss function. For example, if the slope is large, the model takes a large step. If small, it takes a small step. For example, if the gradient's magnitude is 2.5 and the learning rate is 0.01, then the model will change the parameter by 0.025.

The ideal learning rate helps the model to converge within a reasonable number of iterations. 

### Batch Size

Batch size is a hyperparameter that refers to the number of examples the model processes before updating its weights and bias. You might think that the model should calculate the loss for every example in the dataset before updating the weights and bias. However, when a dataset contains hundreds of thousands or even millions of examples, using the full batch isn't practical.

Two common techniques to get the right gradient on average without needing to look at every example in the dataset before updating the weights and bias are: 

1. stochastic gradient descent
2. mini-batch stochastic gradient descent

When training a model, you might think that noise is an undesirable characteristic that should be eliminated. However, a certain amount of noise can be a good thing.

#### Stochastic gradient descent (SGD)

Stochastic gradient descent uses only a single example (a batch size of one) per iteration. Given enough iterations, SGD works but is very noisy. "Noise" refers to variations during training that cause the loss to increase rather than decrease during an iteration. The term "stochastic" indicates that the one example comprising each batch is chosen at random.

#### Mini-batch stochastic gradient descent (mini-batch SGD)

Mini-batch stochastic gradient descent is a compromise between full-batch and SGD. For number of data points, the batch size can be any number greater than 1 and less than N. The model chooses the examples included in each batch at random, averages their gradients, and then updates the weights and bias once per iteration.

Determining the number of examples for each batch depends on the dataset and the available compute resources. In general, small batch sizes behaves like SGD, and larger batch sizes behaves like full-batch gradient descent.

### Epochs

During training, an epoch means that the model has processed every example in the training set once. For example, given a training set with 1,000 examples and a mini-batch size of 100 examples, it will take the model 10 iterations to complete one epoch.

Training typically requires many epochs. That is, the system needs to process every example in the training set multiple times.

The number of epochs is a hyperparameter you set before the model begins training. In many cases, you'll need to experiment with how many epochs it takes for the model to converge. In general, more epochs produces a better model, but also takes more time to train.

The following table describes how batch size and epochs relate to the number of times a model updates its parameters.

| Batch type                             | When weights and bias updates occur                          |
| :------------------------------------- | :----------------------------------------------------------- |
| Full batch                             | After the model looks at all the examples in the dataset. For instance, if a dataset contains 1,000 examples and the model trains for 20 epochs, the model updates the weights and bias 20 times, once per epoch. |
| Stochastic gradient descent            | After the model looks at a single example from the dataset. For instance, if a dataset contains 1,000 examples and trains for 20 epochs, the model updates the weights and bias 20,000 times. |
| Mini-batch stochastic gradient descent | After the model looks at the examples in each batch. For instance, if a dataset contains 1,000 examples, and the batch size is 100, and the model trains for 20 epochs, the model updates the weights and bias 200 times. |
