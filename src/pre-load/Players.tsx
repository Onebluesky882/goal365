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
  const { session, isLoading } = useAuth();
  const [players, setPlayers] = useState<Player[]>([]);
  const [limitReached, setLimitReached] = useState(false);
  const { setPlayerId } = usePlayerStore();
  const router = useRouter();
  const [db, setDb] = useState(false);

  // call userID
  useEffect(() => {
    if (isLoading) return;
    if (!session?.user?.id) {
      router.push("/");
      return;
    }
    const fetchPlayers = async () => {
      try {
        const userId = session.user?.id;
        const res = await playersApi.getPlayers(userId ?? "");

        setDb(true);
        const playersData = Array.isArray(res.data) ? res.data : [];
        setPlayers(playersData);

        if (res.data.length >= 2) {
          setLimitReached(true);
        } else {
          setLimitReached(false);
        }
      } catch (err) {
        handleError(err, "cannot fetch players");
        setPlayers([]);
        setLimitReached(false);
      } finally {
      }
    };

    fetchPlayers();
  }, [isLoading, session?.user?.id]);

  // if no db
  useEffect(() => {
    if (isLoading) return;

    const timer = setTimeout(() => {
      if (!db) {
        toast.error("server not connect");
        router.push("/");
      }
    }, 5000);

    return () => clearTimeout(timer);
  }, [isLoading, db]);
  if (isLoading) return <LoadingIndicators />;

  const handlePlayerClick = async (playerNo: number) => {
    setPlayerId(playerNo);
    router.push(`/player/${playerNo}`);
  };

  return (
    <div className="pt-2">
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
        db={db}
      />
    </div>
  );
};

export default Players;
