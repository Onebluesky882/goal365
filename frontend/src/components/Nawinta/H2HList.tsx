// components/prediction/H2HList.tsx

import { H2HFixture } from "../../../types/predictions";

type Props = {
  h2h: H2HFixture[];
};

export default function H2HList({ h2h }: Props) {
  if (!h2h?.length) return null;

  return (
    <>
      {h2h && (
        <div>
          <h3 className="font-semibold mb-2">Head to Head</h3>
          <ul className="space-y-2">
            {h2h.map((m) => (
              <li key={m.fixture.id} className="text-sm border p-2 rounded">
                {m.teams.home.name} {m.goals.home} - {m.goals.away}{" "}
                {m.teams.away.name}
              </li>
            ))}
          </ul>
        </div>
      )}
    </>
  );
}
