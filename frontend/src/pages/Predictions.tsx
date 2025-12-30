import { predictions } from "@/api/predictions";
import MatchList from "@/components/Prediction/MatchList";
import type { Match } from "@/types/predictions/types";
import { useEffect, useState } from "react";

const Predictions = () => {
  const [data, setData] = useState<Match[] | null>(null);
  useEffect(() => {
    // const date = new Date().toISOString().split("T")[0];
    const getPrediction = async () => {
      const res = await predictions.get(`2025-12-30`);
      if (Array.isArray(res.data)) {
        setData(res.data);
      }
    };

    getPrediction();
  }, [data, setData]);
  return (
    <div className="min-h-screen  ">
      {/* Hero Section */}
      <section className="py-12 md:py-16 border-b border-border">
        <div className="container ">
          <div className="text-center max-w-2xl mx-auto">
            <h2 className="font-display text-4xl md:text-5xl text-foreground tracking-wider mb-4">
              วิเคราะห์ฟอร์ม<span className="text-gradient-gold">ทีม</span>
            </h2>
            <p className="text-muted-foreground text-lg">
              ดูสถิติคะแนนหลังสุด 14, 12, 10, 7, 5 นัด พร้อม BetPick แนะนำ
            </p>
          </div>

          {data && <MatchList matches={data} />}
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-border py-6">
        <div className="container text-center text-sm text-muted-foreground">
          <p>ข้อมูลฟอร์มทีมและ BetPick สำหรับการวิเคราะห์เท่านั้น</p>
        </div>
      </footer>
    </div>
  );
};

export default Predictions;
