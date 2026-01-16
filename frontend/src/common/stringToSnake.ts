// @typescript-eslint/no-explicit-any
export const toSnake = (obj: Record<string, unknown>) =>
  Object.fromEntries(
    Object.entries(obj).map(([k, v]) => [
      k.replace(/[A-Z]/g, (l) => `_${l.toLowerCase()}`),
      v,
    ])
  );
