import { io, Socket } from "socket.io-client";

let socket: Socket | null = null;

export const getSocket = (): Socket => {
  if (!socket) {
    socket = io("http://localhost:4000", {
      autoConnect: true,
    });

    socket.on("connect", () => {
      console.log("Connected socket:", socket?.id);
    });

    socket.on("disconnect", () => {
      console.log("Socket disconnected");
    });
  }
  return socket;
};