import "react-day-picker/dist/style.css";
import { Calendar } from "../ui/calendar";

type PickDateProps = {
  date: Date | undefined;
  setDate: (date: Date | undefined) => void;
};

export const PickDate = ({ date, setDate }: PickDateProps) => {
  return (
    <Calendar
      mode="single"
      selected={date}
      onSelect={setDate}
      className="rounded-lg border"
    />
  );
};
