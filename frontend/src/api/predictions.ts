import { api } from "./Instand";

export const predictions = {
  get: (date: string) => api.get(`/today?date=${date}`),
};
