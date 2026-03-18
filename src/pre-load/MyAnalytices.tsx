"use client";

import { myAnalyticApi } from "@/api/api";
import React, { useEffect, useState } from "react";
import { Match, PickedDto } from "../types/myAnalytic
import { toast } from "sonner";
import MatchCard from "@/components/MyAnalytic/MyAnalytic";
import { PickDate } from "@/components/MyAnalytic/PickDate";
import { formatDate } from "../utils/formatDate";
import { useStoreDate } from "@/store/date";

const MyAnalytics = () => {
  const [matchesData, setMatchesData] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);

  const [date, setDate] = useState<Date>(new Date());
  const pickDate = formatDate(date);
  useEffect(() => {
    const fetchMatches = async () => {
      try {
        setLoading(true);
        const res = await myAnalyticApi.getAnalytics(pickDate);
        setMatchesData(res.data || []);
      } finally {
        setLoading(false);
      }
    };

    fetchMatches();
  }, [date]);

  const handlePickToggle = async (fixtureId: number, picked: boolean) => {
    setMatchesData((prev) =>
      prev.map((m) => (m.fixture_id === fixtureId ? { ...m, picked } : m)),
    );

    if (!fixtureId) return;
    const body: PickedDto = {
      date: pickDate,
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
          <PickDate
            date={date}
            setDate={(newDate) => {
              if (newDate) setDate(newDate);
            }}
          />
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
                  date={pickDate}
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
