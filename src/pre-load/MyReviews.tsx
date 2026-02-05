"use client";

import { useEffect, useState } from "react";
import { myAnalyticApi } from "@/api/api";
import { Match } from "../../types/myAnalytic";
import MatchCard from "@/components/Myanalytic/MyAnalytic";

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
        const res = await myAnalyticApi.getReview(date, picked);
        setMatchesData(res.data || []);
      } finally {
        setLoading(false);
      }
    };

    fetchMatches();
  }, [date, picked]);

  return (
    <>
      <div className="container mx-auto px-4 py-8">
        {loading ? (
          <div>Loading...</div>
        ) : (
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            {matchesData.map((match) => (
              <MatchCard key={match.id} match={match} date={date} />
            ))}
          </div>
        )}
      </div>
    </>
  );
};

export default MyReviews;
