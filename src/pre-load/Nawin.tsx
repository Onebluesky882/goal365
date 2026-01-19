"use client";

import { useEffect, useState } from "react";
import MatchTeams from "@/components/Nawinta/MatchTeams";
import LoadingIndicators from "@/components/Loading_indicators";
import { nawinApi } from "@/api/api";
import { H2HMatch, TeamsRoot } from "../../types/nawin";
import InsertNawinFixTrue from "@/components/Nawinta/InputFixture";
import { useAuth } from "@/GlobalContext/auth-provider";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

export default function PredictionView() {
  const { session } = useAuth();
  const router = useRouter();

  const [fixtureId, setFixtureId] = useState("");
  const [teams, setTeams] = useState<TeamsRoot[]>([]);
  const [loading, setLoading] = useState(false);
  const [h2h, setH2h] = useState<H2HMatch[]>([]);

  useEffect(() => {
    setLoading(true);
    if (session && session.user?.email !== "wansing882@gmail.com") {
      toast.error("denied access");
      router.replace("/");
    }
  }, [session, router]);

  // ===== POST FIXTURE =====
  const submitFixture = async () => {
    if (!fixtureId.trim()) return;

    try {
      await nawinApi.postNawin(fixtureId);

      const res = await nawinApi.getNawin();

      const mappedTeams: TeamsRoot[] = [];
      const mappedH2H: H2HMatch[] = [];

      res.forEach((nawin) => {
        nawin.Response.forEach((r) => {
          mappedTeams.push(r.teams);

          if (r.h2h?.length) {
            mappedH2H.push(...r.h2h);
          }
        });
      });

      setTeams(mappedTeams);
      setH2h(mappedH2H);
      setFixtureId("");
    } catch (err) {
      console.error("postNawin failed:", err);
    }
  };

  useEffect(() => {
    const getData = async () => {
      setLoading(true);

      const res = await nawinApi.getNawin(); // Nawin[]

      const mappedTeams: TeamsRoot[] = [];
      const mappedH2H: H2HMatch[] = [];

      res.forEach((nawin) => {
        nawin.Response.forEach((r) => {
          mappedTeams.push(r.teams);

          if (r.h2h?.length) {
            mappedH2H.push(...r.h2h);
          }
        });
      });

      setTeams(mappedTeams);
      setH2h(mappedH2H);
      setLoading(false);
    };

    getData();
  }, []);

  if (loading) return <LoadingIndicators />;

  return (
    <div className="flex  relative flex-col items-center w-full  mb-20  ">
      <div className="space-y-6 w-full max-w-3xl">
        {teams.length === 0 && (
          <p className="text-sm text-gray-500 text-center">No fixtures yet</p>
        )}

        {teams.map((t, i) => (
          <MatchTeams
            key={`${t.home.id}-${t.away.id}-${i}`}
            teams={t}
            h2h={h2h}
          />
        ))}
      </div>
      <div className="absolute bottom-2    ">
        <InsertNawinFixTrue
          fixtureId={fixtureId}
          setFixtureId={setFixtureId}
          onSubmit={submitFixture}
        />
      </div>
    </div>
  );
}
