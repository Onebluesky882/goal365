export function getUsageAge(createdAt: string) {
  const created = new Date(createdAt);
  const now = new Date();

  const diffMs = now.getTime() - created.getTime();
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

  if (diffDays < 1) return "วันนี้";
  if (diffDays < 30) return `${diffDays} วัน`;
  if (diffDays < 365) return `${Math.floor(diffDays / 30)} เดือน`;
  return `${Math.floor(diffDays / 365)} ปี`;
}
