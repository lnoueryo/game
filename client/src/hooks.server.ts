import type { Handle, RequestEvent, ResolveOptions } from '@sveltejs/kit'
import { db } from '$lib/database'
import { base } from '$app/paths'
import type { MaybePromise } from '@sveltejs/kit/types/private';




const next = async(event: RequestEvent<Partial<Record<string, string>>>, resolve: (event: RequestEvent<Partial<Record<string, string>>>, opts?: ResolveOptions | undefined) => MaybePromise<Response>, path = ''): Promise<Response> => {
  let response;
  if(path) response = await redirect(path.replace('__data.js', ''))
  else response = await resolve(event)
  return response
}

const redirect = (location: string, body?: string) => {
	return new Response(body, {
		status: 303,
		headers: { location }
	});
}

export const handle: Handle = async ({ event, resolve }) => {

  const session = event.cookies.get('session')
  if(!session) return await next(event, resolve)

  const user = await db.players.findUnique({
    where: { userAuthToken: session },
    select: { username: true, role: true },
  })

  // if `user` exists set `events.local`
  if (user) {
    event.locals.user = {
      name: user.username,
      role: user.role.name,
    }
  }

  return await next(event, resolve)
}
  // load page as normal