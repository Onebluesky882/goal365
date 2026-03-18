// betTypeTable/MatchWinner.tsx
import { TableCell } from "@/components/ui/table";
import React from "react";
import { OddsButton } from "../TableSportBook";

export type MatchWinnerProps = {
  markets: Record<string, number | string>; // Home/Draw/Away + odd
  onOddsClick: (data: {
    market: string;
    selection: string;
    odd: string;
  }) => void;
  getVal: (market: Record<string, number | string>, key: string) => string; // always return string
};

export const MatchWinner = ({
  markets,
  onOddsClick,
  getVal,
}: MatchWinnerProps) => {
  return (
    <>
      {["Home", "Draw", "Away"].map((selection) => (
        <TableCell
          key={selection}
          className="text-center border-r border-gray-700 p-0"
        >
          <OddsButton
            market="1X2"
            selection={selection}
            odd={getVal(markets, selection)}
            color="text-blue-400"
            onClick={() =>
              onOddsClick({
                market: "1X2",
                selection,
                odd: getVal(markets, selection),
              })
            }
          />
        </TableCell>
      ))}
    </>
  );
};
