# Used Graphviz

### Input

1. Run script with input.txt

2. Generate the svg, this should output a graph with 2 blobs, just look for the nodes you can disconnect

```bash
dot -Tsvg -Kneato  days/day25/output.dot > days/day25/output.svg
```

3. Update the input and remove the connections from the list.

4. Re-run the script with input.txt

5. After that, you're left with 2 distinct blobs. Then use ccomps to output the subgraphs and count their nodes

```bash
ccomps -e -s -v days/day25/output.dot
```
