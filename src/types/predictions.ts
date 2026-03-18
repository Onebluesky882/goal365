// types/prediction.ts

export type PredictionsRoot = {
  get: string;
  parameters: Parameters;
  errors: any[];
  results: number;
  paging: Paging;
  response: PredictionResponse[];
};

export type Parameters = {
  fixture: string;
};

export type Paging = {
  current: number;
  total: number;
};

export type PredictionResponse = {
  predictions: Predictions;
  league: League;
  teams: {
    home: TeamData;
    away: TeamData;
  };
  comparison: Comparison;
  h2h: H2HFixture[];
};

// ---------- Predictions ----------
export type Predictions = {
  winner: {
    id: number | null;
    name: string | null;
    comment: string | null;
  };
  win_or_draw: boolean | null;
  under_over: string | null;
  goals: {
    home: string | null;
    away: string | null;
  };
  advice: string | null;
  percent: {
    home: string;
    draw: string;
    away: string;
  };
};

// ---------- League ----------
export type League = {
  id: number;
  name: string;
  country: string;
  logo: string;
  flag: string | null;
  season: number;
};

// ---------- Team ----------
export type TeamData = {
  id: number;
  name: string;
  logo: string;
  last_5: Last5;
  league: TeamLeagueStats;
};

export type Last5 = {
  played: number;
  form: string | null;
  att: string | null;
  def: string | null;
  goals: {
    for: { total: number; average: string | null };
    against: { total: number; average: string | null };
  };
};

export type TeamLeagueStats = {
  form: string | null;
  fixtures: {
    played: FixtureCount;
    wins: FixtureCount;
    draws: FixtureCount;
    loses: FixtureCount;
  };
  goals: {
    for: GoalStats;
    against: GoalStats;
  };
  biggest: {
    wins: { home: string | null; away: string | null };
    loses: { home: string | null; away: string | null };
    goals: {
      for: { home: number | null; away: number | null };
      against: { home: number | null; away: number | null };
    };
  };
  clean_sheet: FixtureCount;
  failed_to_score: FixtureCount;
  penalty: {
    scored: { total: number; percentage: string | null };
    missed: { total: number; percentage: string | null };
    total: number;
  };
  lineups: Array<{ formation: string; played: number }>;
  cards: {
    yellow: Record<string, MinuteStat>;
    red: Record<string, MinuteStat>;
  };
};

export type FixtureCount = {
  home: number;
  away: number;
  total: number;
};

export type GoalStats = {
  total: FixtureCount;
  average: {
    home: string | null;
    away: string | null;
    total: string | null;
  };
  minute: Record<string, MinuteStat>;
  under_over: Record<string, { over: number; under: number }>;
};

export type MinuteStat = {
  total: number | null;
  percentage: string | null;
};

// ---------- Comparison ----------
export type Comparison = {
  form: ComparisonDetail;
  att: ComparisonDetail;
  def: ComparisonDetail;
  poisson_distribution: ComparisonDetail;
  h2h: ComparisonDetail;
  goals: ComparisonDetail;
  total: ComparisonDetail;
};

export type ComparisonDetail = {
  home: string;
  away: string;
};

// ---------- H2H ----------
export type H2HFixture = {
  fixture: {
    id: number;
    referee: string | null;
    timezone: string;
    date: string;
    timestamp: number;
    periods: {
      first: number | null;
      second: number | null;
    };
    venue: {
      id: number | null;
      name: string | null;
      city: string | null; // ✅ FIX
    };
    status: {
      long: string;
      short: string;
      elapsed: number | null;
      extra: number | null;
    };
  };
  league: League & {
    round: string | null;
    standings: boolean | null;
  };
  teams: {
    home: {
      id: number;
      name: string;
      logo: string;
      winner: boolean | null;
    };
    away: {
      id: number;
      name: string;
      logo: string;
      winner: boolean | null;
    };
  };
  goals: {
    home: number | null;
    away: number | null;
  };
  score: {
    halftime: {
      home: number | null;
      away: number | null;
    };
    fulltime: {
      home: number | null;
      away: number | null;
    };
    extratime: {
      home: number | null;
      away: number | null;
    };
    penalty: {
      home: number | null;
      away: number | null;
    };
  };
};
