import { MatchWinner } from "@/components/sportbook/betTypeTable/MatchWinner";

export type SportbookRoot = {
  id: number;
  date: string;
  timestamp: string;
  fixture_id: number;
  home: string;
  away: string;
  league: string;
  country: string;
  status: string;
  MatchInfo: MatchInfo;
  bookmaker: string;
  match_winner?: MatchWinner;
  asian_handicap?: AsianHandicap[];
  over_under_full_time?: OverUnder[];
  both_teams_score?: BothTeamsScore;
  odd_even?: OddEven;
  win_to_nil_home?: WinToNilHome;
  win_to_nil_away?: WinToNilAway;
  exact_score?: ExactScore[];
  exact_goals_number?: ExactGoalsNumber[];
  double_chance?: DoubleChance;
  result_and_total_goals?: ResultAndTotalGoal[];
  corners_over_under?: CornersOverUnder[];
  cards_over_under?: CardsOverUnder;
  home_corners?: HomeCorners;
  away_corners?: AwayCorners;
  first_half_winner?: FirstHalfWinner;
  over_under_fh?: OverUnderFh[];
  asian_handicap_fh?: AsianHandicapFh[];
  corners_over_under_fh?: CornersOverUnderFh[];
  both_teams_score_fh?: BothTeamsScoreFh;
  odd_even_fh?: OddEvenFh;
  created_at?: string;
  updated_at?: string;
};

export type MatchInfo = {
  long: string;
  short: string;
};

export type MatchWinner = {
  home: number;
  draw: number;
  away: number;
};
export type FirstHalfWinner = MatchWinner;


export type AsianHandicap = {
  line: number;
  home_odd: number;
  away_odd: number;
  favorite: string;
  is_favorite: boolean;
};
export type AsianHandicapFh = AsianHandicap;

export type OverUnder = {
  value: string;
  over: number;
  under: number;
};
export type OverUnderFh = OverUnder;

export type BothTeamsScore = {
  yes: number;
  no: number;
};

export type OddEven = {
  odd: number;
  even: number;
};

export type WinToNilHome = {
  yes: number;
  no: number;
};

export type WinToNilAway = {
  yes: number;
  no: number;
};

export type ExactScore = {
  score: string;
  odd: number;
};

export type ExactGoalsNumber = {
  value: string;
  odd: number;
};

export type DoubleChance = {
  home_draw: number;
  home_away: number;
  draw_away: number;
};

export type ResultAndTotalGoal = {
  value: string;
  odd: number;
};

export type CornersOverUnder = {
  value: string;
  over: number;
  under: number;
};

export type CardsOverUnder = {
  value: string;
  over: number;
  under: number;
};

export type HomeCorners = {
  value: string;
  over: number;
  under: number;
};

export type AwayCorners = {
  value: string;
  over: number;
  under: number;
};

export type CornersOverUnderFh = {
  value: string;
  over: number;
  under: number;
};

export type BothTeamsScoreFh = {
  yes: number;
  no: number;
};

export type OddEvenFh = {
  odd: number;
  even: number;
};
