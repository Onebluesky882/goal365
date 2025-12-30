import axios from "axios";

export const api = axios.create({
  baseURL: "https://mytipster-production.up.railway.app",
});
