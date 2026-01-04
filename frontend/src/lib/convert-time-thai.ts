export const formatDate = (dateString: string, timestamp?: string) => {
  const date = new Date(dateString);
  if (timestamp) {
    const [hours, minutes] = timestamp.split(":");
    date.setHours(parseInt(hours), parseInt(minutes));
  }
  return date.toLocaleDateString("th-TH", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};
