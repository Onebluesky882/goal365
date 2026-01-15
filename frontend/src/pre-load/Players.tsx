"use client";
import { playersApi } from "@/api/api";
import { useEffect, useState } from "react";
import { Player } from "../../types/player";
import LoadingIndicators from "@/components/Loading_indicators";
import { handleError } from "@/lib/handleErrors";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import { usePlayerStore } from "@/store/playerNo";
import { useAuth } from "@/GlobalContext/auth-provider";
import PlayersGrid from "@/components/Player/PlayerGrid";

const Players = () => {
  const { session } = useAuth();
  const [players, setPlayers] = useState<Player[]>([]);
  const [loading, setLoading] = useState(false);
  const [limitReached, setLimitReached] = useState(false);
  const { setPlayerId } = usePlayerStore();
  const router = useRouter();

  useEffect(() => {
    if (!session?.user?.id) return;

    const fetchPlayersAndCheckLimit = async () => {
      setLoading(true);
      try {
        const res = await playersApi.getPlayers(session?.user?.id!);
        setPlayers(res.data);

        if (res.data.length >= 2) {
          setLimitReached(true);
        } else {
          setLimitReached(false);
        }
      } catch (err) {
        handleError(err, "cannot fetch players");
      } finally {
        setLoading(false);
      }
    };

    fetchPlayersAndCheckLimit();
  }, [session?.user?.id]);
  if (loading) return <LoadingIndicators />;

  const handlePlayerClick = async (playerNo: number) => {
    setPlayerId(playerNo);
    router.push(`/player/${playerNo}`);
  };

  return (
    <div className="">
      <PlayersGrid
        onClick={handlePlayerClick}
        players={players}
        onCreate={() => {
          if (!session?.user) {
            toast.error("กรุณาเข้าสู่ระบบ");
            return;
          }
          if (limitReached) {
            toast.error("สามารถสร้าง Player ได้สูงสุด 2 คน");
            return;
          }
          router.push("/new-player");
        }}
        canCreate={!limitReached}
      />
    </div>
  );
};

export default Players;
