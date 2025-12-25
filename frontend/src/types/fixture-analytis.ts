export interface TeamIdentity {
  id: number;
  name: string;
  logo: string;
  winner?: boolean;
}

export interface FixtureData {
  id: number;
  date: string;
  venue: { name: string; city: string };
  status: { long: string; short: string };
}

export interface GoalStats {
  home: number;
  away: number;
}

export interface MinuteStats {
  [key: string]: { total: number | null; percentage: string | null };
}

export interface LeagueInfo {
  id: number;
  name: string;
  logo: string;
  flag: string;
  season: number;
  country: string;
}

export interface TeamStats {
  id: number;
  name: string;
  logo: string;
  last_5: {
    played: number;
    form: string;
    att: string;
    def: string;
    goals: {
      for: { total: number; average: string };
      against: { total: number; average: string };
    };
  };
  league: {
    form: string;
    goals: {
      for: { minute: MinuteStats };
      against: { minute: MinuteStats };
    };
    clean_sheet: { total: number };
    biggest: { wins: { home: string; away: string } };
  };
}

export interface Comparison {
  form: { home: string; away: string };
  att: { home: string; away: string };
  def: { home: string; away: string };
  poisson_distribution: { home: string; away: string };
  h2h: { home: string; away: string };
  goals: { home: string; away: string };
  total: { home: string; away: string };
}

export interface PredictionData {
  predictions: {
    winner: { id: number; name: string; comment: string };
    advice: string;
    goals: { home: string; away: string };
    percent: { home: string; draw: string; away: string };
  };
  teams: { home: TeamStats; away: TeamStats };
  comparison: Comparison;
  h2h: any[];
}

export interface PredictionItem {
  FixtureID: number;
  Fixture: {
    fixture: FixtureData;
    league: LeagueInfo;
    teams: { home: TeamIdentity; away: TeamIdentity };
    goals: GoalStats;
  };
  Predictions: PredictionData;
}

export interface PredictionResponse {
  Items: PredictionItem[];
}
