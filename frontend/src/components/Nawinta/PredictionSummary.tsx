// components/prediction/PredictionSummary.tsx

import { Predictions } from "../../../types/predictions";

type Props = {
  predictions: Predictions;
};

export default function PredictionSummary({ predictions }: Props) {
  return (
    <>
      {predictions && (
        <div className="bg-gray-50 rounded-lg p-4">
          <p className="font-semibold mb-1">Advice</p>
          <p className="text-sm text-gray-700">{predictions.advice}</p>

          {predictions.winner?.name && (
            <p className="mt-2 text-sm">
              <strong>Winner:</strong> {predictions.winner.name}
            </p>
          )}
        </div>
      )}
    </>
  );
}
