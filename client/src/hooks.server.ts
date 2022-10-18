import type { Handle, RequestEvent, ResolveOptions } from '@sveltejs/kit'
import { db } from '$lib/database'

export const handle: Handle = async ({ event, resolve }) => {

  const session = event.cookies.get('session')
  if(!session) return await resolve(event)

  const user = await db.player.findUnique({
    where: { userAuthToken: session },
    select: { id: true, username: true, tableId: true, role: true, table: true },
  })

  // if `user` exists set `events.local`
  if (user) {
    event.locals.user = {
      id: user.id,
      username: user.username,
      role: user.role.name,
      table: user.table
    }
  }
  console.info('session checked')
  return await resolve(event)
}
  // load page as normal