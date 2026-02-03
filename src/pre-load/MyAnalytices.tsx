"use client";
// app/myanalytics/page.tsx
"use client";

import { myAnalyticApi } from "@/api/api";
import MatchCard from "@/components/Myanalytic/myanalytic";
import React, { useEffect, useState } from "react";
import { Match } from "../../types/myAnalytic";
type Props = {
  params: { date: string };
};

const MyAnalytics = ({ params }: Props) => {
  const [matchesData, setMatchesData] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);
  const { date } = params;
  useEffect(() => {
    const fetchMatches = async () => {
      try {
        setLoading(true);
        // const currentDate = date || pickDate;

        const res = await myAnalyticApi.getAnalytics(date);

        if (res.data && Array.isArray(res.data)) {
          setMatchesData(res.data);
        } else {
          setMatchesData([]);
        }

        console.log("Fetched data:", res.data);
      } catch (error) {
        console.error("Error fetching matches:", error);
        setMatchesData([]);
      } finally {
        setLoading(false);
      }
    };

    fetchMatches();
  }, [params]);

  const handlePickChange = (fixtureId: number, picked: boolean) => {
    console.log(`Match ${fixtureId} picked:`, picked);

    setMatchesData((prev) =>
      prev.map((match) =>
        match.fixture_id === fixtureId ? { ...match, picked } : match,
      ),
    );
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-6">
        <h1 className="text-2xl font-bold mb-2">Football Matches Analytics</h1>
        <p className="text-gray-600">Date: {date}</p>
      </div>

      {loading ? (
        <div className="flex justify-center items-center py-20">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
        </div>
      ) : matchesData.length === 0 ? (
        <div className="text-center py-20 text-gray-500">
          <p className="text-lg">No matches found for this date</p>
        </div>
      ) : (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {matchesData.map((match) => (
            <MatchCard
              key={match.id}
              match={match}
              onPickChange={handlePickChange}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default MyAnalytics;
