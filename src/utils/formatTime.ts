export const formatTime = (iso: string) => {
  const date = new Date(iso);

  return date.toLocaleDateString("en-GB", {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
    timeZone: "Asia/Bangkok",
  });
};

export const getTime = (str: string) => {
  return str.split(", ")[1];
};
