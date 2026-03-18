// app/my-reviews/page.tsx
import MyReviews from "@/pre-load/MyReviews";

type Props = {
  searchParams: {
    date?: string;
    picked?: string;
  };
};

export default async function Page({ searchParams }: Props) {
  const { date, picked } = await searchParams;

  return <MyReviews date={date ?? ""} picked={picked === "true"} />;
}
