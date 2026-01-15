"use client";

import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { toast } from "sonner";

// --- Data Section ---
const sboHandicapData = {
  values: [
    { value: "Home +0", odd: "2.12" },
    { value: "Away +0", odd: "1.80" },
    { value: "Home +0.25", odd: "1.82" },
    { value: "Away +0.25", odd: "2.10" },
    { value: "Home +0.5", odd: "1.56" },
    { value: "Away +0.5", odd: "2.36" },
  ],
};

const matchWinner = {
  values: [
    { value: "Home", odd: "3.15" },
    { value: "Draw", odd: "3.20" },
    { value: "Away", odd: "2.42" },
  ],
};

const sboOverUnder = {
  values: [
    { value: "Over 2.25", odd: "1.86" },
    { value: "Under 2.25", odd: "2.04" },
    { value: "Over 2.5", odd: "2.07" },
    { value: "Under 2.5", odd: "1.83" },
  ],
};

const matchInfo = {
  isLive: true, // ✅ เปลี่ยนเป็น false เพื่อดูเวลาปกติ
  time: "21:00",
  home: "Barcelona BA",
  away: "Galícia",
  league: "BRAZIL BAIANO - 1",
  score: "0 - 0",
};

export default function UltraCompactClassicTable() {
  const handleOddsClick = (market: string, selection: string, odd: string) => {
    const logData = {
      market,
      selection,
      odd,
      match: `${matchInfo.home} vs ${matchInfo.away}`,
    };
    console.log("Selected Odds:", logData);
    toast.success(`เลือก: ${selection} (${market}) ราคา ${odd}`);
  };

  const getVal = (data: any, key: string) =>
    data.values.find((v: any) => v.value === key)?.odd || "-";

  const isHomeFav = sboHandicapData.values.some((v) =>
    v.value.includes("Home -")
  );
  const isAwayFav = sboHandicapData.values.some((v) =>
    v.value.includes("Away -")
  );

  const OddsButton = ({
    market,
    selection,
    odd,
    color = "text-foreground",
  }: any) => (
    <button
      onClick={() => odd !== "-" && handleOddsClick(market, selection, odd)}
      className={`w-full h-full font-mono font-bold text-[10px] transition-all active:scale-90 hover:bg-[#00acec]/20 ${color} ${
        odd === "-" ? "opacity-20 cursor-default" : "cursor-pointer"
      }`}
    >
      {odd}
    </button>
  );

  return (
    <div className="w-full bg-background font-sans select-none overflow-hidden ">
      {/* 🟠 Header */}
      <div className="bg-secondary border-l-[3px] border-[#ff5f00] px-2 py-1 mb-0.5 flex justify-between items-center">
        <h2 className="text-[9px] font-bold uppercase tracking-tighter text-foreground truncate">
          {matchInfo.league}
        </h2>
        {matchInfo.isLive && (
          <span className="text-[7px] text-[#ff5f00] font-black animate-pulse uppercase tracking-widest">
            Live Now
          </span>
        )}
      </div>

      <div className="overflow-x-auto border border-border bg-card scrollbar-hide">
        <Table className="min-w-125 border-collapse table-fixed">
          <TableHeader className="bg-card">
            <TableRow className="border-b border-border hover:bg-transparent h-7">
              {/* ✅ ปรับหัวคอลัมน์แรกให้ยืดหยุ่น */}
              <TableHead className="w-10 text-center text-[7px] font-bold border-r border-border p-0 uppercase">
                {matchInfo.isLive ? "Score" : "Time"}
              </TableHead>
              <TableHead className="w-30 text-left text-[7px] font-bold px-2 border-r border-border uppercase">
                Match
              </TableHead>
              <TableHead className="w-10 text-center text-[7px] font-bold border-r border-border p-0">
                1X2
              </TableHead>
              <TableHead
                colSpan={3}
                className="text-center text-[7px] font-bold border-r border-border p-0 uppercase"
              >
                Handicap
              </TableHead>
              <TableHead
                colSpan={3}
                className="text-center text-[7px] font-bold p-0 uppercase"
              >
                Over/User
              </TableHead>
            </TableRow>
            <TableRow className="bg-secondary/40 border-b border-border h-5 text-[7px] uppercase opacity-50">
              <TableCell className="border-r border-border p-0 text-center">
                {matchInfo.isLive ? "LIVE" : ""}
              </TableCell>
              <TableCell className="border-r border-border p-0"></TableCell>
              <TableCell className="text-center border-r border-border p-0">
                1X2
              </TableCell>
              <TableCell className="text-center border-r border-border">
                HDP
              </TableCell>
              <TableCell className="text-center border-r border-border">
                H
              </TableCell>
              <TableCell className="text-center border-r border-border">
                A
              </TableCell>
              <TableCell className="text-center border-r border-border">
                O/U
              </TableCell>
              <TableCell className="text-center border-r border-border">
                O
              </TableCell>
              <TableCell className="text-center">U</TableCell>
            </TableRow>
          </TableHeader>

          <TableBody>
            {/* Row 1: Home */}
            <TableRow className="border-b border-border h-9">
              {/* ✅ คอลัมน์แรก: แสดง สกอร์ ถ้า Live, แสดง เวลา ถ้าไม่ Live */}
              <TableCell
                className={`text-center font-bold text-[9px] border-r border-border p-0 ${
                  matchInfo.isLive ? "text-[#ff5f00]" : "text-muted-foreground"
                }`}
              >
                {matchInfo.isLive ? matchInfo.score : matchInfo.time}
              </TableCell>
              <TableCell className="px-2 border-r border-border">
                <span
                  className={`text-[10px] font-bold truncate block ${
                    isHomeFav ? "text-[#ff4d4d]" : "text-white"
                  }`}
                >
                  {matchInfo.home}
                </span>
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="1X2"
                  selection="Home"
                  odd={getVal(matchWinner, "Home")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center text-[9px] text-muted-foreground border-r border-border italic p-0">
                0
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Home"
                  odd={getVal(sboHandicapData, "Home +0")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Away"
                  odd={getVal(sboHandicapData, "Away +0")}
                />
              </TableCell>
              <TableCell className="text-center text-[9px] text-muted-foreground border-r border-border italic p-0">
                2.25
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="OU"
                  selection="Over"
                  odd={getVal(sboOverUnder, "Over 2.25")}
                />
              </TableCell>
              <TableCell className="text-center p-0">
                <OddsButton
                  market="OU"
                  selection="Under"
                  odd={getVal(sboOverUnder, "Under 2.25")}
                />
              </TableCell>
            </TableRow>

            {/* Row 2: Draw / VS */}
            <TableRow className="border-b border-border h-9 bg-secondary/10">
              <TableCell className="border-r border-border p-0 text-center text-[7px] text-[#ff5f00] font-black italic">
                {matchInfo.isLive ? "75'" : ""}
              </TableCell>
              <TableCell className="px-2 border-r border-border text-center text-[8px] font-bold text-muted-foreground italic uppercase">
                vs
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="1X2"
                  selection="Draw"
                  odd={getVal(matchWinner, "Draw")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center text-[9px] text-muted-foreground border-r border-border italic p-0">
                0.25
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Home"
                  odd={getVal(sboHandicapData, "Home +0.25")}
                />
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Away"
                  odd={getVal(sboHandicapData, "Away +0.25")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center text-[9px] text-muted-foreground border-r border-border italic p-0">
                2.5
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="OU"
                  selection="Over"
                  odd={getVal(sboOverUnder, "Over 2.5")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center p-0">
                <OddsButton
                  market="OU"
                  selection="Under"
                  odd={getVal(sboOverUnder, "Under 2.5")}
                />
              </TableCell>
            </TableRow>

            {/* Row 3: Away */}
            <TableRow className="border-b-2 border-border h-9">
              <TableCell className="border-r border-border p-0"></TableCell>
              <TableCell className="px-2 border-r border-border">
                <span
                  className={`text-[10px] font-bold truncate block ${
                    isAwayFav ? "text-[#ff4d4d]" : "text-white"
                  }`}
                >
                  {matchInfo.away}
                </span>
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="1X2"
                  selection="Away"
                  odd={getVal(matchWinner, "Away")}
                  color="text-[#00acec]"
                />
              </TableCell>
              <TableCell className="text-center text-[9px] text-muted-foreground border-r border-border italic p-0">
                0.5
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Home"
                  odd={getVal(sboHandicapData, "Home +0.5")}
                />
              </TableCell>
              <TableCell className="text-center border-r border-border p-0">
                <OddsButton
                  market="HDP"
                  selection="Away"
                  odd={getVal(sboHandicapData, "Away +0.5")}
                />
              </TableCell>
              <TableCell
                colSpan={3}
                className="text-center p-0 italic text-[7px] opacity-20 uppercase font-black"
              >
                Next Markets Soon
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <div className="flex justify-center mt-2">
        <span className="text-[6px] text-muted-foreground uppercase tracking-widest border border-border px-2 py-0.5 rounded-full opacity-40">
          SBO Real-time Data Sync
        </span>
      </div>
    </div>
  );
}
