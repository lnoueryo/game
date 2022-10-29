import { db } from '$lib/database';
import type { Action } from '.svelte-kit/types/src/routes/$types';
import { invalid, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from '../../../../.svelte-kit/types/src/routes/black-jack/$types';
import { gameClient } from '$lib/grpc';
import type{ Error, Game } from '$lib/grpc';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async({params, locals}) => {
  //     console.log(url.searchParams.get('a'))
  if (!locals.user) {
    throw redirect(302, '/login')
  }
  if (locals.user.table) {
    throw redirect(302, `/game/${locals.user.table.gameId}/${locals.user.table.key}`)
  }

  const game = await getGame(params)
  if(!game) throw redirect(302, '/')

  locals.game = game as App.Locals['game']

  return locals
};

/** @type {import('./$types').Action} */
const createTable: Action = async ({ cookies, request, params, locals }) => {

  const form = await request.formData()
  const title = form.get('title')
  const limit = Number(form.get('limit'))
  const session = cookies.get('session')
  if(!session) return await redirect(303, '/login')

  if(isNaN(limit)) throw invalid(400, { invalid: true });
  if (
    !title ||
    !limit
  ) throw invalid(400, { invalid: true })

  // 認証
  const user = await db.player.findUnique({
    where: { userAuthToken: session },
    select: { id: true, tableId: true },
  })
  if(!user) throw redirect(303, '/login')
  if(user.tableId) throw invalid(400, { invalid: true });

  const gameId = Number((params as {id: string}).id)
  const key = crypto.randomUUID()

  /*
  extraFieldsを取得するためだけにクエリを投げている。
  game.extraFieldsは参照のみのため、JSONファイルで管理する選択肢もあり。
  */
  locals.game = await getGame(params) as App.Locals['game']

  const table = {
    key: key,
    title: title,
    gameId: gameId,
    adminId: user.id,
    limit: limit,
    extraFields: locals.game.extraFields
  }

  let res;
  try {
    res = await new Promise((resolve, reject) => gameClient.createTable({table: table}, (err: Error, response: App.Locals['game']) => {
      if(err) {
        return reject(err)
      }
      resolve(response)
    }))
  } catch (error: any) {
    if(error.details == 'failed create') throw invalid(400, { failed: true });
    else throw invalid(400, { network: true });
  }

  locals.game = res  as App.Locals['game']

  return {game: locals.game, url: `/game/${gameId}/${key}`}

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

    locals.game = await getGame(params) as App.Locals['game']

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

    locals.game = await getGame(params) as App.Locals['game']
    const gameId = Number((params as {id: string}).id)
    return {game: locals.game, url: `/game/${gameId}/${key}`}

  } catch (err) {
    console.error(err)
    invalid(400, {credentials: true})
  }

}



const getGame = async(params: any) => {
  const id = Number(params.id)
  let res;
  try {
    res = await new Promise((resolve, reject) => gameClient.getGame({id: id}, (err: Error, response: App.Locals['game']) => {
      if(err) {
        return reject(err)
      }
      resolve(response)
    }))
  } catch (error) {
    console.log(error)
  }
  return res
}

/** @type {import('./$types').Actions} */
export const actions: Actions = { createTable, deleteTable, goToTable }
