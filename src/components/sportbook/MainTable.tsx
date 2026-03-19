"use client";
import { SportbookRoot } from "@/types/sportbook";
import { Table, TableRow } from "../ui/table";
import { MatchWinner } from "./betTypeTable/MatchWinner";
import { useState } from "react";
import { TableBodySection } from "./TableBodySection";
import { TableHeaderCard } from "./TableHeaderCard";
import { TableHeaderSection } from "./TableHeaderSection";

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

export function MainTable() {
  const [search, setSearch] = useState(false);
  console.log(search);
  return (
    <>
      <TableHeaderCard search={search} setSearch={setSearch} />

      <Table className="min-w-[800px] border border-gray-700">
        <TableHeaderSection />
        <TableBodySection />
      </Table>
    </>
  );
}
