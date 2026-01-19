export type InsertNawinFixTrueProp = {
  fixtureId: string;
  setFixtureId: (value: string) => void;
  onSubmit: () => void;
};

export default function InsertNawinFixTrue({
  fixtureId,
  setFixtureId,
  onSubmit,
}: InsertNawinFixTrueProp) {
  return (
    <div className="flex  items-center gap-2 bg-background/80 p-2 rounded-md   ">
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
