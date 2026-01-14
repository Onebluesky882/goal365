import PredictionView from "@/pre-load/Nawin";
import data from "./data.json";
import type { PredictionsRoot } from "../../../types/predictions";

const Page = () => {
  const roots = data as PredictionsRoot[];

  return (
    <div className="">
      {roots.map((root, rootIndex) =>
        root.response.map((item, index) => (
          <div className=" ">
            <PredictionView key={`${rootIndex}-${index}`} data={item} />
          </div>
        ))
      )}
    </div>
  );
};

export default Page;
