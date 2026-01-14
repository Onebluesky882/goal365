"use client";

import { Plus } from "lucide-react";
import { Player } from "../../../types/player";
import { getUsageAge } from "@/app/common/getUsageAge";
import Link from "next/link";

type PlayerProps = {
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
  // Helper: สร้างตัวย่อชื่อสำหรับ avatar
  const getInitials = (name: string) => {
    return name
      .split(" ")
      .map((n) => n[0])
      .join("")
      .toUpperCase()
      .substring(0, 2);
  };

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 p-3">
      {players.map((p) => (
        <div
          key={p.playerNo}
          onClick={() => onClick(p.playerNo)}
          className="rounded-xl bg-white shadow-sm border border-gray-100 overflow-hidden transition-all hover:shadow-md active:scale-[0.99] cursor-pointer"
        >
          <div className="flex items-start gap-3 p-4">
            {/* Avatar */}
            <div className="relative shrink-0">
              {p.imageUrl ? (
                <img
                  src={p.imageUrl}
                  alt={p.name}
                  className="h-14 w-14 rounded-full object-cover border-2 border-white shadow-sm"
                />
              ) : (
                <div className="h-14 w-14 rounded-full bg-linear-to-br from-blue-400 to-indigo-500 flex items-center justify-center text-white font-semibold text-sm">
                  {getInitials(p.name)}
                </div>
              )}
            </div>

            {/* Info */}
            <div className="flex-1 min-w-0">
              <div className="flex items-baseline gap-2">
                <h3 className="font-bold text-gray-900 truncate">{p.name}</h3>
                <span className="text-xs text-gray-500 bg-gray-100 px-2 py-0.5 rounded-full">
                  #{p.playerNo}
                </span>
              </div>
              <p className="text-xs text-gray-500 mt-1">
                ใช้งานมาแล้ว {getUsageAge(p.createdAt)}
              </p>

              {p.bio && (
                <p className="text-sm text-gray-600 mt-2 line-clamp-2">
                  {p.bio}
                </p>
              )}

              <div className="mt-3 flex justify-between items-center text-xs text-gray-600">
                <span className="font-medium">กระเป๋า: {p.wallet}</span>
                <span className="bg-indigo-50 text-indigo-700 px-2 py-0.5 rounded-full">
                  Lv.{p.level}
                </span>
              </div>
            </div>
          </div>
        </div>
      ))}

      {/* ➕ Create Player */}
      <button
        onClick={onCreate}
        disabled={!canCreate}
        className={`flex flex-col items-center justify-center rounded-xl border-2 border-dashed ${
          canCreate
            ? "border-gray-300 hover:border-indigo-400 hover:bg-indigo-50 text-gray-600"
            : "border-gray-200 text-gray-400 opacity-70 cursor-not-allowed"
        } transition-colors p-6 min-h-30`}
      >
        {canCreate ? (
          <>
            <Link href={"/new-player"}>
              <Plus size={24} className="mb-2 text-gray-500" />
              <span className="text-sm font-medium">สร้างตัวละครใหม่</span>
            </Link>
          </>
        ) : (
          <>
            <span className="text-xs mb-1">ปลดล็อกที่เลเวล 10</span>
            <Plus size={24} className="text-gray-300" />
          </>
        )}
      </button>
    </div>
  );
}
