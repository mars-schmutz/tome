export const signed = (num: number): string => {
  if (num < 0) return `${num}`;
  else return `+${num}`;
};
