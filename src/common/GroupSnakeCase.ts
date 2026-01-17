export function toSnakeCase<T extends Record<string, unknown>>(obj: T) {
  const result: Record<string, unknown> = {};

  for (const key in obj) {
    result[camelToSnake(key)] = obj[key];
  }

  return result;
}

export function camelToSnake(key: string) {
  return key.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`);
}
