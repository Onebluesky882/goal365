"use client";

import { useEffect, useRef, useState } from "react";
import MatchTeams from "@/components/Nawinta/MatchTeams";
import LoadingIndicators from "@/components/Loading_indicators";
import { nawinApi } from "@/api/api";
import { H2HMatch, TeamsRoot } from "../../types/nawin";
import InsertNawinFixTrue from "@/components/Nawinta/InputFixture";
import { useAuth } from "@/GlobalContext/auth-provider";
import { useRouter } from "next/navigation";

export default function PredictionView() {
  const [fixtureId, setFixtureId] = useState("");
  const [teams, setTeams] = useState<TeamsRoot[]>([]);
  const [loading, setLoading] = useState(true);
  const [h2h, setH2h] = useState<H2HMatch[]>([]);
  const [protecting, setProtecting] = useState(true);
  const { session, isLoading } = useAuth();
  const router = useRouter();
  const bottomRef = useRef<HTMLDivElement | null>(null);
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
      try {
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
      } catch (err) {
        console.error("getData failed:", err);
      } finally {
        setLoading(false);
      }
    };

    getData();
  }, []);

  useEffect(() => {
    if (isLoading) return;

    if (session?.user?.email === "wansing882@gmail.com") {
      setProtecting(false);
    } else {
      setProtecting(true);
    }
  }, [session, isLoading]);

  if (loading) return <LoadingIndicators />;

  return (
    <>
      {protecting ? (
        <div className="flex min-h-[70vh] flex-col items-center justify-center text-center px-4">
          <div className="max-w-md w-full bg-card border rounded-xl shadow-md p-6">
            <div className="flex justify-center mb-4">
              <div className="w-14 h-14 rounded-full bg-red-100 text-red-600 flex items-center justify-center text-2xl font-bold">
                !
              </div>
            </div>

            <h1 className="text-xl font-semibold text-foreground mb-2">
              Access Denied
            </h1>

            <p className="text-sm text-muted-foreground mb-6">
              You don’t have permission to access this page.
            </p>

            <button
              onClick={() => router.push("/")}
              className="w-full rounded-md bg-primary text-primary-foreground py-2 text-sm font-medium hover:opacity-90 transition"
            >
              Go to Home
            </button>
          </div>
        </div>
      ) : (
        <div className="flex relative flex-col items-center w-full mb-20">
          <button
            onClick={() =>
              bottomRef.current?.scrollIntoView({ behavior: "smooth" })
            }
            className=" border bg-accent/20 p-3 m-2 rounded-3xl cursor-pointer"
          >
            ล่างสุด
          </button>
          <div className="space-y-6 w-full max-w-3xl">
            {teams.length === 0 && (
              <p className="text-sm text-gray-500 text-center">
                No fixtures yet
              </p>
            )}

            {teams.map((t, i) => (
              <MatchTeams
                key={`${t.home.id}-${t.away.id}-${i}`}
                teams={t}
                h2h={h2h}
              />
            ))}
          </div>

          <div ref={bottomRef} className="mt-10">
            <InsertNawinFixTrue
              fixtureId={fixtureId}
              setFixtureId={setFixtureId}
              onSubmit={submitFixture}
            />
          </div>
        </div>
      )}
    </>
  );
}
