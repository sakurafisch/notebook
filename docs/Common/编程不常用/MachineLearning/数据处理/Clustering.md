# Clustering

## Goal

The goal of clustering is to 

1. group data points that are close (or similar) to each other.

2. identify such groupings(or clusters) in an unsupervised manner

## Hierarchical methods

- Single-link method

- Complete-link method

- Sum-of-squares method

### Single-link method

We shall illustrate the method by example with an applomerative algorithm in which, at each stage of the algorithm, the closest two groups are fused to form a new group where the distance between two groups, A and B, is the distance between their closest members.

### Complete-link method

In the complete-link or furthest-neighbour method the distance between two groups A and B is the distance between the two furthest points, one taken from each group.

### Sum-of-squares method

At each stage of the algorithm, the two groups that produce the smallest increase in the total within-group sum of squares are amalgamated. The dissimilarity between two groups is defined to be the increase in the total sum of squares that would result if they were amalgamated. The updating formula for the dissimiliarity matrix is……

## K-means

k-means 的思想不同于层次聚类 hierarchical methods，它属于分散性聚类。