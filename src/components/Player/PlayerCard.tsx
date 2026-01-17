"use client";

import React from "react";
import { Player } from "../../../types/player";
import { formatNumber } from "@/common/converts_number";
import { Wallet, Plus, User } from "lucide-react";
import Image from "next/image";

type Props = {
  player: Player | null;
};

export default function WalletMinibar({ player }: Props) {
  if (!player) return null;

  return (
    <div className="sticky top-0 z-50 w-full px-4  flex justify-center pointer-events-none">
      <div className="pointer-events-auto flex items-center gap-4 bg-secondary/80 backdrop-blur-md border border-border px-4 py-2 rounded-full shadow-2xl shadow-black/50">
        {/* User Info Section */}
        <div className="flex items-center gap-2 pr-3 border-r border-border">
          <div className="h-7 w-7 rounded-full bg-card border border-border overflow-hidden flex items-center justify-center">
            {player.imageUrl ? (
              <Image
                src={player.imageUrl}
                alt={player.name}
                className="h-full w-full object-cover"
              />
            ) : (
              <User size={14} className="text-foreground" />
            )}
          </div>
          <div className="flex flex-col">
            <span className="text-[10px] font-bold text-foreground truncate max-w-20">
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
              <span className="text-[9px] text-foreground leading-none mb-0.5">
                ยอดเงินคงเหลือ
              </span>
              <span className="text-sm font-black text-foreground leading-none">
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
