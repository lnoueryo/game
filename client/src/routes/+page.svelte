<script>
	import { page } from '$app/stores';
  </script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div class="container">
		{#each $page.data.games as game}
			<div class="card" style="width: 18rem;">
				<img src="/games/{game.image}" class="card-img-top" alt="">
				<div class="card-body">
				<h5 class="card-title">{game.name}</h5>
				<p class="card-text">{game.description}</p>
				<a href="/game/{game.id}" class="btn btn-primary">Go To Play</a>
				</div>
			</div>
		{/each}
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	/* h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	} */
</style>
