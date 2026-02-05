import { formatDate } from "@/common/format_date";
import { MatchCardProps } from "./Myanalytic";

export const AnalyserHeader = ({
  match,
  handlePickToggle,
  isPicked,
}: MatchCardProps & {
  handlePickToggle: () => void;
}) => {
  return (
    <>
      <div className="flex justify-center">
        <p className="text-[12px] text-gray-500 text-center">
          Match id: {match.fixture_id}
        </p>
      </div>
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
    </>
  );
};
