"use client";
import LeagueHeader from "@/components/Nawinta/LeagueHeader";
import { PredictionResponse } from "../../types/predictions";
import MatchTeams from "@/components/Nawinta/MatchTeams";
import PredictionSummary from "@/components/Nawinta/PredictionSummary";
import PredictionPercent from "@/components/Nawinta/PredictionPercent";
import ComparisonTable from "@/components/Nawinta/ComparisonTable";
import H2HList from "@/components/Nawinta/H2HList";

type Props = {
  data: PredictionResponse;
};

export default function PredictionView({ data }: Props) {
  return (
    <div className="space-y-6">
      <LeagueHeader league={data.league} />
      <MatchTeams teams={data.teams} />
    </div>
  );
}
