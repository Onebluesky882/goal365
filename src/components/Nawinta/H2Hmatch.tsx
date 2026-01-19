"use client";

import Image from "next/image";

type H2HTeam = {
  name: string;
  logo: string;
  winner: boolean | null;
};

type H2HMatch = {
  fixture: {
    id: number;
    date: string;
    status: {
      short: string;
    };
    venue: {
      name: string;
      city: string;
    };
  };
  league: {
    name: string;
    round: string;
    season: number;
    logo: string;
  };
  teams: {
    home: H2HTeam;
    away: H2HTeam;
  };
  goals: {
    home: number;
    away: number;
  };
};

type Props = {
  matches: H2HMatch[];
};

export default function H2HStatistics({ matches }: Props) {
  if (!matches || matches.length === 0) return null;

  const stats = matches.reduce(
    (acc, m) => {
      if (m.teams.home.winner) acc.home++;
      else if (m.teams.away.winner) acc.away++;
      else acc.draw++;
      return acc;
    },
    { home: 0, away: 0, draw: 0 }
  );

  const homeName = matches[0].teams.home.name;
  const awayName = matches[0].teams.away.name;

  return (
    <div className="mt-10 space-y-6">
      {/* Header */}
      <div className="text-center">
        <h2 className="text-xl font-bold">Head to Head</h2>
        <p className="text-sm text-gray-500">
          {homeName} vs {awayName} (Last {matches.length})
        </p>
      </div>

      {/* Summary */}
      <div className="grid grid-cols-3 gap-4">
        <Summary title={`${homeName} Wins`} value={stats.home} />
        <Summary title="Draws" value={stats.draw} />
        <Summary title={`${awayName} Wins`} value={stats.away} />
      </div>

      {/* Match list */}
      <div className="space-y-3">
        {matches.map((m) => (
          <div
            key={m.fixture.id}
            className="border rounded-lg p-3 flex items-center justify-between text-sm bg-card/50"
          >
            {/* Home */}
            <TeamMini team={m.teams.home} align="end" />

            {/* Score */}
            <div className="text-center font-bold">
              {m.goals.home} - {m.goals.away}
              <div className="text-[10px] text-gray-500">
                {new Date(m.fixture.date).toLocaleDateString()}
              </div>
            </div>

            {/* Away */}
            <TeamMini team={m.teams.away} align="start" />
          </div>
        ))}
      </div>
    </div>
  );
}

function Summary({ title, value }: { title: string; value: number }) {
  return (
    <div className="border rounded-md p-3 text-center bg-muted">
      <div className="text-xs text-gray-500">{title}</div>
      <div className="text-xl font-bold">{value}</div>
    </div>
  );
}

function TeamMini({
  team,
  align,
}: {
  team: H2HTeam;
  align: "start" | "end";
}) {
  return (
    <div className={`flex items-center gap-2 justify-${align}`}>
      <Image src={team.logo} alt="" width={28} height={28} />
      <span
        className={`text-xs ${
          team.winner ? "font-bold text-green-600" : ""
        }`}
      >
        {team.name}
      </span>
    </div>
  );
}