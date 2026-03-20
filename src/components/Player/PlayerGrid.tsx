"use client";

import { Plus, Wallet, ShieldCheck, ChevronRight } from "lucide-react";
import { getUsageAge } from "@/common/getUsageAge";
import Image from "next/image";
import LoadingIndicators from "../Loading_indicators";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Player } from "@/types/player";

type PlayerProps = {
  players: Player[];
  canCreate: boolean;
  db: boolean;
  onCreate: () => void;
  onClick: (playerNo: number) => void;
};

export default function PlayersGrid({
  players,
  canCreate,
  onCreate,
  onClick,
  db,
}: PlayerProps) {
  const router = useRouter();
  const getInitials = (name: string) => {
    return name
      .split(" ")
      .map((n) => n[0])
      .join("")
      .toUpperCase()
      .substring(0, 2);
  };

  useEffect(() => {
    if (db) return;
    const timer = setTimeout(() => {
      router.push("/sign-in");
    }, 5000);

    return () => clearTimeout(timer);
  }, [db, router]);

  if (!db) return <LoadingIndicators />;
  return (
    <div className=" w-max-160  p-2 mx-2 sm:p-10 sm:mx-10 bg-card/10 border  rounded-md  animate-in fade-in slide-in-from-bottom-4 duration-500 ">
      {/* Header Section */}
      <div className="flex items-center justify-between mb-6 px-2">
        <div>
          <h2 className="text-xl font-bold text-foreground">DashBoard</h2>
          <p className="text-xs text-foreground">
            จัดการและเลือกตัวละครเพื่อเข้าสู่ระบบ
          </p>
        </div>
        <div className="text-right">
          <span className="text-[10px] font-bold text-foreground uppercase tracking-wider">
            บัญชีสมาชิก
          </span>
          <div className="text-sm font-bold text-[#00acec]">
            Lv. {players[0]?.level || 1}
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-5 ">
        {players.map((p) => (
          <div
            key={p.playerNo}
            onClick={() => onClick(p.playerNo)}
            className="group relative rounded-xl bg-card border border-border overflow-hidden transition-all hover:border-[#00acec]/50 hover:shadow-[0_0_20px_rgba(0,172,236,0.1)] cursor-pointer active:scale-[0.98]"
          >
            {/* ป้าย Player No (Floating Badge สไตล์ taptap) */}
            <div className="absolute top-0 left-0 bg-[#ff5f00] text-white text-[10px] font-bold px-2 py-1 uppercase z-10 rounded-br-lg">
              id : {p.playerNo}
            </div>

            <div className="p-5 flex items-center gap-5 ">
              {/* Avatar Section */}
              <div className="relative">
                {p.imageUrl ? (
                  <Image
                    src={p.imageUrl}
                    alt={p.name}
                    className="h-16 w-16 rounded-lg object-cover border border-border]"
                  />
                ) : (
                  <div className="h-16 w-16 rounded-lg bg-secondary flex items-center justify-center text-foreground font-bold text-xl border border-border">
                    {getInitials(p.name)}
                  </div>
                )}
                <div className="absolute -bottom-2 -right-2 bg-[#1a1a1a] border border-border rounded-full p-1">
                  <ShieldCheck size={14} className="text-[#00acec]" />
                </div>
              </div>

              {/* Info Section */}
              <div className="flex-1 min-w-0">
                <div className="flex justify-between items-start">
                  <h3 className="font-bold text-lg text-foreground truncate group-hover:text-[#00acec] transition-colors">
                    {p.name}
                  </h3>
                  <div className="text-[10px] bg-secondary text-foreground px-2 py-0.5 rounded border border-border">
                    Lv.{p.level}
                  </div>
                </div>

                <p className="text-[11px] text-foreground mt-0.5">
                  ใช้งานมาแล้ว {getUsageAge(p.createdAt)}
                </p>

                {/* Stats Area (คล้ายช่อง Input ของหน้า Login) */}
                <div className="mt-4 flex items-center gap-3 bg-input rounded-lg p-2 border border-border">
                  <div className="flex items-center gap-1.5 flex-1 border-r border-border">
                    <Wallet size={12} className="text-[#00acec]" />
                    <span className="text-xs font-bold text-foreground">
                      {p.wallet.toLocaleString()}
                    </span>
                  </div>
                  <div className="flex items-center gap-1 px-1">
                    <span className="text-[10px] text-foreground font-medium">
                      เข้าสู่ระบบ
                    </span>
                    <ChevronRight size={14} className="text-[#00acec]" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        ))}

        {/* ➕ Create Player Card (ออกแบบให้เหมือน Placeholder) */}
        <div
          onClick={() => {
            if (canCreate) onCreate();
          }}
          className={`relative flex flex-col items-center justify-center rounded-xl border-2 border-dashed h-31.5 transition-all
            ${
              canCreate
                ? "border-border bg-card hover:border-[#00acec]/40 hover:bg-[#00acec]/5 cursor-pointer"
                : "border-border opacity-40 cursor-not-allowed"
            }`}
        >
          {canCreate ? (
            <div className="flex flex-col items-center gap-2">
              <div className="p-2 rounded-full bg-secondary text-[#00acec]">
                <Plus size={24} />
              </div>
              <span className="text-sm font-bold text-foreground">
                สร้างตัวละครใหม่
              </span>
              <p className="text-[10px] text-foreground">
                คุณยังเหลือโควตาว่างสำหรับตัวละครใหม่
              </p>
            </div>
          ) : (
            <div className="flex flex-col items-center gap-1">
              <Plus size={20} className="text-foreground" />
              <span className="text-xs font-bold text-foreground">
                โควตาเต็มแล้ว
              </span>
              <p className="text-[10px] text-foreground">
                ปลดล็อกช่องเพิ่มที่เลเวล 10
              </p>
            </div>
          )}
        </div>
      </div>

      {/* Footer Support Info */}
      <div className="mt-8 text-center">
        <p className="text-[11px] text-foreground">
          มีปัญหาในการเข้าถึงตัวละคร?{" "}
          <span className="text-[#00acec] cursor-pointer hover:underline">
            ติดต่อฝ่ายบริการลูกค้า
          </span>
        </p>
      </div>
    </div>
  );
}
