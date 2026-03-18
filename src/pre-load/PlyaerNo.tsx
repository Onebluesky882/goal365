"use client";
import { useAuth } from "@/GlobalContext/auth-provider";
import { usePlayerStore } from "@/store/playerNo";
import { useParams, useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import PlayerCard from "@/components/Player/PlayerCard";
import { playersApi } from "@/api/api";
import LoadingIndicators from "@/components/Loading_indicators";
import { Player } from "@/types/player";

const PlayerNo = () => {
  const { session, isLoading } = useAuth();
  const params = useParams<{ player_no: string }>();
  const playerNo = params.player_no;
  const { playerId, setPlayerId: setPlayer } = usePlayerStore();
  const [playerData, setPlayerData] = useState<Player | null>();
  const router = useRouter();

  useEffect(() => {
    if (!isLoading && !session) {
      router.push("/sign-in");
    }
  }, [isLoading, session, router]);

  useEffect(() => {
    if (!session || !playerNo) return;

    const getPlayer = async () => {
      const res = await playersApi.getPlayerById(playerNo);
      setPlayerData(res.data);
      setPlayer(Number(playerNo));
    };

    getPlayer();
  }, [session, playerNo, setPlayer]);

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

export default PlayerNo;
