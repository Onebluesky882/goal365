import MyReviews from "@/pre-load/MyReviews";

type Props = {
  searchParams: {
    date?: string;
    picked?: string;
  };
};
export default function Page({ searchParams }: Props) {
  console.log(searchParams);
  return (
    <MyReviews
      date={searchParams.date ?? ""}
      picked={searchParams.picked === "true"}
    />
  );
}
