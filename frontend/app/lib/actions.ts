"use server";

export type NodeData = {
  release_date: number;
  title: string;
  description: string;
};

export type NodeShape = {
  data: NodeData;
  children: NodeShape[];
};

function exportToList(tree: NodeShape) {
  const root = tree.data;
  const nodes = [[root]];
  let level = 1;
  let hasNextLevel = false;
  let nextLevel: NodeShape[] | undefined[] = [];
  if (tree.children !== null) {
    hasNextLevel = true;
    nextLevel = tree.children;
  }
  while (hasNextLevel) {
    let currentNodes = nextLevel;
    let currentLevel = new Array(2 ** level++);
    nextLevel = new Array(2 ** level);
    hasNextLevel = false;
    let j = 0;
    for (let i = 0; i < currentNodes.length; i++) {
      let thisNode = currentNodes[i];
      if (thisNode === undefined || thisNode.data.release_date === -1) {
        currentLevel[i] = undefined;
        nextLevel[j++] = undefined;
        nextLevel[j++] = undefined;
      } else {
        currentLevel[i] = thisNode.data;
        nextLevel[j++] =
          thisNode.children === null ? undefined : thisNode.children[0];
        nextLevel[j++] =
          thisNode.children === null ? undefined : thisNode.children[1];
        if (thisNode.children) {
          hasNextLevel = true;
        }
      }
    }
    nodes.push(currentLevel);
  }
  console.log(nodes);
  return nodes;
}

export async function getTree() {
  const res = await fetch("http://localhost:5000/v1/trees", {
    cache: "no-store",
  });
  const tree = (await res.json()) as NodeShape;
  const nodes = exportToList(tree);
  return { tree, nodes };
}
