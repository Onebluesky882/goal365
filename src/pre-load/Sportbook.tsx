"use client";

import { sportbookApi } from "@/api/api";
import React, { useEffect, useState } from "react";

const SportBook = () => {
  const [date, setDate] = useState<string>("");
  const [sportbook, setSportbook] = useState<any>([]);
  useEffect(() => {
    const today = new Date().toISOString().slice(0, 10);

    const fetch = async () => {
      const res = await sportbookApi.getPreMatch(today, "pre_match");
      setSportbook(res.data);
    };

    fetch();
  }, [date]);

  console.log(sportbook);
  return <pre>{JSON.stringify(sportbook, null, 2)}</pre>;
};

export default SportBook;

