import { Score as S } from "@/types/myAnalytic";

type Props = {
  score?: S; // optional กัน undefined
  homeName?: string;
  awayName?: string;
};

export function ScoreFullTime({ score, homeName, awayName }: Props) {
  const home = score?.fulltime?.home ?? "-";
  const away = score?.fulltime?.away ?? "-";

  return (
    <div className="flex items-center justify-between bg-gray-400 rounded-sm my-2 px-4 py-2 font-bold text-sm">
      <span className="truncate">{homeName ?? "-"}</span>

      <div className="text-sm px-3 flex flex-col text-center">
        <span className="text-[9px] text-red-700">FT</span>
        <div>
          {home} - {away}
        </div>
      </div>

      <span className="truncate text-right">{awayName ?? "-"}</span>
    </div>
  );
}
