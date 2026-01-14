import PredictionView from "@/pre-load/Nawin";
import data from "./data.json";
import type { PredictionsRoot } from "../../../types/predictions";

const Page = () => {
  const roots = data as PredictionsRoot[];

  return (
    <div className="space-y-10">
      {roots.map((root, rootIndex) =>
        root.response.map((item, index) => (
          <PredictionView key={`${rootIndex}-${index}`} data={item} />
        ))
      )}
    </div>
  );
};

export default Page;
