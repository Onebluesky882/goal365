import { create } from "zustand";

type Day = {
  date: string | null;
  setDate: (date: string) => void;
};

export const useStoreDate = create<Day>((set) => ({
  date: null,
  setDate: (day) => set({ date: day }),
}));
