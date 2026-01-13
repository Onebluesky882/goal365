"use client";

import { Plus } from "lucide-react";
import { Player } from "../../../types/player";
import { getUsageAge } from "@/app/common/getUsageAge";

type PlayerProps = {
  newPlayerLink: string;
  players: Player[];
  canCreate: boolean;
  onCreate: () => void;
  onClick: (playerNo: number) => void;
};

export default function PlayersGrid({
  players,
  canCreate,
  onCreate,
  onClick,
}: PlayerProps) {
  return (
    <div className="grid grid-cols-2 md:grid-cols-4 gap-4 p-2">
      {players &&
        players.map((p) => (
          <div
            onClick={() => onClick(p.playerNo)}
            key={p.playerNo}
            className="rounded-2xl bg-white shadow-md p-4 flex flex-col justify-between"
          >
            <div className="flex items-center gap-3">
              <div className="h-12 w-12 rounded-full bg-gray-200 flex items-center justify-center">
                {p.imageUrl ? (
                  <img
                    src={p.imageUrl}
                    alt={p.name}
                    className="h-12 w-12 rounded-full object-cover"
                  />
                ) : (
                  <span className="text-gray-400 text-lg">👤</span>
                )}
              </div>

              <div>
                <p className="font-semibold">{p.playerNo}</p>
                <p className="font-semibold">{p.name}</p>
                <p className="text-xs text-gray-500">
                  ใช้งาน {getUsageAge(p.createdAt)}
                </p>
              </div>

              {p.bio && (
                <p className="text-sm text-gray-600 mt-2 line-clamp-2">
                  {p.bio}
                </p>
              )}
            </div>

            <div className="mt-4 flex justify-between text-sm text-gray-600">
              <span>กระเป๋าตัง {p.wallet}</span>
              <span>Lv.{p.level}</span>
            </div>
          </div>
        ))}
      {/* ➕ Create Player */}
      <button
        className={`flex  flex-col items-center justify-center rounded-2xl border-2 border-dashed border-gray-200 
        hover:border-gray-400 hover:bg-gray-150 transition min-h-40 ${
          canCreate ? "hover:bg-gray-100" : "opacity-50 cursor-not-allowed"
        }`}
        onClick={onCreate}
        disabled={!canCreate}
      >
        <p className="text-gray-400">{canCreate ? "New" : "Next level : 10"}</p>
        <Plus size={32} className="text-gray-300" />
      </button>
    </div>
  );
}
