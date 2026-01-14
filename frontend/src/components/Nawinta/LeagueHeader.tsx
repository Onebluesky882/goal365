import { League } from "../../../types/predictions";

type Props = {
  league: League;
};

export default function LeagueHeader({ league }: Props) {
  return (
    <>
      {league && (
        <div className="flex items-center gap-3">
          <img
            src={league?.logo ?? ""}
            alt={league?.name ?? ""}
            className="h-8"
          />
          <div>
            <h2 className="font-bold">{league?.name}</h2>
            <p className="text-xs text-gray-500">
              {league.country} • Season {league.season}
            </p>
          </div>
        </div>
      )}
    </>
  );
}
