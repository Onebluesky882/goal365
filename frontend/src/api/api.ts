import { toSnakeCase } from "@/app/common/GroupSnakeCase";
import { api } from "./instand";
import { toSnake } from "@/app/common/stringToSnake";

export const playersApi = {
  getPlayers: (userId: string) =>
    api.get("/players", { params: toSnake({ userId }) }),
  getPlayerById: (playerNo: string) =>
    api.get("/player", { params: toSnake({ playerNo }) }),
  CreatePlayer: (payload: { name: string; bio?: string; userId: string }) =>
    api.post(`/new-player`, toSnakeCase(payload)),
};
