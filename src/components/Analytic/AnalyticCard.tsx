"use client";

import { useState } from "react";
import { ScoreFullTime } from "./Score";
import { AnalyserHeader } from "./AnalyserHeader";
import StatsSummary from "./StatsSummary";
import { Match } from "@/types/myAnalytic";
import AsianHandicapSection from "./AsianHandicapSection";
import FormDisplay from "./FormDisplay";
import Teams from "./Teams";
import H2HPreviews from "./H2HPreviews";

export type MatchCardProps = {
  date?: string;
  match: Match;
  isPicked?: boolean;
  handlePickToggle?: (fixtureId: number, picked: boolean) => Promise<void>;
};

export default function MatchCard({ match, handlePickToggle }: MatchCardProps) {
  const [showDetails, setShowDetails] = useState(false);

  // ✅ ใช้ timestamp number (ถูกต้องกับ API)
  const formatH2HDate = (timestamp: number) => {
    try {
      return new Date(timestamp * 1000).toLocaleDateString("th-TH", {
        year: "numeric",
        month: "short",
        day: "numeric",
      });
    } catch {
      return "";
    }
  };

  const getFormArray = (form: string) => {
    if (!form) return [];
    return form.split("").reverse();
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

  // ✅ safe handler (แก้ !)
  const onPickToggle = async () => {
    if (!handlePickToggle) return;
    await handlePickToggle(match.fixture_id, !match.picked);
  };

  return (
    <div
      className="
        bg-background rounded-lg shadow-md hover:shadow-xl transition-shadow border border-gray-100
            "
    >
      {/* Main Card Content */}
      <div className="p-4 max-sm:">
        {/* Header */}
        <AnalyserHeader match={match} handlePickToggle={onPickToggle} />

        {/* Teams */}
        <div className="max-sm:scale-[0.95] max-sm:origin-top">
          <Teams match={match} />
        </div>

        {/* Form Display */}
        <div className="max-sm:mt-1">
          <FormDisplay
            match={match}
            getFormArray={getFormArray}
            getFormColor={getFormColor}
          />
        </div>

        {/* Score */}
        <div className="max-sm:mt-1 max-sm:text-sm">
          <ScoreFullTime
            score={match.score}
            homeName={match.home}
            awayName={match.away}
          />
        </div>

        {/* Stats */}
        <div className="max-sm:mt-2 max-sm:scale-[0.95]">
          <StatsSummary match={match} />
        </div>

        {/* Handicap */}
        <div className="max-sm:mt-2 max-sm:scale-[0.95]">
          <AsianHandicapSection match={match} />
        </div>

        {/* Button */}
        <button
          onClick={() => setShowDetails(!showDetails)}
          className="
            w-full py-2 text-sm font-medium text-blue-600 hover:bg-blue-50 rounded transition-colors
            max-sm:py-1.5 max-sm:text-xs max-sm:rounded-md
          "
        >
          {showDetails ? "Hide H2H ▲" : "Show H2H ▼"}
        </button>
      </div>

      {/* H2H */}
      {showDetails && (match.H2H?.length ?? 0) > 0 && (
        <div className="max-sm:text-xs">
          <H2HPreviews match={match} formatH2HDate={formatH2HDate} />
        </div>
      )}
    </div>
  );
}
