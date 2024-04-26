import { Hierarchy, getTree } from "./lib/actions";

export default async function Home() {
  const tree = await getTree();
  console.log(tree)
  const root = tree.data;
  const nodes = [[root]];
  let level = 1;
  let hasNextLevel = false;
  let nextLevel: Hierarchy[] | undefined[] = [];
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
      if (thisNode === undefined || thisNode.data.datetime === -1) {
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

  return (
    <main className="flex min-h-screen flex-col items-center p-24 justify-start gap-10">
      {nodes.map((level, index) => {
        return (
          <div key={index} className="flex items-center justify-center gap-48">
            {level.map((node, index) => {
              if (node) {
                return (
                  <div
                    key={node.datetime}
                    className="rounded-full bg-green-700 w-10 h-10 leading-10 text-center"
                  >
                    {node.datetime ?? ""}
                  </div>
                );
              }
              return <div key={`${index}-nil`} className="w-8 h-8"></div>;
            })}
          </div>
        );
      })}
    </main>
  );
}
