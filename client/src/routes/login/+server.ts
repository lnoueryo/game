import { json } from '@sveltejs/kit';
 
/** @type {import('./$types').RequestHandler} */
export async function POST({ request }: any) {
    console.log(typeof request)
    console.log('HE')
  const { a, b } = await request.json();
  return json(a + b);
}