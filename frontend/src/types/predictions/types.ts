export type FormData = {
  home_form_14: number;
  away_form_14: number;
  home_form_12: number;
  away_form_12: number;
  home_form_10: number;
  away_form_10: number;
  home_form_7: number;
  away_form_7: number;
  home_form_5: number;
  away_form_5: number;
};

export type HandicapValue = {
  value: string;
  odd: string;
};

export type Handicap = {
  id: number;
  name: string;
  values: HandicapValue[];
};

export type BetPick = {
  odds: string;
  picked: string;
  stake: string;
};

export type Match = {
  fixture_id: number;
  date: string;
  league: string;
  timestamp: string;
  country: string;
  home: string;
  away: string;
  match_finish: string;
  handicap: Handicap;
  form_league_home_count: number;
  form_league_away_count: number;
  home_form_14: number;
  away_form_14: number;
  home_form_12: number;
  away_form_12: number;
  home_form_10: number;
  away_form_10: number;
  home_form_7: number;
  away_form_7: number;
  home_form_5: number;
  away_form_5: number;
  home_score: string;
  away_score: string;
  match_result: string;
  bet_pick: BetPick;
};
