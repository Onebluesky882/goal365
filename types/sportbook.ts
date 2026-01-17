// Types
type OddsValue = {
  value: string;
  odd: string;
};

type Market = {
  type: string;
  values: OddsValue[];
};

type MatchInfo = {
  isLive: boolean;
  time: string;
  home: string;
  away: string;
  league: string;
  score: string;
};

type BookmakerData = {
  id: number;
  matchInfo: MatchInfo;
  markets: {
    matchWinner: Market;
    handicap: Market;
    overUnder: Market;
  };
};
