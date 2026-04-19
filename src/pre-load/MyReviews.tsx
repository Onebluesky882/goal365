"use client";

import { useEffect, useState } from "react";
import { myAnalyticApi } from "@/api/api";
import MatchCard from "@/components/Analytic/AnalyticCard";
import { Match } from "@/types/myAnalytic";
import { useAuth } from "@/GlobalContext/auth-provider";
import { useRouter } from "next/navigation";

type Props = {
  date: string;
  picked: boolean;
};

const MyReviews = ({ date, picked }: Props) => {
  const [matchesData, setMatchesData] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);
  useEffect(() => {
    const fetchMatches = async () => {
      try {
        setLoading(true);
        if (!date) return;
        const res = await myAnalyticApi.getReview(date, picked);
        setMatchesData(res.data || []);
      } finally {
        setLoading(false);
      }
    };

    fetchMatches();
  }, [date, picked]);

  const handlePickToggle = async (fixtureId: number, picked: boolean) => {
    // optimistic update
    setMatchesData((prev) =>
      prev.map((m) => (m.fixture_id === fixtureId ? { ...m, picked } : m)),
    );

    try {
      await myAnalyticApi.picked({
        date,
        id: String(fixtureId),
        picked,
      });
    } catch {
      // rollback
      setMatchesData((prev) =>
        prev.map((m) => {
          if (m.fixture_id === fixtureId) {
            return { ...m, picked: !picked };
          }
          return m;
        }),
      );
    }
  };

  const { session, isLoading } = useAuth();
  const router = useRouter();

  const userAllow = ["wansing882@gmail.com"];

  useEffect(() => {
    if (!isLoading) {
      if (!userAllow.includes(session?.user?.email || "")) {
        router.replace("/"); // 🔥 redirect
      }
    }
  }, [session, isLoading]);

  return (
    <>
      <div className="   px-4 py-8 max-sm:px-3 max-sm:py-2">
        {loading ? (
          <div>Loading...</div>
        ) : (
          <div className="grid max-sm:grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
            {matchesData
              .filter((item) => item.picked == true)
              .sort(
                (a, b) =>
                  new Date(a.timestamp).getTime() -
                  new Date(b.timestamp).getTime(),
              )

              .map((match) => (
                <MatchCard
                  key={match.id}
                  match={match}
                  date={date}
                  handlePickToggle={handlePickToggle}
                />
              ))}
          </div>
        )}
      </div>
    </>
  );
};

export default MyReviews;
