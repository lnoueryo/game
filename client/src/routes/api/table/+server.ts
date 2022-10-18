import { db } from '$lib/database';
import { json } from '@sveltejs/kit';
import { createCards } from '$lib/card';
import type { player } from '@prisma/client';

/** @type {import('./$types').RequestHandler} */
export async function PUT({ request, locals, url }: any) {
  const { key, players } = await request.json()
  const newTable = await db.$transaction(async (tx) => {

    const table = await tx.table.updateMany({
      where: {
        key: key,
        start: false
      },
      data: {
        start: true,
      },
    })
    if(table.count == 0) return{};
    const cards = handCards(createCards(key), players)
    const newCards = await tx.card.createMany({
      data: cards,
      skipDuplicates: true, // Skip 'Bobo'
    })
    console.log(newCards)
    return table
  })
  locals.table = await db.table.findUnique({
    where: {key: key},
    select: { key: true, title: true, cards: true, players: true },
  })
  return json(locals.table);
}

export async function POST({ request }: any) {
  const { a, b } = await request.json();
  console.log(a, b)
  return json(a + b);
}

const handCards = (cards: {playerId: number}[], players: player[]) => {
  const ids: number[] = []
  players.forEach(player => {
    for (let i = 0; i < 2; i++) {
      ids.push(player.id)
    }
  })
  for (let i = 0; i < ids.length; i++) {
    cards[i]['playerId'] = ids[i];
  }
  return cards
}
