"use client";
import MenuBar from "@/components/sportbook/MenuBar";
import { TableBodySection } from "@/components/sportbook/TableBodySection";
import { TableHeaderCard } from "@/components/sportbook/TableHeaderCard";
import { TableHeaderSection } from "@/components/sportbook/TableHeaderSection";
import { Table } from "@/components/ui/table";
import { useSportbookData } from "@/hooks/useSportBookData";
import matchJsonv from "@/app/sportsbook/match_demo.json";
import { useState } from "react";

function PreMatch() {
  const { preMatch, comingSoon, loading, error } = useSportbookData();

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error</p>;

  // const groupCountry = Array.from(uniqueMap.values());
  // console.log("groupCountry :", groupCountry);
  // group league

  // diff time

  // search team
  if (!preMatch) return;
  const firstPreMatch = preMatch[0];
  const [search, setSearch] = useState(false);
  return (
    <div>
      <MenuBar />
      <div className="overflow-y-auto overflow-x-hidden">
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
      </div>
    </div>
  );
}
export default PreMatch;
