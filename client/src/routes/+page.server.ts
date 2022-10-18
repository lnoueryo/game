import { db } from '$lib/database';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

/** @type {import('@sveltekit/types').Load} */
export const load: PageServerLoad = async({params, url, locals}) => {
  //     console.log(url.searchParams.get('a'))
  if (!locals.user) {
    throw redirect(302, '/login')
  }

  if (locals.user.table) {
    throw redirect(302, `/game/${locals.user.table.gameId}/${locals.user.table.key}`)
  }

  locals.games = await db.game.findMany()

  return locals
};

