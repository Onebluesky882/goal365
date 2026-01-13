// PlayerCard.tsx
"use client";

import React from "react";
import { Player } from "../../../types/player";
import Image from "next/image";
import LoadingIndicators from "../Loading_indicators";
import { formatNumber } from "@/app/common/converts_number";

type Props = {
  player: Player;
  onClick?: () => void;
};

export default function PlayerCard({ player, onClick }: Props) {
  if (!player) return <LoadingIndicators />;
  return (
    <div
      onClick={onClick}
      className="cursor-pointer rounded-2xl bg-white shadow-md p-4 flex flex-col justify-between hover:shadow-lg transition"
    >
      <div className="flex items-center gap-3">
        {/* Avatar */}
        <div className="h-12 w-12 rounded-full bg-gray-200 flex items-center justify-center overflow-hidden">
          {player?.imageUrl ? (
            <Image
              src={player.imageUrl}
              alt={player.name}
              className="h-12 w-12 object-cover"
            />
          ) : (
            <span className="text-gray-400 text-lg">👤</span>
          )}
        </div>

        {/* Name + PlayerNo */}
        <div className="flex flex-col">
          <p className="font-semibold text-gray-800">{player.name}</p>
          <p className="text-sm text-gray-500">#{player.playerNo}</p>
        </div>
      </div>

      {/* Bio */}
      {player.bio && (
        <p className="text-sm text-gray-600 mt-2 line-clamp-2">{player.bio}</p>
      )}

      {/* Wallet / Level */}
      <div className="mt-4 flex justify-between text-sm text-gray-600">
        <span>Wallet: {formatNumber(player.wallet)}</span>
        <span>Lv.{player.level}</span>
      </div>

      {/* Created At */}
      <p className="mt-2 text-xs text-gray-400">
        Created: {new Date(player.createdAt).toLocaleDateString()}
      </p>
    </div>
  );
}
