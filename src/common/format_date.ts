export const formatDate = (timestamp: string) => {
  try {
    const date = new Date(timestamp);
    return date.toLocaleString("th-TH", {
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  } catch {
    return timestamp;
  }
};
