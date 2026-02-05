import React from "react";
import { MatchCardProps } from "./MyAnalytic";
import { Match } from "../../../types/myAnalytic";

type FormProps = {
  match: Match;
  getFormArray: (form: string) => string[];
  getFormColor: (result: string) => string;
};

const FormDisplay = ({ match, getFormArray, getFormColor }: FormProps) => {
  return (
    <div className="mb-3">
      <div className="space-y-2">
        {/* Home Form */}
        {match.HomeStatic?.form && (
          <div className="flex items-center gap-2">
            <span className="text-xs text-gray-500 w-12">Home:</span>

            {/* scroll container */}
            <div className="flex gap-1 flex-1 min-w-0 overflow-x-auto">
              {getFormArray(match.HomeStatic.form).map((result, idx) => (
                <div
                  key={idx}
                  className={`w-6 h-6  shrink-0 flex items-center justify-center rounded text-xs font-bold ${getFormColor(result)}`}
                >
                  {result}
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Away Form */}
        {match.AwayStatic?.form && (
          <div className="flex items-center gap-2">
            <span className="text-xs text-gray-500 w-12">Away:</span>
            <div className="flex gap-1 flex-1 min-w-0 overflow-x-auto">
              {getFormArray(match.AwayStatic.form).map((result, idx) => (
                <div
                  key={idx}
                  className={`w-6 h-6  shrink-0 flex items-center justify-center rounded text-xs font-bold ${getFormColor(
                    result,
                  )}`}
                >
                  {result}
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default FormDisplay;
