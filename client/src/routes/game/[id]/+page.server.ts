import { db } from '$lib/database';
import type { Action } from '.svelte-kit/types/src/routes/$types';
import { invalid, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from '../../../../.svelte-kit/types/src/routes/black-jack/$types';

/** @type {import('@sveltekit/types').Load} */
export const load: PageServerLoad = async({params, url, locals}) => {
  //     console.log(url.searchParams.get('a'))
  if (!locals.user) {
    throw redirect(302, '/login')
  }

  locals.game = await db.games.findUnique({
    where: { id: Number(params.id) },
    include: { tables: {select: {key: true, title: true, players: {select: {id: true, username: true}}}} },
  })

  return locals
};

/** @type {import('./$types').Actions} */
const createTable: Action = async ({ cookies, request, params }) => {
  const data = await request.formData()
  const title = data.get('title')
  const limit = Number(data.get('limit'))
  const session = cookies.get('session')
  if(isNaN(limit)) return invalid(400, { invalid: true });
  if(!session) return await redirect(303, '/login')
  if (
    !title ||
    !limit
  ) throw invalid(400, { invalid: true })

  const gameId = Number(params.id)

  const table = {
    title: title,
    limit: limit,
    gameId: gameId
  }

  let newTable;
  try {
    newTable = await db.$transaction(async (tx) => {
      const newTable = await tx.tables.create({
        data: table,
      })

      tx.players.update({
        where: {
          userAuthToken: session,
        },
        data: {
          tableId: newTable.key,
        },
      })

      return newTable
    })
    // console.log(`/game/${gameId}/${newTable.key}`)

  } catch (err) {
    console.error(err)
    throw invalid(400, {credentials: true})
  }
  throw redirect(302, `/game/${gameId}/${newTable.key}`)
}

/** @type {import('./$types').Actions} */
const deleteTable: Action = async ({ cookies, request, params }) => {
  const data = await request.formData()
  const key = data.get('table-id')
  try {
    const deleteUser = await db.tables.delete({
      where: {
        key: key,
      },
    })
  } catch (err) {
    console.error(err)
    invalid(400, {credentials: true})
  }

}
/** @type {import('./$types').Actions} */
export const actions: Actions = { createTable, deleteTable }