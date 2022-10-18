import { db } from '$lib/database';
import type { Action, RouteParams } from '.svelte-kit/types/src/routes/$types';
import { invalid, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from '../../../../.svelte-kit/types/src/routes/black-jack/$types';

/** @type {import('@sveltekit/types').Load} */
export const load: PageServerLoad = async({params, url, locals}) => {
  //     console.log(url.searchParams.get('a'))
  if (!locals.user) {
    throw redirect(302, '/login')
  }

  if (locals.user.table) {
    throw redirect(302, `/game/${locals.user.table.gameId}/${locals.user.table.key}`)
  }

  locals.game = await getGame(params)

  return locals
};

/** @type {import('./$types').Actions} */
const createTable: Action = async ({ cookies, request, params, locals }) => {

  const form = await request.formData()
  const title = form.get('title')
  const limit = Number(form.get('limit'))
  const session = cookies.get('session')
  if(isNaN(limit)) return invalid(400, { invalid: true });
  if(!session) return await redirect(303, '/login')
  if (
    !title ||
    !limit
  ) throw invalid(400, { invalid: true })

  const user = await db.player.findUnique({
    where: { userAuthToken: session },
    select: { id: true },
  })
  if(!user) throw redirect(303, '/login')

  locals.game = await getGame(params)

  const gameId = Number(params.id)
  const key = crypto.randomUUID()
  const table = {
    key: key,
    title: title,
    limit: limit,
    gameId: gameId,
    start: false,
    adminId: user.id,
    extraFields: locals.game.extraFields
  }

  let newTable;

  try {
    newTable = await db.$transaction(async (tx) => {
      const newTable = await tx.table.create({
        data: table,
      })
      const player = await tx.player.update({
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

  return {game: locals.game, url: `/game/${gameId}/${newTable.key}`}

}

/** @type {import('./$types').Actions} */
const deleteTable: Action = async ({ request, locals, params }) => {
  const form = await request.formData()
  const key = form.get('table-id')

  try {
    const deletetable = await db.table.delete({
      where: {
        key: key,
      },
    })

    locals.game = await getGame(params)

    return locals.game

  } catch (err) {
    console.error(err)
    invalid(400, {credentials: true})
  }

}

/** @type {import('./$types').Actions} */
const goToTable: Action = async ({ request, locals, params, cookies }) => {
  const form = await request.formData()
  const key = form.get('table-id') as string
  const session = cookies.get('session')
  if(!session) return await redirect(303, '/login')
  try {
    const player = await db.player.update({
      where: {
        userAuthToken: session,
      },
      data: {
        tableId: key,
      },
    })

    locals.game = await getGame(params)
    const gameId = Number(params.id)
    return {game: locals.game, url: `/game/${gameId}/${key}`}

  } catch (err) {
    console.error(err)
    invalid(400, {credentials: true})
  }

}

const getGame = (params: RouteParams) => {
  const id = Number(params.id)
  return db.game.findUnique({
    where: { id:  id},
    select: { id: true, name: true, extraFields: true, tables: {select: {key: true, title: true, adminId: true, players: {select: {id: true, username: true}}}} },
  })
}

/** @type {import('./$types').Actions} */
export const actions: Actions = { createTable, deleteTable, goToTable }

