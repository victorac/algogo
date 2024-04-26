"use client";

import { NodeShape } from "../lib/actions";
import { useMemo } from "react";
import { Group } from "@visx/group";
import { Cluster, hierarchy } from "@visx/hierarchy";
import {
  HierarchyPointNode,
  HierarchyPointLink,
} from "@visx/hierarchy/lib/types";
import { LinkVertical } from "@visx/shape";
import { LinearGradient } from "@visx/gradient";

const citrus = "#ddf163";
const white = "#ffffff";
export const green = "#79d259";
const aqua = "#37ac8c";
const merlinsbeard = "#f7f7f3";
export const background = "#306c90";

function Node({ node }: { node: HierarchyPointNode<NodeShape> }) {
  const isRoot = node.depth === 0;
  const isParent = !!node.children;

  if (isRoot) return <RootNode node={node} />;

  return (
    <Group top={node.y} left={node.x}>
      {node.depth !== 0 && (
        <circle
          r={12}
          fill={background}
          stroke={isParent ? white : citrus}
          onClick={() => {
            alert(`clicked: ${JSON.stringify(node.data.data.datetime)}`);
          }}
        />
      )}
      <text
        dy=".33em"
        fontSize={9}
        fontFamily="Arial"
        textAnchor="middle"
        style={{ pointerEvents: "none" }}
        fill={isParent ? white : citrus}
      >
        {node.data.data.datetime}
      </text>
    </Group>
  );
}

function RootNode({ node }: { node: HierarchyPointNode<NodeShape> }) {
  const width = 40;
  const height = 20;
  const centerX = -width / 2;
  const centerY = -height / 2;

  return (
    <Group top={node.y} left={node.x}>
      <rect
        width={width}
        height={height}
        y={centerY}
        x={centerX}
        fill="url('#top')"
      />
      <text
        dy=".33em"
        fontSize={9}
        fontFamily="Arial"
        textAnchor="middle"
        style={{ pointerEvents: "none" }}
        fill={background}
      >
        {node.data.data.datetime}
      </text>
    </Group>
  );
}

export type DendrogramProps = {
  tree: NodeShape;
  width: number;
  height: number;
  margin?: { top: number; right: number; bottom: number; left: number };
};

const defaultMargin = { top: 40, left: 0, right: 0, bottom: 40 };
export default function Tree({ tree, width, height, margin }: DendrogramProps) {
  const data = useMemo(() => hierarchy<NodeShape>(tree), []);
  const xMax = width - (margin?.left ?? 0) - (margin?.right ?? 0);
  const yMax = height - (margin?.top ?? 0) - (margin?.bottom ?? 0);
  return width < 10 ? null : (
    <svg width={width} height={height}>
      <LinearGradient id="top" from={green} to={aqua} />
      <rect width={width} height={height} rx={14} fill={background} />
      <Cluster<NodeShape> root={data} size={[xMax, yMax]}>
        {(cluster) => (
          <Group top={margin?.top} left={margin?.left}>
            {cluster.links().map((link, i) => (
              <LinkVertical<
                HierarchyPointLink<NodeShape>,
                HierarchyPointNode<NodeShape>
              >
                key={`cluster-link-${i}`}
                data={link}
                stroke={merlinsbeard}
                strokeWidth="1"
                strokeOpacity={0.2}
                fill="none"
              />
            ))}
            {cluster.descendants().map((node, i) => (
              <Node key={`cluster-node-${i}`} node={node} />
            ))}
          </Group>
        )}
      </Cluster>
    </svg>
  );
}
