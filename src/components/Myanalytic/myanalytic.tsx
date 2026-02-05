// components/Myanalytic/myanalytic.tsx
"use client";

import { useState } from "react";
import { ScoreFullTime } from "./Score";
import { AnalyserHeader } from "./AnalyserHeader";
import StatsSummary from "./StatsSummary";
import { AsianHandicap, Match } from "../../../types/myAnalytic";
import AsianHandicapSection from "./AsianHandicapSection";
import FormDisplay from "./FormDisplay";
import Teams from "./Teams";
import H2HPreviews from "./H2HPreviews";

export type MatchCardProps = {
  match: Match;
  isPicked?: boolean;
  onPickChange?: (fixtureId: number, picked: boolean) => void;
  getFavoriteTeamName?: (handicap: AsianHandicap) => void;
  getFormArray?: (form: string) => string[];
  getFormColor?: (result: string) => string;
};

export default function MatchCard({ match, onPickChange }: MatchCardProps) {
  const [isPicked, setIsPicked] = useState(match.picked || false);
  const [showDetails, setShowDetails] = useState(false);

  const handlePickToggle = () => {
    const newPickedState = !isPicked;
    setIsPicked(newPickedState);
    onPickChange?.(match.fixture_id, newPickedState);
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
  return (
    <div className="bg-background rounded-lg shadow-md hover:shadow-xl transition-shadow border border-gray-100">
      {/* Main Card Content */}
      <div className="p-4">
        {/* Header */}

        <AnalyserHeader
          match={match}
          handlePickToggle={handlePickToggle}
          isPicked={isPicked}
        />

        {/* Teams */}
        <Teams match={match} />

        {/* Form Display */}
        <FormDisplay
          match={match}
          getFormArray={getFormArray}
          getFormColor={getFormColor}
        />

        <ScoreFullTime
          score={match.score}
          homeName={match.home}
          awayName={match.away}
        />

        {/* Stats Summary */}
        <StatsSummary match={match} isPicked={isPicked} />

        {/* Asian Handicap */}
        <AsianHandicapSection
          match={match}
          getFavoriteTeamName={getFavoriteTeamName}
        />
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
        <H2HPreviews match={match} formatH2HDate={formatH2HDate} />
      )}
    </div>
  );
}
