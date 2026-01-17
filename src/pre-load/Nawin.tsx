"use client";
import { PredictionResponse } from "../../types/predictions";
import MatchTeams from "@/components/Nawinta/MatchTeams";

type Props = {
  data: PredictionResponse;
};

export default function PredictionView({ data }: Props) {
  return <MatchTeams teams={data.teams} />;
}
