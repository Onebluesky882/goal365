export type Player = {
  playerNo: number;
  name: string;
  bio?: string; // optional
  imageUrl?: string;
  wallet: number;
  level: number;
  exp: number;
  createdAt: string;
};
