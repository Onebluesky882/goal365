"use client";

import { myAnalyticApi } from "@/api/api";
import MatchCard from "@/components/Analytic/AnalyticCard";
import { PickDate } from "@/components/Analytic/PickDate";
import { Match, PickedDto } from "@/types/myAnalytic";
import { formatDate } from "@/utils/formatDate";
import React, { useEffect, useState } from "react";
import { toast } from "sonner";

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

  const goDown = () => {
    window.scrollTo({
      top: document.body.scrollHeight,
      behavior: "smooth",
    });
  };
  const goTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  return (
    <div className="container mx-auto px-4 py-8">
      {loading ? (
        <div>Loading...</div>
      ) : (
        <div className="grid gap-4 max-sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
          <button
            onClick={goDown}
            className="hidden  max-sm:block w-20 justify-self-center   items-center justify-center border border-gray-400 rounded-lg"
          >
            godown
          </button>
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
      <div className="flex justify-center">
        <button
          onClick={goTop}
          className="hidden mt-4 mb-8  max-sm:block   w-20    items-center  border border-gray-400 rounded-lg"
        >
          goTop
        </button>
      </div>
    </div>
  );
};

export default MyAnalytics;
