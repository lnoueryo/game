import { db } from '$lib/database';

import type { PageServerLoad } from '.svelte-kit/types/src/routes/$types';
import { redirect, type Actions } from '@sveltejs/kit';

/** @type {import('@sveltekit/types').Load} */
export const load: PageServerLoad = async({params, url, locals}) => {
  //     console.log(url.searchParams.get('a'))
  if (!locals.user) {
    throw redirect(302, '/login')
  }

  locals.table = await db.table.findUnique({
    where: {key: (params as {key: string}).key},
    select: { key: true, title: true, cards: true, players: true },
  })

  return locals
};