"use client";

import { sportbookApi } from "@/api/api";
import React, { useEffect, useState } from "react";
import { Match } from "../../types/myAnalytic";
type Props = {
  date: string;
};

const SportBookMain = ({ date }: Props) => {
  const [matchesData, setMatchesData] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);
  useEffect(() => {
    const fetchMatches = async () => {
      try {
        setLoading(true);

        const res = await sportbookApi.getComingSoon(date, "coming_soon");

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
  }, [date]);
};
export default SportBookMain;
