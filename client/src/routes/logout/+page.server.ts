import { goto } from '$app/navigation'
import { redirect } from '@sveltejs/kit'
import type { Action, Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
  // we only use this endpoint for the api
  // and don't need to see the page
//   throw redirect(302, '/')
}
const logout: Action = async ({ cookies }) => {
  cookies.set('session', '', {
    path: '/',
    expires: new Date(0),
  })
}

export const actions: Actions = { logout }
// export const actions: Actions = {
//   default({ cookies }) {
//     // eat the cookie
//     cookies.set('session', '', {
//       path: '/',
//       expires: new Date(0),
//     })

//     // goto('/')
//   },
// }