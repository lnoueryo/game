// import type { table } from '@prisma/client';
declare namespace App {
  interface Locals {
    user: {
      id: number
      username: string
      role: string
      tableId: string | null
    },
    games: {
      id: int
      name: string
    }[],
    game: {
      id: int
      name: string
      extraFields: {}
      tables: {
        key: string
        adminId: number
        title: string
        players: {
          id: number
          username: string
          role: string
          tableId: string | null
        }[]
      }[],
    }
    table: {
      key: string
      title: string
      limit: number
      start: boolean
      players: {
        id: number
        username: string
        role: string
        tableId: string | null
      }[]
    }
  }
  // interface PageData {}
  // interface Platform {}
}