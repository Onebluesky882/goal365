"use client";

import React from "react";
import { Player } from "../../../types/player";
import { formatNumber } from "@/app/common/converts_number";
import { Wallet, Plus, User } from "lucide-react";

type Props = {
  player: Player | null;
};

export default function WalletMinibar({ player }: Props) {
  if (!player) return null;

  return (
    <div className="sticky top-0 z-50 w-full px-4  flex justify-center pointer-events-none">
      <div className="pointer-events-auto flex items-center gap-4 bg-[var(--secondary)]/80 backdrop-blur-md border border-[var(--border)] px-4 py-2 rounded-full shadow-2xl shadow-black/50">
        {/* User Info Section */}
        <div className="flex items-center gap-2 pr-3 border-r border-[var(--border)]">
          <div className="h-7 w-7 rounded-full bg-[var(--card)] border border-[var(--border)] overflow-hidden flex items-center justify-center">
            {player.imageUrl ? (
              <img
                src={player.imageUrl}
                alt={player.name}
                className="h-full w-full object-cover"
              />
            ) : (
              <User size={14} className="text-[var(--muted-foreground)]" />
            )}
          </div>
          <div className="flex flex-col">
            <span className="text-[10px] font-bold text-[var(--foreground)] truncate max-w-[80px]">
              {player.name}
            </span>
            <span className="text-[8px] text-[#00acec] font-bold uppercase tracking-tighter">
              Lv.{player.level}
            </span>
          </div>
        </div>

        {/* Wallet Section */}
        <div className="flex items-center gap-3">
          <div className="flex items-center gap-2">
            <div className="p-1 rounded-md bg-[#ff5f00]/10">
              <Wallet size={14} className="text-[#ff5f00]" />
            </div>
            <div className="flex flex-col">
              <span className="text-[9px] text-[var(--muted-foreground)] leading-none mb-0.5">
                ยอดเงินคงเหลือ
              </span>
              <span className="text-sm font-black text-[var(--foreground)] leading-none">
                {formatNumber(player.wallet)}
              </span>
            </div>
          </div>

          {/* Add Money Button */}
          <button className="h-7 w-7 rounded-full bg-[#00acec] flex items-center justify-center text-white hover:bg-[#0092c9] transition-colors shadow-lg shadow-[#00acec]/20">
            <Plus size={16} strokeWidth={3} />
          </button>
        </div>
      </div>
    </div>
  );
}
