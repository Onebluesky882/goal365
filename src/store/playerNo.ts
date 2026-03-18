import { create } from "zustand";

type PlayerArea = {
  playerId: number | null;
  setPlayerId: (id: number) => void;
  logout: () => void;
};

export const rStore = create<PlayerArea>((set) => ({
  playerId: null,
  setPlayerId: (id) => set({ playerId: id }),
  logout: () => set({ playerId: null }),
}));
