"use client";

import { SportbookRoot } from "@/types/sportbook";
import { TableCell, TableHead, TableRow } from "../ui/table";
import { OddsButton } from "./TableSportBook";
import { MatchWinner } from "./betTypeTable/MatchWinner";

export type OddsClickHandler = (
  market: string,
  selection: string,
  odd: string,
) => void;

export const MatchRows = ({ data }: { data: SportbookRoot }) => {
  const {} = data;

  const mockMatch = {
    isLive: true,
    score: "2 - 1",
    time: "45'",
    home: "Manchester United",
  };

  const mockMarkets = {
    matchWinner: {
      Home: 1.85,
      Draw: 3.2,
      Away: 4.1,
    },
  };

  const getVal = (market: Record<string, number | string>, key: string) => {
    const val = market?.[key];
    return val !== undefined ? String(val) : "-";
  };
  const onOddsClick = (data: any) => {
    console.log("ODDS CLICK:", data);
  };
  return (
    <>
      {/* Row 1: Home */}
      <TableRow className="">
        <MatchWinner
          markets={mockMarkets.matchWinner}
          onOddsClick={onOddsClick}
          getVal={getVal}
        />
      </TableRow>
    </>
  );
};

export const TableHeaderRows = () => {
  return (
    <>
      <TableRow className="hover:bg-transparent h-7">
        <TableHead className="w-10 text-center text-[7px] font-bold border-r border-gray-700 p-0 uppercase text-gray-300">
          Time
        </TableHead>
        <TableHead className="w-30 text-left text-[7px] font-bold px-2 border-r border-gray-700 uppercase text-gray-300">
          Match
        </TableHead>
        <TableHead className="w-10 text-center text-[7px] font-bold border-r border-gray-700 p-0 text-gray-300">
          1X2
        </TableHead>
        <TableHead
          colSpan={3}
          className="text-center text-[7px] font-bold border-r border-gray-700 p-0 uppercase text-gray-300"
        >
          Handicap
        </TableHead>
        <TableHead
          colSpan={3}
          className="text-center text-[7px] font-bold p-0 uppercase text-gray-300"
        >
          Over/Under
        </TableHead>
      </TableRow>
      <TableRow className="bg-gray-800/40 h-5 text-[7px] uppercase opacity-50">
        <TableCell className="border-r border-gray-700 p-0 text-center text-gray-400">
          LIVE
        </TableCell>
        <TableCell className="border-r border-gray-700 p-0"></TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0 text-gray-400">
          1X2
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 text-gray-400">
          HDP
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 text-gray-400">
          H
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 text-gray-400">
          A
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 text-gray-400">
          O/U
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 text-gray-400">
          O
        </TableCell>
        <TableCell className="text-center text-gray-400">U</TableCell>
      </TableRow>
    </>
  );
};
