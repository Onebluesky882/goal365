import type { Match } from "@/types/predictions/types";
import { BarChart3, Trophy, Target } from "lucide-react";
import MatchCard from "./MatchCard";
type MatchListProps = {
  matches: Match[];
};

const MatchList = ({ matches }: MatchListProps) => {
  const winCount = matches.filter((m) => m.bet_pick.stake === "win").length;
  const winRate =
    matches.length > 0 ? Math.round((winCount / matches.length) * 100) : 0;

  return (
    <>
      {/* Stats Summary */}
      <div className="grid grid-cols-3 gap-4 mt-8 max-w-lg mx-auto">
        <div className="text-center p-4 rounded-xl bg-card border border-border">
          <BarChart3 className="h-6 w-6 mx-auto text-primary mb-2" />
          <div className="font-display text-2xl text-foreground">
            {matches.length}
          </div>
          <div className="text-xs text-muted-foreground">แมตช์ทั้งหมด</div>
        </div>
        <div className="text-center p-4 rounded-xl bg-card border border-border">
          <Trophy className="h-6 w-6 mx-auto text-secondary mb-2" />
          <div className="font-display text-2xl text-foreground">
            {winCount}
          </div>
          <div className="text-xs text-muted-foreground">ชนะ</div>
        </div>
        <div className="text-center p-4 rounded-xl bg-card border border-border">
          <Target className="h-6 w-6 mx-auto text-primary mb-2" />
          <div className="font-display text-2xl text-foreground">
            {winRate}%
          </div>
          <div className="text-xs text-muted-foreground">อัตราชนะ</div>
        </div>
      </div>

      {/* Matches Grid */}
      <section className="py-8 md:py-12">
        <div className="container">
          <div className="grid gap-6 md:grid-cols-2">
            {matches.map((match, index) => (
              <MatchCard key={match.fixture_id} match={match} index={index} />
            ))}
          </div>
        </div>
      </section>
    </>
  );
};

export default MatchList;
