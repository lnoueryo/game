<script lang="ts">
	import type { ActionData } from './$types'
	import { applyAction, enhance } from '$app/forms';
	import { goto, invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	export let form: ActionData;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div class="container">
		<!-- Button trigger modal -->
		<button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#createModal">
			新しく作成
		</button>
		<!-- Modal -->
		<div class="modal fade" id="createModal" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title" id="modalLabel">新しいテーブル</h5>
						<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
					</div>
					<div class="modal-body">
						<form
							action="?/createTable"
							method="POST"
							use:enhance={() => {
								return async ({ result }) => {
									invalidateAll()
									await applyAction(result)
									console.log(result)
									console.log(form)
									if(result.type == 'redirect') {
										goto(result.location)
									}
								}
								}}
							>
							<div class="mb-3">
								<label for="title" class="mb-2">タイトル</label>
								<input id="title" class="form-control form-control-lg" name="title" type="text" placeholder="タイトル" required>
							</div>
							<div class="mb-3">
								<label for="limit" class="mb-2">人数</label>
								<input id="limit" class="form-control form-control-lg" name="limit" type="text" placeholder="人数" required>
							</div>
							<div class="message-box">
								{#if form?.invalid}
								<p class="text-danger">Username and password is required.</p>
								{/if}

								{#if form?.credentials}
								<p class="text-danger">You have entered the wrong credentials.</p>
								{/if}
							</div>
							<div class="modal-footer">
								<button type="button" class="btn btn-secondary" data-bs-dismiss="modal" aria-label="Close">削除</button>
								<button type="submit" class="btn btn-primary" data-bs-dismiss="modal" aria-label="Close">作成</button>
							</div>
						</form>
					</div>

				</div>
			</div>
		</div>
		{#each $page.data.game.tables as table}
			<div class="card" style="width: 18rem;">
				<form
					class="text-end"
					action="?/deleteTable"
					method="POST"
					use:enhance={() => {
						return async ({ result }) => {
							invalidateAll()
							await applyAction(result)
						}
						}}
					>
					<input id="table-id" name="table-id" type="hidden" value={table.key}>
					<button type="submit" class="btn btn-danger">削除</button>
				</form>
				<div class="card-body">
					<h5 class="card-title">{table.title}</h5>
					{#each table.players as player}
						{player.username}
					{/each}
				</div>
				<a href="/game/{$page.data.game.id}/{table.key}" class="btn btn-primary">Go To Play</a>
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
