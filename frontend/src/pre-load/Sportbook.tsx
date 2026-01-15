"use client";
import { MatchHeader } from "@/components/MatchTable/MatchHeader";
import { MatchRows, TableHeaderRows } from "@/components/MatchTable/MatchRows";
import { Table, TableBody, TableHeader } from "@/components/ui/table";

import React, { useState } from "react";

export default function BettingOddsDisplay() {
  const [matchesData] = useState<BookmakerData[]>([
    {
      id: 1,
      matchInfo: {
        isLive: true,
        time: "21:00",
        home: "Barcelona BA",
        away: "Galícia",
        league: "BRAZIL BAIANO - 1",
        score: "0 - 0",
      },
      markets: {
        matchWinner: {
          type: "MATCH_WINNER",
          values: [
            { value: "Home", odd: "3.15" },
            { value: "Draw", odd: "3.20" },
            { value: "Away", odd: "2.42" },
          ],
        },
        handicap: {
          type: "HDP",
          values: [
            { value: "Home +0", odd: "2.12" },
            { value: "Away +0", odd: "1.80" },
            { value: "Home +0.25", odd: "1.82" },
            { value: "Away +0.25", odd: "2.10" },
            { value: "Home +0.5", odd: "1.56" },
            { value: "Away +0.5", odd: "2.36" },
          ],
        },
        overUnder: {
          type: "OU",
          values: [
            { value: "Over 2.25", odd: "1.86" },
            { value: "Under 2.25", odd: "2.04" },
            { value: "Over 2.5", odd: "2.07" },
            { value: "Under 2.5", odd: "1.83" },
          ],
        },
      },
    },
    {
      id: 2,
      matchInfo: {
        isLive: false,
        time: "22:30",
        home: "Manchester United",
        away: "Liverpool",
        league: "ENGLAND PREMIER LEAGUE",
        score: "0 - 0",
      },
      markets: {
        matchWinner: {
          type: "MATCH_WINNER",
          values: [
            { value: "Home", odd: "2.91" },
            { value: "Draw", odd: "3.15" },
            { value: "Away", odd: "2.38" },
          ],
        },
        handicap: {
          type: "HDP",
          values: [
            { value: "Home +0", odd: "1.95" },
            { value: "Away +0", odd: "1.95" },
            { value: "Home +0.25", odd: "1.75" },
            { value: "Away +0.25", odd: "2.15" },
            { value: "Home +0.5", odd: "1.60" },
            { value: "Away +0.5", odd: "2.40" },
          ],
        },
        overUnder: {
          type: "OU",
          values: [
            { value: "Over 2.25", odd: "1.90" },
            { value: "Under 2.25", odd: "2.00" },
            { value: "Over 2.5", odd: "2.10" },
            { value: "Under 2.5", odd: "1.80" },
          ],
        },
      },
    },
    {
      id: 3,
      matchInfo: {
        isLive: true,
        time: "01:30",
        home: "Real Madrid",
        away: "Barcelona",
        league: "SPAIN LA LIGA",
        score: "1 - 1",
      },
      markets: {
        matchWinner: {
          type: "MATCH_WINNER",
          values: [
            { value: "Home", odd: "2.65" },
            { value: "Draw", odd: "3.10" },
            { value: "Away", odd: "2.55" },
          ],
        },
        handicap: {
          type: "HDP",
          values: [
            { value: "Home +0", odd: "2.05" },
            { value: "Away +0", odd: "1.85" },
            { value: "Home +0.25", odd: "1.88" },
            { value: "Away +0.25", odd: "2.02" },
            { value: "Home +0.5", odd: "1.68" },
            { value: "Away +0.5", odd: "2.22" },
          ],
        },
        overUnder: {
          type: "OU",
          values: [
            { value: "Over 2.25", odd: "1.85" },
            { value: "Under 2.25", odd: "2.05" },
            { value: "Over 2.5", odd: "2.05" },
            { value: "Under 2.5", odd: "1.85" },
          ],
        },
      },
    },
  ]);

  const [selectedBet, setSelectedBet] = useState<any>(null);

  const handleOddsClick =
    (matchId: number, matchInfo: MatchInfo) =>
    (market: string, selection: string, odd: string) => {
      const betData = {
        matchId,
        match: `${matchInfo.home} vs ${matchInfo.away}`,
        market,
        selection,
        odd,
      };
      console.log("Selected Odds:", betData);
      setSelectedBet(betData);
    };

  return (
    <div className="">
      {/* Matches List */}
      <div className="space-y-1">
        {matchesData.map((data) => (
          <div
            key={data.id}
            className="bg-gray-900  overflow-hidden border border-gray-800"
          >
            <MatchHeader matchInfo={data.matchInfo} />

            <div className="overflow-x-auto">
              <Table className="min-w-125">
                <TableHeader>
                  <TableHeaderRows />
                </TableHeader>
                <TableBody>
                  <MatchRows
                    data={data}
                    onOddsClick={handleOddsClick(data.id, data.matchInfo)}
                  />
                </TableBody>
              </Table>
            </div>
          </div>
        ))}
      </div>

      {/* Selected Bet Footer */}
      {selectedBet && (
        <div className="fixed text[12px] bottom-10 left-0 right-0 bg-gray-900 border-t-2 border-green-500 shadow-lg ">
          <div className="flex justify-between items-center mb-2">
            <div className="flex-1">
              <p className="text-xs text-gray-400">{selectedBet.market}</p>
              <p className="font-bold text-sm text-white">
                {selectedBet.match}
              </p>
              <p className="text-xs text-green-400">{selectedBet.selection}</p>
            </div>
            <div className="text-right">
              <p className="text-xs text-gray-400">ราคา</p>
              <p className="text-2xl font-bold text-green-400">
                {selectedBet.odd}
              </p>
            </div>
          </div>
          <button className="w-full bg-green-600 text-white py-2 rounded font-bold text-sm active:bg-green-700">
            วางเดิมพัน
          </button>
        </div>
      )}
    </div>
  );
}
