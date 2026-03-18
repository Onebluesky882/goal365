export interface SportbookRoot {
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
  over_under_full_time?: OverUnderFullTime[];
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
}

export interface MatchInfo {
  long: string;
  short: string;
}

export interface MatchWinner {
  home: number;
  draw: number;
  away: number;
}

export interface AsianHandicap {
  line: number;
  home_odd: number;
  away_odd: number;
  favorite: string;
  is_favorite: boolean;
}

export interface OverUnderFullTime {
  value: string;
  over: number;
  under: number;
}

export interface BothTeamsScore {
  yes: number;
  no: number;
}

export interface OddEven {
  odd: number;
  even: number;
}

export interface WinToNilHome {
  yes: number;
  no: number;
}

export interface WinToNilAway {
  yes: number;
  no: number;
}

export interface ExactScore {
  score: string;
  odd: number;
}

export interface ExactGoalsNumber {
  value: string;
  odd: number;
}

export interface DoubleChance {
  home_draw: number;
  home_away: number;
  draw_away: number;
}

export interface ResultAndTotalGoal {
  value: string;
  odd: number;
}

export interface CornersOverUnder {
  value: string;
  over: number;
  under: number;
}

export interface CardsOverUnder {
  value: string;
  over: number;
  under: number;
}

export interface HomeCorners {
  value: string;
  over: number;
  under: number;
}

export interface AwayCorners {
  value: string;
  over: number;
  under: number;
}

export interface FirstHalfWinner {
  home: number;
  draw: number;
  away: number;
}

export interface OverUnderFh {
  value: string;
  over: number;
  under: number;
}

export interface AsianHandicapFh {
  line: number;
  home_odd: number;
  away_odd: number;
  favorite: string;
  is_favorite: boolean;
}

export interface CornersOverUnderFh {
  value: string;
  over: number;
  under: number;
}

export interface BothTeamsScoreFh {
  yes: number;
  no: number;
}

export interface OddEvenFh {
  odd: number;
  even: number;
}
