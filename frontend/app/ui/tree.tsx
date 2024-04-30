"use client";

import { NodeShape } from "../lib/actions";
import { useEffect, useMemo, useRef, useState } from "react";
import { Group } from "@visx/group";
import { Cluster, hierarchy } from "@visx/hierarchy";
import {
  HierarchyPointNode,
  HierarchyPointLink,
} from "@visx/hierarchy/lib/types";
import { LinkVertical } from "@visx/shape";
import { LinearGradient } from "@visx/gradient";
import { useTooltip } from "@visx/tooltip";
import Tooltip, { TooltipData } from "./tooltip";

const citrus = "#ddf163";
const white = "#ffffff";
export const green = "#79d259";
const aqua = "#37ac8c";
const merlinsbeard = "#f7f7f3";
export const background = "#306c90";
let interval: number;

function Node({
  node,
  containerBounds,
  showTooltip,
  hideTooltip,
}: {
  node: HierarchyPointNode<NodeShape>;
  containerBounds: { top: number; left: number };
  showTooltip: (args: any) => void;
  hideTooltip: () => void;
}) {
  const isRoot = node.depth === 0;
  const isParent = !!node.children;

  // event handlers
  function handleMouseEnter(event: React.MouseEvent<SVGGElement>) {
    if (interval) clearInterval(interval);
    // coordinates should be relative to the container in which Tooltip is rendered
    const containerX =
      ("clientX" in event ? event.clientX : 0) - containerBounds.left;
    const containerY =
      ("clientY" in event ? event.clientY : 0) - containerBounds.top;
    showTooltip({
      tooltipLeft: containerX,
      tooltipTop: containerY,
      tooltipData: node.data.data,
    });
  }

  function handleMouseExit() {
    interval = setTimeout(hideTooltip, 200) as any;
  }

  if (isRoot)
    return (
      <RootNode
        node={node}
        handleMouseEnter={handleMouseEnter}
        handleMouseLeave={handleMouseExit}
      />
    );

  return (
    <Group
      top={node.y}
      left={node.x}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseExit}
    >
      {node.depth !== 0 && (
        <circle
          r={16}
          fill={background}
          stroke={isParent ? white : citrus}
          strokeWidth={2}
        />
      )}
      <text
        dy=".33em"
        fontSize={12}
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

function RootNode({
  node,
  handleMouseEnter,
  handleMouseLeave,
}: {
  node: HierarchyPointNode<NodeShape>;
  handleMouseEnter: (args: any) => void;
  handleMouseLeave: () => void;
}) {
  const width = 50;
  const height = 30;
  const centerX = -width / 2;
  const centerY = -height / 2;

  return (
    <Group
      top={node.y}
      left={node.x}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      <rect
        width={width}
        height={height}
        y={centerY}
        x={centerX}
        fill="url('#top')"
      />
      <text
        dy=".33em"
        fontSize={14}
        fontFamily="Arial"
        textAnchor="middle"
        fontWeight="bold"
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

export default function Tree({ tree, width, height, margin }: DendrogramProps) {
  const ref = useRef<SVGSVGElement>(null);
  const data = useMemo(() => hierarchy<NodeShape>(tree), []);
  const [containerBounds, setContainerBounds] = useState({ left: 0, top: 0 });

  useEffect(() => {
    if (ref.current) {
      const bounds = ref.current.getBoundingClientRect();
      setContainerBounds(bounds);
    }
  }, []);

  const {
    showTooltip,
    hideTooltip,
    tooltipOpen,
    tooltipData,
    tooltipLeft = 0,
    tooltipTop = 0,
  } = useTooltip<TooltipData>({
    // initial tooltip state
    tooltipOpen: false,
    tooltipLeft: width / 3,
    tooltipTop: height / 3,
  });

  const xMax = width - (margin?.left ?? 0) - (margin?.right ?? 0);
  const yMax = height - (margin?.top ?? 0) - (margin?.bottom ?? 0);
  return width < 10 ? null : (
    <div className="relative">
      <Tooltip
        tooltipData={tooltipData}
        tooltipLeft={tooltipLeft}
        tooltipTop={tooltipTop}
        tooltipOpen={tooltipOpen}
      />

      <svg ref={ref} width={width} height={height}>
        <LinearGradient id="top" from={green} to={aqua} />
        <rect width={width} height={height} rx={14} fill={background} />
        <Cluster<NodeShape> root={data} size={[xMax, yMax]}>
          {(cluster) => (
            <Group top={margin?.top} left={margin?.left}>
              {cluster.links().map((link, i) => {
                if (link.target.data.data.datetime === -1) {
                  return null;
                }
                return (
                  <LinkVertical<
                    HierarchyPointLink<NodeShape>,
                    HierarchyPointNode<NodeShape>
                  >
                    key={`cluster-link-${i}`}
                    data={link}
                    stroke={merlinsbeard}
                    strokeWidth="2"
                    strokeOpacity={0.2}
                    fill="none"
                  />
                );
              })}
              {cluster.descendants().map((node, i) => {
                if (node.data.data.datetime === -1) {
                  return null;
                }
                return (
                  <Node
                    key={`cluster-node-${i}`}
                    node={node}
                    showTooltip={showTooltip}
                    hideTooltip={hideTooltip}
                    containerBounds={containerBounds}
                  />
                );
              })}
            </Group>
          )}
        </Cluster>
      </svg>
    </div>
  );
}
