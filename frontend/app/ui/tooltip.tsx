export type TooltipData = {
  title: string;
  description: string;
};

import { TooltipWithBounds, defaultStyles } from "@visx/tooltip";

export default function Tooltip({
  tooltipOpen,
  tooltipTop,
  tooltipLeft,
  tooltipData,
}: {
  tooltipOpen: boolean;
  tooltipLeft: number;
  tooltipTop: number;
  tooltipData?: TooltipData;
}) {
  const tooltipStyles = {
    ...defaultStyles,
    backgroundColor: "rgba(53,71,125,0.8)",
    color: "white",
  };
  return tooltipOpen ? (
    <TooltipWithBounds
      key={Math.random()} // needed for bounds to update correctly
      left={tooltipLeft}
      top={tooltipTop}
      style={tooltipStyles}
      className="flex flex-col gap-2 max-w-[192px]"
    >
      <h3 className="text-sm font-semibold">{tooltipData?.title}</h3>
      <h4 className="text-xs line-clamp-4">{tooltipData?.description}</h4>
    </TooltipWithBounds>
  ) : null;
}
