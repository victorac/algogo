import { getTree } from "./lib/actions";
import Tree from "./ui/tree";

export default async function Home() {
  const { tree } = await getTree();
  console.log(tree);
  return (
    <main className="flex min-h-screen flex-col items-center p-10 justify-between gap-10">
      <Tree
        tree={tree}
        height={700}
        width={900}
        margin={{ top: 20, bottom: 20, left: 20, right: 20 }}
      />
    </main>
  );
}
