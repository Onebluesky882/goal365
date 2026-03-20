// hooks/useSportbookData.ts
"use client";

import useSWR from "swr";
import { sportbookApi } from "@/api/api";
import { useStoreDate } from "@/store/date";
import { SportbookRoot } from "@/types/sportbook";

const fetchMatches = async ([status, date]: [string, string]): Promise<
  SportbookRoot[]
> => {
  const res = await sportbookApi.getPreMatch(date, status);
  return res.data ?? [];
};

export const useSportbookData = () => {
  const date = useStoreDate((s) => s.date);

  // preMatch
  const {
    data: preMatch,
    error: err1,
    isLoading: loading1,
  } = useSWR<SportbookRoot[]>(date ? ["pre_match", date] : null, fetchMatches, {
    revalidateOnFocus: false,
    keepPreviousData: true,
  });

  // comingSoon
  const {
    data: comingSoon,
    error: err2,
    isLoading: loading2,
  } = useSWR<SportbookRoot[]>(
    date ? ["coming_soon", date] : null,
    fetchMatches,
    {
      revalidateOnFocus: false,
      keepPreviousData: true,
    },
  );

  return {
    preMatch,
    comingSoon,
    error: err1 || err2,
    loading: loading1 || loading2,
  };
};
