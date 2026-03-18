import { SportbookRoot } from "@/types/sportbook";
import { create } from "zustand@/types/sportbook

// store/sportbook.ts

type SportbookStore = {
  preMatch: SportbookRoot[];
  comingSoon: SportbookRoot[];
  setPreMatch: (data: SportbookRoot[]) => void;
  setComingSoon: (data: SportbookRoot[]) => void;
};

export const useSportbook = create<SportbookStore>((set) => ({
  preMatch: [],
  comingSoon: [],
  setPreMatch: (data) => set({ preMatch: data }),
  setComingSoon: (data) => set({ comingSoon: data }),
}));
