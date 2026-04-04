import { MatchCardProps } from "./MyAnalytic";

const AsianHandicapSection = ({ match }: MatchCardProps) => {
  return (
    <div className="mb-3">
      <div className="text-xs text-gray-500 mb-2 font-medium">
        Asian Handicap
      </div>

      <div className="space-y-2">
        {[...(match.AsianHandicap || [])]
          .sort((a, b) => Number(b.is_favorite) - Number(a.is_favorite))
          .map((handicap, idx) => {
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
                        ? "text-orange-600 "
                        : "text-gray-600"
                    }`}
                  >
                    {handicap.line > 0 ? "+" : ""}
                    {handicap.line}
                  </div>
                </div>

                {/* ODDS */}
                <div className="flex gap-6 text-center">
                  <div>
                    <div className="text-[10px] text-gray-500">Home</div>
                    <div
                      className={` text-[10px] ${
                        handicap.favorite === "Home"
                          ? "text-red-400 font-bold "
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
                          ? "text-red-400 font-bold"
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
  );
};

export default AsianHandicapSection;
