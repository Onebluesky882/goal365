"use client";

import { myAnalyticApi } from "@/api/api";
import React, { useEffect, useState } from "react";
import { Match, PickedDto } from "../../types/myAnalytic";
import MatchCard from "../components/Myanalytic/MyAnalytic";
import { toast } from "sonner";

type Props = {
  date: string;
};

const MyAnalytics = ({ date }: Props) => {
  const [matchesData, setMatchesData] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchMatches = async () => {
      try {
        setLoading(true);
        const res = await myAnalyticApi.getAnalytics(date);
        setMatchesData(res.data || []);
      } finally {
        setLoading(false);
      }
    };

    fetchMatches();
  }, [date]);

  const handlePickToggle = async (fixtureId: number, picked: boolean) => {
    console.log("click");
    setMatchesData((prev) =>
      prev.map((m) => (m.fixture_id === fixtureId ? { ...m, picked } : m)),
    );

    if (!fixtureId) return;
    const body: PickedDto = {
      date: date,
      id: String(fixtureId),
      picked: picked,
    };

    try {
      await myAnalyticApi.picked(body);
      toast.success("picked! ");
    } catch (error) {
      toast.error("Fail! ");
    }

    setMatchesData((prev) =>
      prev.map((m) =>
        m.fixture_id === fixtureId ? { ...m, picked: picked } : m,
      ),
    );
  };

  return (
    <div className="container mx-auto px-4 py-8">
      {loading ? (
        <div>Loading...</div>
      ) : (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {matchesData
            .filter((item) => item.picked == false)
            .sort(
              (a, b) =>
                new Date(a.timestamp).getTime() -
                new Date(b.timestamp).getTime(),
            )

            .map((match) => {
              return (
                <MatchCard
                  key={match.id}
                  match={match}
                  date={date}
                  handlePickToggle={handlePickToggle}
                />
              );
            })}
        </div>
      )}
    </div>
  );
};

export default MyAnalytics;
