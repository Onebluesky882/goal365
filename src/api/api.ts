import { toSnakeCase } from "@/common/GroupSnakeCase";
import { api } from "./instand";
import { toSnake } from "@/common/stringToSnake";
import { Nawin } from "../../types/nawin";

export const playersApi = {
  getPlayers: (userId: string) =>
    api.get("/players", { params: toSnake({ userId }) }),
  getPlayerById: (playerNo: string) =>
    api.get("/player", { params: toSnake({ playerNo }) }),
  CreatePlayer: (payload: { name: string; bio?: string; userId: string }) =>
    api.post(`/new-player`, toSnakeCase(payload)),
};

export const nawinApi = {
  getNawin: async (): Promise<Nawin[]> => {
    const res = await api.get<Nawin[]>("/nawin");
    return res.data;
  },
  postNawin: (fixtureId: string) => api.post("/nawin", { fixtureId }),
};
