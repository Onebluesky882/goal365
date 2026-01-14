// components/prediction/ComparisonTable.tsx

import { Comparison } from "../../../types/predictions";

type Props = {
  comparison: Comparison;
};

export default function ComparisonTable({ comparison }: Props) {
  return (
    <>
      {comparison && (
        <table className="w-full text-sm border rounded-lg overflow-hidden">
          <tbody>
            {Object.entries(comparison).map(([key, value]) => (
              <tr key={key} className="border-t">
                <td className="p-2 capitalize">{key.replace("_", " ")}</td>
                <td className="p-2 text-center">{value.home}</td>
                <td className="p-2 text-center">{value.away}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </>
  );
}
