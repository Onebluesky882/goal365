// components/prediction/MatchTeams.tsx

import { ReactNode } from "react";
import { TeamData } from "../../../types/predictions";
import Image from "next/image";
import { H2HMatch } from "../../../types/nawin";

type Props = {
  teams: {
    home: TeamData;
    away: TeamData;
  };
  h2h: H2HMatch[];
};

type SectionProps = {
  title: string;
  children: ReactNode;
};
type RowProps = {
  label: string;
  value: ReactNode;
};

export default function MatchTeams({ teams, h2h }: Props) {
  if (!teams) return null;
  const filteredH2H = h2h.filter((m) => {
    const homeName = teams.home.name;
    const awayName = teams.away.name;

    return (
      (m.teams.home.name === homeName && m.teams.away.name === awayName) ||
      (m.teams.home.name === awayName && m.teams.away.name === homeName)
    );
  });
  return (
    <div className="flex justify-center">
      <div className="grid grid-cols-7 p-5  w-250 ">
        <div className="col-span-3">
          <TeamFullCard title="home" team={teams.home} />
        </div>
        <div className="col-span-1 flex items-center justify-center">
          <span className="font-bold  ">VS</span>
        </div>
        <div className="col-span-3">
          <TeamFullCard title="away" team={teams.away} />
        </div>
      </div>
      {filteredH2H.length > 0 && (
        <div className="col-span-7 mt-6 ">
          <div className="border rounded-lg p-4 bg-muted/30">
            <div className="font-semibold mb-3">H2H (Last matches)</div>

            <div className="space-y-2">
              {filteredH2H.slice(0, 5).map((m) => (
                <div
                  key={m.fixture.id}
                  className="flex justify-between text-xs"
                >
                  <span>
                    {m.teams.home.name} {m.teams.away.name}
                  </span>
                  <span className="font-bold">
                    {m.goals.home} {m.goals.away}
                  </span>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

function Section({ title, children }: SectionProps) {
  return (
    <div className="space-y-1  ">
      <div className="font-semibold ">{title}</div>
      {children}
    </div>
  );
}

function Row({ label, value }: RowProps) {
  return (
    <div className="flex justify-between text-xs ">
      <span>{label}</span>
      <span className="font-medium">{value}</span>
    </div>
  );
}

function Divider() {
  return <div className="border-t my-1" />;
}

function TeamFullCard({ team, title }: { team: TeamData; title: string }) {
  const { league, last_5 } = team;
  const f = league.fixtures;

  return (
    <div className="border rounded-lg p-4 space-y-4 text-sm bg-card/50 ">
      {/* Header */}
      <div className="flex items-center gap-3">
        <Image
          src={team.logo}
          className="h-10 w-10"
          alt={""}
          width={100}
          height={100}
        />
        <div>
          <div className="font-bold">{team.name}</div>
          <div className="text-xs text-gray-500">{title}</div>
        </div>
      </div>

      {/* Form */}
      <div>
        <div className="font-semibold">Form</div>
        <div className="tracking-widest text-[12px] overflow-scroll ">
          {league.form}
        </div>
      </div>

      {/* Fixtures */}
      <Section title="Fixtures">
        <Row label="แข่งทั้งหมด" value={f.played.total} />
        <Row
          label="W / D / L"
          value={`${f.wins.total} / ${f.draws.total} / ${f.loses.total}`}
        />

        <Divider />

        <Row
          label="เหย้า"
          value={`แข่ง ${f.played.home} | W ${f.wins.home} D ${f.draws.home} L ${f.loses.home}`}
        />
        <Row
          label="เยือน"
          value={`แข่ง ${f.played.away} | W ${f.wins.away} D ${f.draws.away} L ${f.loses.away}`}
        />
      </Section>

      {/* Last 5 */}
      <Section title="Last 5">
        <Row label="Played" value={last_5.played} />
        <Row label="Form" value={last_5.form} />
        <Row label="Attack" value={last_5.att} />
        <Row label="Defence" value={last_5.def} />
        <Row
          label="Goals For"
          value={`${last_5.goals.for.total} (${last_5.goals.for.average})`}
        />
        <Row
          label="Goals Against"
          value={`${last_5.goals.against.total} (${last_5.goals.against.average})`}
        />
      </Section>

      {/* Goals */}
      <Section title="Goals">
        <Row label="ยิง (รวม)" value={league.goals.for.total.total} />
        <Row label="เสีย (รวม)" value={league.goals.against.total.total} />
        <Row label="เฉลี่ยยิง" value={league.goals.for.average.total} />
        <Row label="เฉลี่ยเสีย" value={league.goals.against.average.total} />
      </Section>

      {/* Biggest */}
      <Section title="Biggest">
        <Row label="ชนะมากสุด (เหย้า)" value={league.biggest.wins.home} />
        <Row label="ชนะมากสุด (เยือน)" value={league.biggest.wins.away} />
        <Row label="แพ้มากสุด (เหย้า)" value={league.biggest.loses.home} />
        <Row label="แพ้มากสุด (เยือน)" value={league.biggest.loses.away} />
      </Section>

      {/* Clean sheet */}
      <Section title="Clean Sheet">
        <Row label="Home" value={league.clean_sheet.home} />
        <Row label="Away" value={league.clean_sheet.away} />
        <Row label="Total" value={league.clean_sheet.total} />
      </Section>

      {/* Penalty */}
      <Section title="Penalty">
        <Row label="Scored" value={league.penalty.scored.total} />
        <Row label="Missed" value={league.penalty.missed.total} />
      </Section>
    </div>
  );
}
