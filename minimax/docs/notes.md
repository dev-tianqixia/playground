## minimax algorithm

#### basic minimax





#### alpha-beta pruning

##### how to make pruning effective?

the effectiveness of alpha-beta pruning depends on the order in which children are visited. 

ideally for max nodes, we want to visit the best child first so that rest of the children can all be pruned; similarly, for min nodes, we want to visit the worst child first to avoid exploring redundant scenarios.

2 obvious ways to rank each child:

1. the evaluator function.
2. previous search result of the game tree.



#### references

https://towardsdatascience.com/how-a-chess-playing-computer-thinks-about-its-next-move-8f028bd0e7b1

https://www.cs.cornell.edu/courses/cs312/2002sp/lectures/rec21.htm

https://cs.stanford.edu/people/eroberts/courses/soco/projects/2003-04/intelligent-search/minimax.html

https://cs.stanford.edu/people/eroberts/courses/soco/projects/2003-04/intelligent-search/alphabeta.html