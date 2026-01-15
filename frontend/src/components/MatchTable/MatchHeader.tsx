"use client";

export const MatchHeader = ({ matchInfo }: { matchInfo: MatchInfo }) => {
  return (
    <div className="bg-gray-800 border-l-[3px] border-orange-500 px-2 py-1 mb-0.5 flex justify-between items-center">
      <h2 className="text-[9px] font-bold uppercase tracking-tighter text-white truncate">
        {matchInfo.league}
      </h2>
      {matchInfo.isLive && (
        <span className="text-[7px] text-orange-500 font-black animate-pulse uppercase tracking-widest">
          Live Now
        </span>
      )}
    </div>
  );
};
