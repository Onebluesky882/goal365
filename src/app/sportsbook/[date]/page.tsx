import { parseDate } from "@/common/pareDate";
import MenuBar from "@/components/sportbook/MenuBar";

type props = {
  params: Promise<{ date: string }>;
};

export default async function page({ params }: props) {
  const { date } = await params;
    const formatted = await parseDate(date);
    
    
  return (
    <>
      <MenuBar />
      <p>{formatted}</p>
    </>
  );
}
