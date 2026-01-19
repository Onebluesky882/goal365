import { TeamData } from "./predictions";

export type Nawin = {
  Id: number;
  Response: PredictionResponse[];
};

type PredictionResponse = {
  teams: {
    home: TeamData;
    away: TeamData;
  };
  h2h: H2HMatch[];
};

export type H2HMatch = {
  fixture: {
    id: number;
    date: string;
  };
  teams: {
    home: { name: string; winner: boolean | null };
    away: { name: string; winner: boolean | null };
  };
  goals: {
    home: number;
    away: number;
  };
};

export type TeamsRoot = {
  home: TeamData;
  away: TeamData;
};

