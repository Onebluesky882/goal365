import { toast } from "sonner";

export function handleError(err: unknown, fallback = "เกิดข้อผิดพลาด") {
  if (err instanceof Error) {
    toast.error(err.message);
  } else {
    toast.error(fallback);
  }
}
