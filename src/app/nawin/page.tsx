import PredictionView from "@/pre-load/Nawin";
import data from "./data.json";
import type { PredictionsRoot } from "../../../types/predictions";

const Page = () => {
  const roots = data as PredictionsRoot[];
  console.log("NEXT_PUBLIC_API_URL", process.env.NEXT_PUBLIC_API_URL);
  const env = process.env.NEXT_PUBLIC_API_URL;

  return (
    <div>
      <h1>{env}</h1>
      {roots.map((root, rootIndex) =>
        root.response.map((item, index) => (
          <PredictionView key={`${rootIndex}-${index}`} data={item} />
        ))
      )}
    </div>
  );
};

export default Page;
