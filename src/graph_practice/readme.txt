This folder contains the code to create a bi-directional graph, add node and edge, and traverse based on BFS algorithm.

Please note the code is based on the web post https://flaviocopes.com/golang-data-structure-graph/
I made necessary changes to simplify the code structure, including:
- Note struct now is simpler by using string-type as its value, not generic type.
- Remove the parameter of Traverse method. The code does not do any other operation on the traversed item, only print the traverse path.

Running Method:
1. Copy the whole folder to your editor (such as VSCode)
2. Set it as current working directory
3. Type "go run ."
