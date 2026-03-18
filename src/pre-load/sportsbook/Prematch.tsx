"use client";
import MenuBar from "@/components/sportbook/MenuBar";
import { useSportbookData } from "@/hooks/useSportBookData";
function PreMatch() {
  const { preMatch, comingSoon, loading, error } = useSportbookData();

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error</p>;

  // const groupCountry = Array.from(uniqueMap.values());
  // console.log("groupCountry :", groupCountry);
  // group league

  // diff time

  // search team
  if (!preMatch) return;
  const firstPreMatch = preMatch[0];
  const firstComing = comingSoon[0];
  return (
    <>
      <MenuBar />
      <h1>PreMatch</h1>
      <pre>{JSON.stringify(firstPreMatch, null, 2)}</pre>
      <h1>ComingSoon</h1>
      <pre>{JSON.stringify(firstComing, null, 2)}</pre>
    </>
  );
}
export default PreMatch;
