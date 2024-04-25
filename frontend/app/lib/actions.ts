"use server";

type Node = {
  datetime: number;
  title: string;
  description: string;
};

export type Hierarchy = {
  data: Node;
  children: Hierarchy[];
};

export async function getTree() {
  const res = await fetch("http://localhost:5000/v1/trees", { cache: 'no-store' });
  const tree = (await res.json()) as Hierarchy;


  return tree;
}
