// components/Myanalytic/myanalytic.tsx
"use client";

import { useEffect, useState } from "react";
import Image from "next/image";
import { AsianHandicap, Match, Score } from "../../../types/myAnalytic";
import { ScoreFullTime } from "./Score";

interface MatchCardProps {
  match: Match;
  onPickChange?: (fixtureId: number, picked: boolean) => void;
}

export default function MatchCard({ match, onPickChange }: MatchCardProps) {
  const [isPicked, setIsPicked] = useState(match.picked || false);
  const [showDetails, setShowDetails] = useState(false);

  const handlePickToggle = () => {
    const newPickedState = !isPicked;
    setIsPicked(newPickedState);
    onPickChange?.(match.fixture_id, newPickedState);
  };

  const formatDate = (timestamp: string) => {
    try {
      const date = new Date(timestamp);
      return date.toLocaleString("th-TH", {
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    } catch {
      return timestamp;
    }
  };

  const formatH2HDate = (timestamp: string) => {
    try {
      const date = new Date(timestamp);
      return date.toLocaleDateString("th-TH", {
        year: "numeric",
        month: "short",
        day: "numeric",
      });
    } catch {
      return timestamp;
    }
  };

  // แปลง form string เป็น array
  const getFormArray = (form: string) => {
    if (!form) return [];
    return form.split("");
  };

  const getFormColor = (result: string) => {
    switch (result.toUpperCase()) {
      case "W":
        return "bg-green-500 text-white";
      case "D":
        return "bg-yellow-500 text-white";
      case "L":
        return "bg-red-500 text-white";
      default:
        return "bg-gray-300 text-gray-700";
    }
  };

  const getFavoriteTeamName = (handicap: AsianHandicap) => {
    if (handicap.line < 0) return match.home; // home ต่อ
    if (handicap.line > 0) return match.away; // away ต่อ
    return "-";
  };
  console.log("match score", match.Score);
  return (
    <div className="bg-background rounded-lg shadow-md hover:shadow-xl transition-shadow border border-gray-100">
      {/* Main Card Content */}
      <div className="p-4">
        {/* Header */}
        <p className="text[8px] text-gray-500">match id: {match.fixture_id}</p>
        <div className="flex justify-between items-start mb-3">
          <div className="text-xs text-gray-500">
            <div className="font-medium text-gray-700">{match.league}</div>
            <div className="text-gray-500">{match.country}</div>
            <div className="text-gray-400 mt-1">
              {formatDate(match.timestamp)}
            </div>
          </div>
          <button
            onClick={handlePickToggle}
            className={`px-4 py-2 rounded-full font-medium transition-all ${
              isPicked
                ? "bg-green-500 text-white hover:bg-green-600 shadow-md"
                : "bg-gray-100 text-gray-700 hover:bg-gray-200"
            }`}
          >
            {isPicked ? "✓ Picked" : "Pick"}
          </button>
        </div>
        {/* Teams */}
        <div className="space-y-3 mb-3">
          {/* Home Team */}
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2 flex-1 min-w-0">
              <div className="w-6 h-6 relative shrink-0">
                <Image
                  src={match.home_logo || "/placeholder-team.png"}
                  alt={match.home}
                  fill
                  className="object-contain"
                  onError={(e) => {
                    e.currentTarget.src = "/placeholder-team.png";
                  }}
                />
              </div>
              <span className="font-medium text-sm truncate">{match.home}</span>
            </div>
            <span className="text-xs font-bold text-blue-600 bg-blue-50 px-2 py-1 rounded ml-2">
              {match.HomeScore}
            </span>
          </div>

          {/* Away Team */}
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2 flex-1 min-w-0">
              <div className="w-6 h-6 relative shrink-0">
                <Image
                  src={match.away_logo || "/placeholder-team.png"}
                  alt={match.away}
                  fill
                  className="object-contain"
                  onError={(e) => {
                    e.currentTarget.src = "/placeholder-team.png";
                  }}
                />
              </div>
              <span className="font-medium text-sm truncate">{match.away}</span>
            </div>
            <span className="text-xs font-bold text-blue-600 bg-blue-50 px-2 py-1 rounded ml-2">
              {match.AwayScore}
            </span>
          </div>
        </div>
        {/* Form Display */}
        <div className="mb-3">
          <div className="space-y-2">
            {/* Home Form */}
            {match.HomeStatic?.form && (
              <div className="flex items-center gap-2">
                <span className="text-xs text-gray-500 w-12">Home:</span>

                {/* scroll container */}
                <div className="flex gap-1 flex-1 min-w-0 overflow-x-auto">
                  {getFormArray(match.HomeStatic.form).map((result, idx) => (
                    <div
                      key={idx}
                      className={`w-6 h-6  shrink-0 flex items-center justify-center rounded text-xs font-bold ${getFormColor(result)}`}
                    >
                      {result}
                    </div>
                  ))}
                </div>
              </div>
            )}

            {/* Away Form */}
            {match.AwayStatic?.form && (
              <div className="flex items-center gap-2">
                <span className="text-xs text-gray-500 w-12">Away:</span>
                <div className="flex gap-1 flex-1 min-w-0 overflow-x-auto">
                  {getFormArray(match.AwayStatic.form).map((result, idx) => (
                    <div
                      key={idx}
                      className={`w-6 h-6  shrink-0 flex items-center justify-center rounded text-xs font-bold ${getFormColor(
                        result,
                      )}`}
                    >
                      {result}
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
        <div>
          <ScoreFullTime
            score={match.Score}
            homeName={match.home}
            awayName={match.away}
          />
        </div>
        {/* Stats Summary */}
        <div className="grid grid-cols-2 gap-2 mb-3 text-xs">
          <div className="bg-blue-300 rounded p-2">
            <div className="text-gray-600 mb-1 font-medium">Home Last 5</div>
            <div className="space-y-1">
              <div className="flex justify-between">
                <span>Form:</span>
                <span className="font-bold text-blue-600">
                  {match.HomeLast5?.form || "N/A"}
                </span>
              </div>
              <div className="flex justify-between">
                <span>Goals:</span>
                <span className="font-medium">
                  {match.HomeLast5?.goals?.for?.total || 0} /{" "}
                  {match.HomeLast5?.goals?.against?.total || 0}
                </span>
              </div>
              <div className="flex justify-between">
                <span>Avg:</span>
                <span className="font-medium">
                  {match.HomeLast5?.goals?.for?.average || "0"} /{" "}
                  {match.HomeLast5?.goals?.against?.average || "0"}
                </span>
              </div>
            </div>
          </div>

          <div className="bg-red-300 rounded p-2">
            <div className="text-gray-600 mb-1 font-medium">Away Last 5</div>
            <div className="space-y-1">
              <div className="flex justify-between">
                <span>Form:</span>
                <span className="font-bold text-red-600">
                  {match.AwayLast5?.form || "N/A"}
                </span>
              </div>
              <div className="flex justify-between">
                <span>Goals:</span>
                <span className="font-medium">
                  {match.AwayLast5?.goals?.for?.total || 0} /{" "}
                  {match.AwayLast5?.goals?.against?.total || 0}
                </span>
              </div>
              <div className="flex justify-between">
                <span>Avg:</span>
                <span className="font-medium">
                  {match.AwayLast5?.goals?.for?.average || "0"} /{" "}
                  {match.AwayLast5?.goals?.against?.average || "0"}
                </span>
              </div>
            </div>
          </div>
        </div>
        {/* Asian Handicap */}
        <div className="mb-3">
          <div className="text-xs text-gray-600 mb-2 font-medium">
            Asian Handicap
          </div>

          <div className="space-y-2">
            {[...(match.AsianHandicap || [])]
              .sort((a, b) => Number(b.is_favorite) - Number(a.is_favorite))
              .map((handicap, idx) => {
                const favTeam = getFavoriteTeamName(handicap);

                return (
                  <div
                    key={idx}
                    className={`
              rounded-lg p-3 border flex justify-between items-center transition-all
              ${
                handicap.is_favorite
                  ? "bg-linear-to-r from-orange-100 to-yellow-100 border-orange-400 shadow-md scale-[1.02]"
                  : "bg-gray-50 border-gray-200 opacity-80"
              }
            `}
                  >
                    {/* LEFT SIDE */}
                    <div className="flex flex-col gap-1">
                      {/* line */}
                      <div
                        className={`text-lg font-bold ${
                          handicap.is_favorite
                            ? "text-orange-600"
                            : "text-gray-600"
                        }`}
                      >
                        {handicap.line > 0 ? "+" : ""}
                        {handicap.line}
                      </div>

                      {/* favorite badge */}
                      {handicap.is_favorite && (
                        <div className="text-[10px] bg-red-400/90 text-white px-2 py-0.5 rounded-full w-fit">
                          Favorite: {favTeam}
                        </div>
                      )}
                    </div>

                    {/* ODDS */}
                    <div className="flex gap-6 text-center">
                      <div>
                        <div className="text-[10px] text-gray-500">Home</div>
                        <div
                          className={` text-[10px] ${
                            handicap.favorite === "Home"
                              ? "text-red-400 "
                              : "text-gray-500"
                          }`}
                        >
                          {handicap.home_odd}
                        </div>
                      </div>

                      <div>
                        <div className="text-[10px] text-gray-500 ">Away</div>
                        <div
                          className={`text-[10px]  ${
                            handicap.favorite === "Away"
                              ? "text-red-400 "
                              : "text-gray-500"
                          }`}
                        >
                          {handicap.away_odd}
                        </div>
                      </div>
                    </div>
                  </div>
                );
              })}
          </div>
        </div>
        {/* View Details Button */}
        <button
          onClick={() => setShowDetails(!showDetails)}
          className="w-full py-2 text-sm font-medium text-blue-600 hover:bg-blue-50 rounded transition-colors"
        >
          {showDetails ? "Hide H2H ▲" : "Show H2H ▼"}
        </button>
      </div>

      {/* H2H Details */}
      {showDetails && match.H2H && match.H2H.length > 0 && (
        <div className="border-t p-4 bg-gray-600">
          <h3 className="text-sm font-bold mb-3 text-gray-300">
            Head to Head ({match.H2H.length} matches)
          </h3>
          <div className="space-y-2 max-h-96 overflow-y-auto">
            {match.H2H.map((h2h, idx) => {
              const isHomeWin = h2h.teams.home.winner;
              const isAwayWin = h2h.teams.away.winner;
              const isDraw = !isHomeWin && !isAwayWin;

              return (
                <div
                  key={idx}
                  className="bg-gray-800 rounded-sm p-3 text-xs border border-gray-400"
                >
                  <div className="flex justify-between items-start mb-2">
                    <div className="text-gray-500">
                      {formatH2HDate(h2h.fixture.date)}
                    </div>
                    <div className="text-gray-500">{h2h.league.round}</div>
                  </div>

                  <div className="space-y-1">
                    {/* Home Team */}
                    <div
                      className={`flex items-center justify-between p-1.5 rounded ${
                        isHomeWin ? "text-blue-500" : ""
                      }`}
                    >
                      <div className="flex items-center gap-2 flex-1 min-w-0">
                        <div className="w-4 h-4 relative shrink-0">
                          <Image
                            src={h2h.teams.home.logo}
                            alt={h2h.teams.home.name}
                            fill
                            className="object-contain"
                          />
                        </div>
                        <span className="font-medium truncate">
                          {h2h.teams.home.name}
                        </span>
                      </div>
                      <span
                        className={`font-bold ml-2 ${
                          isHomeWin ? "text-green-600" : "text-gray-600"
                        }`}
                      >
                        {h2h.goals.home}
                      </span>
                    </div>

                    {/* Away Team */}
                    <div
                      className={`flex items-center justify-between p-1.5 rounded ${
                        isAwayWin ? "text-blue-500" : ""
                      }`}
                    >
                      <div className="flex items-center gap-2 flex-1 min-w-0">
                        <div className="w-4 h-4 relative shrink-0">
                          <Image
                            src={h2h.teams.away.logo}
                            alt={h2h.teams.away.name}
                            fill
                            className="object-contain"
                          />
                        </div>
                        <span className="font-medium truncate">
                          {h2h.teams.away.name}
                        </span>
                      </div>
                      <span
                        className={`font-bold ml-2 ${
                          isAwayWin ? "text-green-600" : "text-gray-600"
                        }`}
                      >
                        {h2h.goals.away}
                      </span>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      )}
    </div>
  );
}
