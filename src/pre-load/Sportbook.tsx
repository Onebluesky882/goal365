"use client";

import { sportbookApi } from "@/api/api";
import React, { useEffect, useState } from "react";
import { Match } from "../../types/myAnalytic";

const SportBook = () => {
  const [sportbook, setSportbook] = useState<Match>();
  const today = new Date().toISOString().slice(0, 10);
  useEffect(() => {
    const fetch = async () => {
      const res = await sportbookApi.getPreMatch(today, "pre_match");
      setSportbook(res.data);
    };

    fetch();
  }, [today]);

  return <pre>{JSON.stringify(sportbook, null, 2)}</pre>;
};

export default SportBook;
