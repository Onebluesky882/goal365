"use client";
import { playersApi } from "@/api/api";
import PlayerCard from "@/components/Player/PlayerCard";
import { useAuth } from "@/GlobalContext/auth-provider";
import { usePlayerStore } from "@/store/playerNo";
import { useParams, useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { Player } from "../../../../types/player";
import LoadingIndicators from "@/components/Loading_indicators";

const page = () => {
  const { session, isLoading } = useAuth();
  const params = useParams<{ player_no: string }>();
  const playerNo = params.player_no;
  const { playerId, setPlayerId: setPlayer } = usePlayerStore();
  const [playerData, setPlayerData] = useState<Player | null>();
  const router = useRouter();

  useEffect(() => {
    if (isLoading) return;
    if (!session) {
      router.push("/sign-in");
      return;
    }
    if (!playerNo) return;

    const getPlayer = async () => {
      try {
        const res = await playersApi.getPlayerById(playerNo);
        setPlayerData(res.data);
        setPlayer(Number(playerNo));
      } catch (err) {
        console.error(err);
      }
    };

    getPlayer();
  }, [isLoading, session, playerNo]);

  if (isLoading) return <LoadingIndicators />;
  return (
    <>
      {playerId && (
        <div className="">
          <PlayerCard player={playerData!} />
        </div>
      )}
    </>
  );
};

export default page;
