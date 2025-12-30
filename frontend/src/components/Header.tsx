import { TrendingUp } from "lucide-react";

const Header = () => {
  return (
    <header className="border-b border-border bg-card/50 backdrop-blur-sm sticky top-0 z-50">
      <div className="container py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="h-10 w-10 rounded-lg gradient-primary flex items-center justify-center shadow-glow-green">
              <TrendingUp className="h-5 w-5 text-primary-foreground" />
            </div>
            <div>
              <h1 className="font-display text-2xl text-foreground tracking-wider">
                FOOTBALL STATS
              </h1>
              <p className="text-xs text-muted-foreground">
                วิเคราะห์ฟอร์มทีม & BetPick
              </p>
            </div>
          </div>
          <div className="hidden md:flex items-center gap-2 text-sm text-muted-foreground">
            <span className="h-2 w-2 rounded-full bg-primary animate-pulse" />
            <span>อัพเดทล่าสุด</span>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
