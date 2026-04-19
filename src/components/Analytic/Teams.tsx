import Image from "next/image";
import { MatchCardProps } from "./AnalyticCard";

const Teams = ({ match }: MatchCardProps) => {
  return (
    <div className="space-y-3 mb-3">
      {/* Home Team */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2 flex-1 min-w-0">
          <div className="w-6 h-6 relative shrink-0">
            <Image
              src={match.home_logo || "/placeholder-team.png"}
              alt={match.home}
              fill
              sizes="50px"
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
              sizes="50px"
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
  );
};

export default Teams;
