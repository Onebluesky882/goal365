"use client";

import { useEffect, useState } from "react";
import MatchTeams from "@/components/Nawinta/MatchTeams";
import LoadingIndicators from "@/components/Loading_indicators";
import { TeamData } from "../../types/predictions";
import { nawinApi } from "@/api/api";
import { H2HMatch } from "../../types/nawin";

type TeamsRoot = {
  home: TeamData;
  away: TeamData;
};

type InsertNawinFixTrueProp = {
  fixtureId: string;
  setFixtureId: (value: string) => void;
  onSubmit: () => void;
};

export default function PredictionView() {
  const [fixtureId, setFixtureId] = useState("");
  const [teams, setTeams] = useState<TeamsRoot[]>([]);
  const [loading, setLoading] = useState(false);
  const [h2h, setH2h] = useState<H2HMatch[]>([]);

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
    <div className="flex flex-col items-center pt-10 w-full">
      <p className="mb-2 font-semibold">Add fixture Id</p>

      <InsertNawinFixTrue
        fixtureId={fixtureId}
        setFixtureId={setFixtureId}
        onSubmit={submitFixture}
      />

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
    </div>
  );
}

function InsertNawinFixTrue({
  fixtureId,
  setFixtureId,
  onSubmit,
}: InsertNawinFixTrueProp) {
  return (
    <div className="flex items-center gap-2 mb-4">
      <input
        type="text"
        placeholder="Fixture ID"
        value={fixtureId}
        onChange={(e) => setFixtureId(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            onSubmit();
          }
        }}
        className="
          w-64
          px-3
          py-2
          border
          rounded-md
          text-sm
          focus:outline-none
          focus:ring-2
          focus:ring-blue-500
        "
      />
      <button
        onClick={onSubmit}
        className="
          px-4 py-2
          bg-blue-600
          text-white
          rounded-md
          hover:bg-blue-700
          text-sm
        "
      >
        Save
      </button>
    </div>
  );
}
