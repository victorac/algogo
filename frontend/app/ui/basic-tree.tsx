import { NodeData } from "../lib/actions";

export default function BasicTree({ nodes }: { nodes: NodeData[][] }) {
  return (
    <div className="flex flex-col items-center p-24 gap-10">
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
    </div>
  );
}
