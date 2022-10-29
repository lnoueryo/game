import { sveltekit } from '@sveltejs/kit/vite';
import { Server } from 'socket.io';
import type {  UserConfig } from 'vite';
const websocket = {
  name: 'sveltekit-socket-io',
  configureServer(server: { httpServer: any }) {
    const io = new Server(server.httpServer)
    io.on('connection', (socket) => {
      // socket.broadcast.emit('hi');
      socket.on('game', (game) => {
        io.emit('game', game);
      });
      socket.on('table', (table) => {
        io.emit('table', table);
      });
    });
  }
}

const config: UserConfig = {
	plugins: [sveltekit(), websocket],
  server: {
    host: true
  }
};

export default config;
