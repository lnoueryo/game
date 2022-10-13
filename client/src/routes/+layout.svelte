<script lang="ts">
	import { page } from '$app/stores'
	import { applyAction, enhance } from '$app/forms'
  	import { goto, invalidateAll } from '$app/navigation'
  </script>

  <svelte:head>
	<title>SvelteKit Auth</title>
  </svelte:head>

  <nav>
	{#if !$page.data.user}
	  <a href="/login">Login</a>
	  <a href="/register">Register</a>
	{/if}

	{#if $page.data.user}
	  <!-- <a href="/admin">Admin</a> -->

	  <form
		action="/logout?/logout"
		method="POST"
		use:enhance={() => {
			return async ({ result }) => {
				invalidateAll()
				await applyAction(result)
			}
		}}
		>
			<button type="submit" class="btn btn-primary">Logout</button>
		</form>
	{/if}
  </nav>

  <main>
	<slot />
  </main>