export function parseDate(input: string) {
  // รับ 05-02-2026
  const [dd, mm, yyyy] = input.split("-");

  if (!dd || !mm || !yyyy) return null;

  // แปลงเป็น ISO (2026-02-05)
  return `${yyyy}-${mm}-${dd}`;
}
