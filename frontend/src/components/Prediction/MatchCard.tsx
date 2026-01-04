import type { Match } from "@/types/predictions/types";
import FormChart from "./FormChart";
import { Calendar, Trophy, Target, MapPin } from "lucide-react";
import { formatDate } from "@/lib/convert-time-thai";

export const matchesData: Match[] = [
  {
    fixture_id: 1437405,
    date: "2025-12-30",
    league: "Division 1",
    timestamp: "2025-12-30 19:25:00",
    country: "Saudi-Arabia",
    home: "Al Arabi SC",
    away: "Al Wehda Club",
    match_finish: "Match Finished",
    handicap: {
      id: 4,
      name: "Asian Handicap",
      values: [
        { value: "Home +0.25", odd: "2.00" },
        { value: "Away +0.25", odd: "1.80" },
        { value: "Home +0.5", odd: "1.80" },
        { value: "Away +0.5", odd: "2.00" },
      ],
    },
    form_league_home_count: 13,
    form_league_away_count: 13,
    home_form_14: 13,
    away_form_14: 15,
    home_form_12: 12,
    away_form_12: 15,
    home_form_10: 11,
    away_form_10: 15,
    home_form_7: 5,
    away_form_7: 14,
    home_form_5: 4,
    away_form_5: 10,
    home_score: "27%",
    away_score: "67%",
    match_result: "0-1",
    bet_pick: {
      odds: "",
      picked: "",
      stake: "",
    },
  },
];

type MatchCardProps = {
  match: Match;
  index?: number;
};

const MatchCard = ({ match, index = 0 }: MatchCardProps) => {
  // สมมติ matchesData: Match[]

  const hasBetPick = match.bet_pick.picked && match.bet_pick.picked.length > 0;

  const formattedDate = formatDate(match.timestamp);
  return (
    <div
      className="gradient-card rounded-xl border border-border overflow-hidden shadow-card animate-slide-up"
      style={{ animationDelay: `${index * 150}ms` }}
    >
      {/* Header */}
      <div className="bg-muted/50 px-4 py-3 flex items-center justify-between flex-wrap gap-2">
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <Calendar className="h-4 w-4" />
          <span>{match.fixture_id}</span>
          <span>{formattedDate}</span>
        </div>
        <div className="flex items-center gap-3">
          <div className="flex items-center gap-1 text-xs text-muted-foreground">
            <MapPin className="h-3 w-3" />
            <span>{match.league} | </span>
            <span>{match.country}</span>
          </div>
          <span className="text-xs px-2 py-1 rounded-full bg-secondary/20 text-secondary">
            {match.league}
          </span>
        </div>
      </div>

      {/* Teams & Score */}
      <div className="p-6">
        <div className="flex items-center justify-between mb-6">
          {/* Home Team */}
          <div className="flex-1 text-center">
            <div className="h-16 w-16 mx-auto mb-2 rounded-full bg-muted flex items-center justify-center text-2xl font-display">
              {match.home.charAt(0)}
            </div>
            <h3 className="font-semibold text-foreground text-sm md:text-base">
              {match.home}
            </h3>
            <p className="text-xs text-muted-foreground mt-1">เจ้าบ้าน</p>
          </div>

          {/* Score */}
          <div className="px-4 text-center">
            <div className="text-3xl md:text-4xl font-display text-foreground tracking-wider">
              {match.match_result}
            </div>
            <div className="mt-2 text-xs text-muted-foreground">
              {match.match_finish}
            </div>
          </div>

          {/* Away Team */}
          <div className="flex-1 text-center">
            <div className="h-16 w-16 mx-auto mb-2 rounded-full bg-muted flex items-center justify-center text-2xl font-display">
              {match.away.charAt(0)}
            </div>
            <h3 className="font-semibold text-foreground text-sm md:text-base">
              {match.away}
            </h3>
            <p className="text-xs text-muted-foreground mt-1">ทีมเยือน</p>
          </div>
        </div>

        {/* Score Percentages */}
        <div className="flex items-center justify-between mb-6 px-4">
          <div className="text-center">
            <div className="text-2xl font-display text-primary">
              {match.home_score}
            </div>
            <div className="text-xs text-muted-foreground">คะแนนรวม</div>
          </div>
          <div className="text-center">
            <Trophy className="h-6 w-6 mx-auto text-secondary mb-1" />
            <div className="text-xs text-muted-foreground">vs</div>
          </div>
          <div className="text-center">
            <div className="text-2xl font-display text-secondary">
              {match.away_score}
            </div>
            <div className="text-xs text-muted-foreground">คะแนนรวม</div>
          </div>
        </div>

        {/* Handicap Odds */}
        {match.handicap.values.length > 0 && (
          <>
            <div className="h-px bg-border mb-6" />
            <div className="mb-6">
              <h4 className="font-display text-sm text-muted-foreground mb-3">
                {match.handicap.name}
              </h4>
              <div className="grid grid-cols-2 gap-2">
                {match.handicap.values.map((hc, idx) => (
                  <div
                    key={idx}
                    className="flex items-center justify-between px-3 py-2 rounded-lg bg-muted/50 text-sm"
                  >
                    <span className="text-foreground">{hc.value}</span>
                    <span className="font-display text-primary">{hc.odd}</span>
                  </div>
                ))}
              </div>
            </div>
          </>
        )}

        {/* Divider */}
        <div className="h-px bg-border mb-6" />

        {/* Form Chart */}
        <div className="mb-6">
          <h4 className="font-display text-lg text-foreground mb-4 flex items-center gap-2">
            <Target className="h-5 w-5 text-primary" />
            ฟอร์มทีม (คะแนนหลังสุด)
          </h4>
          <FormChart
            formData={{
              home_form_14: match.home_form_14,
              away_form_14: match.away_form_14,
              home_form_12: match.home_form_12,
              away_form_12: match.away_form_12,
              home_form_10: match.home_form_10,
              away_form_10: match.away_form_10,
              home_form_7: match.home_form_7,
              away_form_7: match.away_form_7,
              home_form_5: match.home_form_5,
              away_form_5: match.away_form_5,
            }}
          />
        </div>

        {/* Bet Pick - only show if there's a pick */}
        {hasBetPick && (
          <>
            <div className="h-px bg-border mb-6" />
            <div className="gradient-primary rounded-lg p-4 shadow-glow-green">
              <div className="flex items-center justify-between">
                <div>
                  <div className="text-xs text-primary-foreground/80 mb-1">
                    BET PICK
                  </div>
                  <div className="font-display text-xl text-primary-foreground">
                    {match.bet_pick.picked}
                  </div>
                </div>
                {match.bet_pick.odds && (
                  <div className="text-right">
                    <div className="text-xs text-primary-foreground/80 mb-1">
                      ODDS
                    </div>
                    <div className="font-display text-2xl text-primary-foreground">
                      {match.bet_pick.odds}
                    </div>
                  </div>
                )}
              </div>
              {match.bet_pick.stake && (
                <div className="mt-3 pt-3 border-t border-primary-foreground/20">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-primary-foreground/80">สถานะ:</span>
                    <span className="font-semibold text-primary-foreground uppercase tracking-wider">
                      {match.bet_pick.stake === "win"
                        ? "✓ ชนะ"
                        : match.bet_pick.stake}
                    </span>
                  </div>
                </div>
              )}
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default MatchCard;
