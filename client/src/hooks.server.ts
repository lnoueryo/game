import type { Handle, RequestEvent, ResolveOptions } from '@sveltejs/kit'
import { db } from '$lib/database'
import { base } from '$app/paths'
import type { MaybePromise } from '@sveltejs/kit/types/private';
// /** @type {import('@sveltejs/kit').Handle} */
// export const handle: Handle = async ({ event, resolve }) => {
//   console.log('Hello')
//   // get cookies from browser
//   const session = event.cookies.get('session')

//   if (!session) {
//     // if there is no session load page as normal
//     return await resolve(event)
//   }

//   // find the user based on the session
//   const user = await db.user.findUnique({
//     where: { userAuthToken: session },
//     select: { username: true, role: true },
//   })

//   // if `user` exists set `events.local`
//   if (user) {
//     event.locals.user = {
//       name: user.username,
//       role: user.role.name,
//     }
//   }

//   // load page as normal
//   return await resolve(event)
// }

const NO_AUTH_PATH = [
  '/login/',
  '/register/'
]
let previousPage : string = base;

const isUnneededAuth = (path: string) => {
  if(!path.endsWith('/')) path += '/'
  return NO_AUTH_PATH.includes(path)
}



const next = async(event: RequestEvent<Partial<Record<string, string>>>, resolve: (event: RequestEvent<Partial<Record<string, string>>>, opts?: ResolveOptions | undefined) => MaybePromise<Response>, path = ''): Promise<Response> => {
  let response;
  if(path) response = await redirect(path.replace('__data.js', ''))
  else response = await resolve(event)
  previousPage = event.url.pathname.replace('__data.js', '')
  return response
}

const redirect = (location: string, body?: string) => {
	return new Response(body, {
		status: 303,
		headers: { location }
	});
}

export const handle: Handle = async ({ event, resolve }) => {
  console.log('middleware')
  // if(event.url.pathname.includes('/__data.js')) return await next(event, resolve, previousPage)
  // get cookies from browser
  const session = event.cookies.get('session')
  if(!session) return await next(event, resolve)
  // console.log(event.request.method)
  // if(session && isUnneededAuth(event.url.pathname.replace('__data.js', ''))) {
  //   const path = previousPage == '/login' ? '/' : previousPage
  //   return await next(event, resolve, path)
  // }
  
  // if (!session && !isUnneededAuth(event.url.pathname.replace('__data.js', ''))) {
  //   console.log(2)
  //   return await next(event, resolve, '/login')
  // }
  
  // if (!session && isUnneededAuth(event.url.pathname.replace('__data.js', ''))) {
  //   console.log(3)
  //   previousPage = event.url.pathname.replace('__data.js', '')
  //   return await next(event, resolve)
  // }

  // find the user based on the session
  const user = await db.user.findUnique({
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