// types/match.types.ts
export type AsianHandicap = {
  line: number;
  home_odd: number;
  away_odd: number;
  favorite: string;
  is_favorite: boolean;
};

export type Goals = {
  total: number;
  average: string;
};

export type Last5Stats = {
  played: number;
  form: string;
  att: string;
  def: string;
  goals: {
    for: Goals;
    against: Goals;
  };
};

export type Match = {
  id: number;
  Date: string;
  timestamp: string;
  fixture_id: number;
  league: string;
  country: string;
  home: string;
  home_logo: string;
  away: string;
  away_logo: string;
  AsianHandicap: AsianHandicap[];
  picked: boolean;
  review: boolean;
  HomeScore: number;
  AwayScore: number;
  HomeLast5: Last5Stats;
  AwayLast5: Last5Stats;
  HomeStatic?: {
    form: string;
  };
  AwayStatic?: {
    form: string;
  };
  H2H?: H2HMatch[];
  score: Score;
};

// types/match.types.ts - เพิ่ม H2H types
export type H2HFixture = {
  id: number;
  date: string;
  timestamp: number;
  venue: {
    name: string;
    city: string;
  };
  status: {
    long: string;
    short: string;
  };
};

export type H2HTeam = {
  id: number;
  name: string;
  logo: string;
  winner: boolean;
};

export type H2HMatch = {
  fixture: H2HFixture;
  league: {
    id: number;
    name: string;
    country: string;
    round: string;
  };
  teams: {
    home: H2HTeam;
    away: H2HTeam;
  };
  goals: {
    home: number;
    away: number;
  };
  score: {
    fulltime: {
      home: number;
      away: number;
    };
  };
};

export type Score = {
  fulltime?: {
    home?: number | null;
    away?: number | null;
  };
};
