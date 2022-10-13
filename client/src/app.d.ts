declare namespace App {
    interface Locals {
      user: {
        name: string
        role: string
      },
      games: [
        {
          id: int
          name: string
        }
      ],
      game: {
        id: int
        name: string
        tables: [
          {
            id: int
            players: string
          }
        ],
      }
    }
    // interface PageData {}
    // interface Platform {}
  }