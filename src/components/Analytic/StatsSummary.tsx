import React from "react";
import { MatchCardProps } from "./AnalyticCard";

const StatsSummary = ({ match }: MatchCardProps) => {
  return (
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
  );
};

export default StatsSummary;
