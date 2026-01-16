"use client";

import { TableCell, TableHead, TableRow } from "../ui/table";
import { OddsButton } from "./TableSportBook";

export type OddsClickHandler = (
  market: string,
  selection: string,
  odd: string
) => void;

export const MatchRows = ({
  data,
  onOddsClick,
}: {
  data: BookmakerData;
  onOddsClick: OddsClickHandler;
}) => {
  const { matchInfo, markets } = data;

  const getVal = (market: Market, key: string) =>
    market.values.find((v) => v.value === key)?.odd || "-";

  const isHomeFav = markets.handicap.values.some((v) =>
    v.value.includes("Home -")
  );
  const isAwayFav = markets.handicap.values.some((v) =>
    v.value.includes("Away -")
  );

  return (
    <>
      {/* Row 1: Home */}
      <TableRow className="">
        <TableCell
          className={`text-center font-bold text-[9px] border-r border-gray-700 p-0 ${
            matchInfo.isLive ? "text-red-400" : "text-gray-400"
          }`}
        >
          {matchInfo.isLive ? matchInfo.score : matchInfo.time}
        </TableCell>
        <TableCell className="px-2 border-r border-gray-700">
          <span
            className={`text-[10px] font-bold truncate block ${
              isHomeFav ? "text-red-400" : "text-white"
            }`}
          >
            {matchInfo.home}
          </span>
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="1X2"
            selection="Home"
            odd={getVal(markets.matchWinner, "Home")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center text-[9px] text-gray-400 border-r border-gray-700 italic p-0">
          0
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Home +0"
            odd={getVal(markets.handicap, "Home +0")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Away +0"
            odd={getVal(markets.handicap, "Away +0")}
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center text-[9px] text-gray-400 border-r border-gray-700 italic p-0">
          2.25
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="OU"
            selection="Over 2.25"
            odd={getVal(markets.overUnder, "Over 2.25")}
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center p-0">
          <OddsButton
            market="OU"
            selection="Under 2.25"
            odd={getVal(markets.overUnder, "Under 2.25")}
            onClick={onOddsClick}
          />
        </TableCell>
      </TableRow>

      {/* Row 2: Draw / VS */}
      <TableRow className="h-9 bg-gray-800/10">
        <TableCell className="border-r border-gray-700 p-0 text-center text-[7px] text-red-500 font-black  ">
          {matchInfo.isLive ? "75'" : ""}
        </TableCell>
        <TableCell className="px-2 border-r border-gray-700  text-[6px] font-bold text-gray-600 italic uppercase">
          vs
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="1X2"
            selection="Draw"
            odd={getVal(markets.matchWinner, "Draw")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center text-[9px] text-gray-400 border-r border-gray-700 italic p-0">
          0.25
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Home +0.25"
            odd={getVal(markets.handicap, "Home +0.25")}
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Away +0.25"
            odd={getVal(markets.handicap, "Away +0.25")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center text-[9px] text-gray-400 border-r border-gray-700 italic p-0">
          2.5
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="OU"
            selection="Over 2.5"
            odd={getVal(markets.overUnder, "Over 2.5")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center p-0">
          <OddsButton
            market="OU"
            selection="Under 2.5"
            odd={getVal(markets.overUnder, "Under 2.5")}
            onClick={onOddsClick}
          />
        </TableCell>
      </TableRow>

      {/* Row 3: Away */}
      <TableRow className="border-b-2 border-gray-700 h-9">
        <TableCell className="border-r border-gray-700 p-0">
          <span className="text-[6px] text-green-500 uppercase tracking-widest  border-yellow-600  px-2 py-0.5 rounded-full ">
            odds+
          </span>
        </TableCell>
        <TableCell className="px-2 border-r border-gray-700">
          <span
            className={`text-[10px] font-bold truncate block ${
              isAwayFav ? "text-red-400" : "text-white"
            }`}
          >
            {matchInfo.away}
          </span>
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="1X2"
            selection="Away"
            odd={getVal(markets.matchWinner, "Away")}
            color="text-blue-400"
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center text-[9px] text-gray-400 border-r border-gray-700 italic p-0">
          0.5
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Home +0.5"
            odd={getVal(markets.handicap, "Home +0.5")}
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-0">
          <OddsButton
            market="HDP"
            selection="Away +0.5"
            odd={getVal(markets.handicap, "Away +0.5")}
            onClick={onOddsClick}
          />
        </TableCell>
        <TableCell
          colSpan={3}
          className="text-center p-0 italic text-[7px] opacity-20 uppercase font-black text-gray-500"
        >
          Next Markets Soon
        </TableCell>
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
