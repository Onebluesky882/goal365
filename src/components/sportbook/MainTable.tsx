"use client";
import { SportbookRoot } from "@/types/sportbook";
import { Table, TableRow } from "../ui/table";
import { MatchWinner } from "./betTypeTable/MatchWinner";
import { useState } from "react";
import { TableBodySection } from "./TableBodySection";
import { TableHeaderCard } from "./TableHeaderCard";
import { TableHeaderSection } from "./TableHeaderSection";
import matchJsonv from "@/app/sportsbook/match_demo.json";
export type OddsClickHandler = (
  market: string,
  selection: string,
  odd: string,
) => void;

export const MatchRows = ({ data }: { data: SportbookRoot }) => {
  const {} = data;

  return (
    <>
      {/* Row 1: Home */}
      <TableRow className=""></TableRow>
    </>
  );
};

export function MainTable() {
  const [search, setSearch] = useState(false);

  return (
    <>
      <TableHeaderCard search={search} setSearch={setSearch} />

      <Table className="min-w-200 border border-gray-700">
        <TableHeaderSection />
        <TableBodySection
          Away={matchJsonv.away}
          Home={matchJsonv.home}
          asianHandicap={matchJsonv.asian_handicap}
          country={matchJsonv.country}
          firstHapdicap={matchJsonv.asian_handicap_fh}
          leagueName={matchJsonv.league}
          overUnderFistHaft={matchJsonv.over_under_fh}
          overUnderFullIime={matchJsonv.over_under_full_time}
          time={matchJsonv.timestamp}
        />
      </Table>
    </>
  );
}
