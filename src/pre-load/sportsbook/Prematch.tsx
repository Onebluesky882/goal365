"use client";
import MenuBar from "@/components/sportbook/MenuBar";
import { TableBodySection } from "@/components/sportbook/TableBodySection";
import { TableHeaderCard } from "@/components/sportbook/TableHeaderCard";
import { TableHeaderSection } from "@/components/sportbook/TableHeaderSection";
import { Table } from "@/components/ui/table";
import { useSportbookData } from "@/hooks/useSportBookData";
import { useState } from "react";

function PreMatch() {
  const { preMatch, comingSoon, loading, error } = useSportbookData();
  const [search, setSearch] = useState(false);

  // const groupCountry = Array.from(uniqueMap.values());
  // console.log("groupCountry :", groupCountry);
  // group league

  // diff time

  // search team
  //
  const [scrolled, setScrolled] = useState(false);
  if (!preMatch) return;
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error</p>;
  return (
    <div>
      <MenuBar />

      <TableHeaderCard search={search} setSearch={setSearch} />
      <div className="max-h-[900px] overflow-y-auto">
        <Table className="min-w-[1200px] border border-gray-700">
          <TableHeaderSection scrolled={scrolled} setScrolled={setScrolled} />

          {preMatch.map((match, index) => (
            <TableBodySection
              key={index}
              Away={match.away}
              Home={match.home}
              asianHandicap={match.asian_handicap ?? []}
              country={match.country}
              firstHapdicap={match.asian_handicap_fh ?? []}
              leagueName={match.league}
              overUnderFistHaft={match.over_under_fh ?? []}
              overUnderFullIime={match.over_under_full_time ?? []}
              time={match.timestamp}
            />
          ))}
        </Table>
      </div>
    </div>
  );
}
export default PreMatch;
