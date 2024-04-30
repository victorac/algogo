import { getTree } from "./lib/actions";
import Menu from "./ui/menu";
import Tree from "./ui/tree";

export default async function Home() {
  const { tree } = await getTree();
  return (
    <main className="flex min-h-screen items-center justify-center p-10 gap-10">
      <Menu />
      <Tree
        tree={tree}
        height={700}
        width={900}
        margin={{ top: 50, bottom: 50, left: 20, right: 20 }}
      />
    </main>
  );
}
