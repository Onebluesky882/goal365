import PredictionView from "@/pre-load/Nawin";

export default async function page() {
  return (
    <div>
      <h1>{process.env.NEXT_PUBLIC_API_URL}</h1>
      <PredictionView />
    </div>
  );
}
