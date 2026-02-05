import React from "react";
import { MatchCardProps } from "./MyAnalytic";
import Image from "next/image";
type Props = MatchCardProps & {
  formatH2HDate: (timestamp: string) => string;
};

const H2HPreviews = ({ formatH2HDate, match }: Props) => {
  if (!match.H2H || match.H2H.length === 0) return null;
  return (
    <div className="border-t p-4 bg-gray-600">
      <h3 className="text-sm font-bold mb-3 text-gray-300">
        Head to Head ({match.H2H.length} matches)
      </h3>
      <div className="space-y-2 max-h-96 overflow-y-auto">
        {match.H2H.map((h2h, idx) => {
          const isHomeWin = h2h.teams.home.winner;
          const isAwayWin = h2h.teams.away.winner;

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
  );
};

export default H2HPreviews;
